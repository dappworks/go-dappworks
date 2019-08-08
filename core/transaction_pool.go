// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either pubKeyHash 3 of the License, or
// (at your option) any later pubKeyHash.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//

package core

import (
	"bytes"
	"encoding/hex"
	"sort"
	"sync"

	"github.com/asaskevich/EventBus"
	corepb "github.com/dappley/go-dappley/core/pb"
	"github.com/dappley/go-dappley/network/network_model"
	"github.com/dappley/go-dappley/storage"
	"github.com/golang-collections/collections/stack"
	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-core/peer"
	logger "github.com/sirupsen/logrus"
)

const (
	NewTransactionTopic   = "NewTransaction"
	EvictTransactionTopic = "EvictTransaction"

	TxPoolDbKey = "txpool"

	BroadcastTx       = "BroadcastTx"
	BroadcastBatchTxs = "BraodcastBatchTxs"
)

var (
	txPoolSubscribedTopics = []string{
		BroadcastTx,
		BroadcastBatchTxs,
	}
)

type TransactionPool struct {
	txs        map[string]*TransactionNode
	pendingTxs []*Transaction
	tipOrder   []string
	sizeLimit  uint32
	currSize   uint32
	EventBus   EventBus.Bus
	mutex      sync.RWMutex
	netService NetService
}

func NewTransactionPool(netService NetService, limit uint32) *TransactionPool {
	txPool := &TransactionPool{
		txs:        make(map[string]*TransactionNode),
		pendingTxs: make([]*Transaction, 0),
		tipOrder:   make([]string, 0),
		sizeLimit:  limit,
		currSize:   0,
		EventBus:   EventBus.New(),
		mutex:      sync.RWMutex{},
		netService: netService,
	}
	txPool.ListenToNetService()
	return txPool
}

func (txPool *TransactionPool) GetTipOrder() []string { return txPool.tipOrder }

func (txPool *TransactionPool) ListenToNetService() {
	if txPool.netService == nil {
		return
	}

	for _, command := range txPoolSubscribedTopics {
		txPool.netService.Listen(command, txPool.GetCommandHandler(command))
	}
}

func (txPool *TransactionPool) GetCommandHandler(commandName string) network_model.CommandHandlerFunc {

	switch commandName {
	case BroadcastTx:
		return txPool.BroadcastTxHandler
	case BroadcastBatchTxs:
		return txPool.BroadcastBatchTxsHandler
	}
	return nil
}

func (txPool *TransactionPool) DeepCopy() *TransactionPool {
	txPoolCopy := TransactionPool{
		txs:       make(map[string]*TransactionNode),
		tipOrder:  make([]string, len(txPool.tipOrder)),
		sizeLimit: txPool.sizeLimit,
		currSize:  0,
		EventBus:  EventBus.New(),
		mutex:     sync.RWMutex{},
	}

	copy(txPoolCopy.tipOrder, txPool.tipOrder)

	for key, tx := range txPool.txs {
		newTx := tx.Value.DeepCopy()
		newTxNode := NewTransactionNode(&newTx)

		for childKey, childTx := range tx.Children {
			newTxNode.Children[childKey] = childTx
		}
		txPoolCopy.txs[key] = newTxNode
	}

	return &txPoolCopy
}

func (txPool *TransactionPool) SetSizeLimit(sizeLimit uint32) {
	txPool.sizeLimit = sizeLimit
}

func (txPool *TransactionPool) GetSizeLimit() uint32 {
	return txPool.sizeLimit
}

func (txPool *TransactionPool) GetTransactions() []*Transaction {
	txPool.mutex.RLock()
	defer txPool.mutex.RUnlock()
	return txPool.getSortedTransactions()
}

func (txPool *TransactionPool) GetNumOfTxInPool() int {
	txPool.mutex.RLock()
	defer txPool.mutex.RUnlock()

	return len(txPool.txs)
}

func (txPool *TransactionPool) ResetPendingTransactions() {
	txPool.mutex.Lock()
	defer txPool.mutex.Unlock()

	txPool.pendingTxs = make([]*Transaction, 0)
}

func (txPool *TransactionPool) GetAllTransactions() []*Transaction {
	txPool.mutex.RLock()
	defer txPool.mutex.RUnlock()

	txs := []*Transaction{}
	for _, tx := range txPool.pendingTxs {
		txs = append(txs, tx)
	}

	for _, tx := range txPool.getSortedTransactions() {
		txs = append(txs, tx)
	}

	return txs
}

//PopTransactionWithMostTips pops the transactions with the most tips
func (txPool *TransactionPool) PopTransactionWithMostTips(utxoIndex *UTXOIndex) *TransactionNode {
	txPool.mutex.Lock()
	defer txPool.mutex.Unlock()

	txNode := txPool.getMaxTipTransaction()
	tempUtxoIndex := utxoIndex.DeepCopy()
	if txNode == nil {
		return txNode
	}
	//remove the transaction from tip order
	txPool.tipOrder = txPool.tipOrder[1:]

	if result, err := VerifyTransaction(tempUtxoIndex, txNode.Value, 0); result {
		txPool.insertChildrenIntoSortedWaitlist(txNode)
		txPool.removeTransaction(txNode)
	} else {
		logger.WithError(err).Warn("Transaction Pool: Pop max tip transaction failed!")
		txPool.removeTransactionNodeAndChildren(txNode.Value)
		return nil
	}

	txPool.pendingTxs = append(txPool.pendingTxs, txNode.Value)
	return txNode
}

//Rollback adds a popped transaction back to the transaction pool. The existing transactions in txpool may be dependent on the input transaction_base. However, the input transaction should never be dependent on any transaction in the current pool
func (txPool *TransactionPool) Rollback(tx Transaction) {
	txPool.mutex.Lock()
	defer txPool.mutex.Unlock()

	rollbackTxNode := NewTransactionNode(&tx)
	txPool.updateChildren(rollbackTxNode)
	newTipOrder := []string{}

	for _, txid := range txPool.tipOrder {
		if _, exist := rollbackTxNode.Children[txid]; !exist {
			newTipOrder = append(newTipOrder, txid)
		}
	}

	txPool.tipOrder = newTipOrder

	txPool.addTransaction(rollbackTxNode)
	txPool.insertIntoTipOrder(rollbackTxNode)

}

//updateChildren traverses through all transactions in transaction pool and find the input node's children
func (txPool *TransactionPool) updateChildren(node *TransactionNode) {
	for txid, txNode := range txPool.txs {
	loop:
		for _, vin := range txNode.Value.Vin {
			if bytes.Compare(vin.Txid, node.Value.ID) == 0 {
				node.Children[txid] = txNode.Value
				break loop
			}
		}
	}
}

//Push pushes a new transaction into the pool
func (txPool *TransactionPool) Push(tx Transaction) {
	txPool.mutex.Lock()
	defer txPool.mutex.Unlock()
	if txPool.sizeLimit == 0 {
		logger.Warn("TransactionPool: transaction is not pushed to pool because sizeLimit is set to 0.")
		return
	}

	txNode := NewTransactionNode(&tx)

	if txPool.currSize != 0 && txPool.currSize+uint32(txNode.Size) >= txPool.sizeLimit {
		logger.WithFields(logger.Fields{
			"sizeLimit": txPool.sizeLimit,
		}).Warn("TransactionPool: is full.")

		return
	}

	txPool.addTransactionAndSort(txNode)

}

//CleanUpMinedTxs updates the transaction pool when a new block is added to the blockchain.
//It removes the packed transactions from the txpool while keeping their children
func (txPool *TransactionPool) CleanUpMinedTxs(minedTxs []*Transaction) {
	txPool.mutex.Lock()
	defer txPool.mutex.Unlock()

	for _, tx := range minedTxs {

		txNode, ok := txPool.txs[hex.EncodeToString(tx.ID)]
		if !ok {
			continue
		}
		txPool.insertChildrenIntoSortedWaitlist(txNode)
		txPool.removeTransaction(txNode)
		txPool.removeFromTipOrder(tx.ID)
	}
}

func (txPool *TransactionPool) removeFromTipOrder(txID []byte) {
	key := hex.EncodeToString(txID)

	for index, value := range txPool.tipOrder {
		if value == key {
			txPool.tipOrder = append(txPool.tipOrder[:index], txPool.tipOrder[index+1:]...)
			return
		}
	}

}

func (txPool *TransactionPool) cleanUpTxSort() {
	newTxOrder := []string{}
	for _, txid := range txPool.tipOrder {
		if _, ok := txPool.txs[txid]; ok {
			newTxOrder = append(newTxOrder, txid)
		}
	}
	txPool.tipOrder = newTxOrder
}

func (txPool *TransactionPool) getSortedTransactions() []*Transaction {

	nodes := make(map[string]*TransactionNode)
	scDeploymentTxExists := make(map[string]bool)

	for key, node := range txPool.txs {
		nodes[key] = node
		ctx := node.Value.ToContractTx()
		if ctx != nil && !ctx.IsExecutionContract() {
			scDeploymentTxExists[ctx.GetContractPubKeyHash().GenerateAddress().String()] = true
		}
	}

	var sortedTxs []*Transaction
	for len(nodes) > 0 {
		for key, node := range nodes {
			if !checkDependTxInMap(node.Value, nodes) {
				ctx := node.Value.ToContractTx()
				if ctx != nil {
					ctxPkhStr := ctx.GetContractPubKeyHash().GenerateAddress().String()
					if ctx.IsExecutionContract() {
						if !scDeploymentTxExists[ctxPkhStr] {
							sortedTxs = append(sortedTxs, node.Value)
							delete(nodes, key)
						}
					} else {
						sortedTxs = append(sortedTxs, node.Value)
						delete(nodes, key)
						scDeploymentTxExists[ctxPkhStr] = false
					}
				} else {
					sortedTxs = append(sortedTxs, node.Value)
					delete(nodes, key)
				}
			}
		}
	}

	return sortedTxs
}

func checkDependTxInMap(tx *Transaction, existTxs map[string]*TransactionNode) bool {
	for _, vin := range tx.Vin {
		if _, exist := existTxs[hex.EncodeToString(vin.Txid)]; exist {
			return true
		}
	}
	return false
}

func (txPool *TransactionPool) GetTransactionById(txid []byte) *Transaction {
	txPool.mutex.RLock()
	defer txPool.mutex.RUnlock()
	txNode, ok := txPool.txs[hex.EncodeToString(txid)]
	if !ok {
		return nil
	}
	return txNode.Value
}

func (txPool *TransactionPool) getDependentTxs(txNode *TransactionNode) map[string]*TransactionNode {

	toRemoveTxs := make(map[string]*TransactionNode)
	toCheckTxs := []*TransactionNode{txNode}

	for len(toCheckTxs) > 0 {
		currentTxNode := toCheckTxs[0]
		toCheckTxs = toCheckTxs[1:]
		for key, _ := range currentTxNode.Children {
			toCheckTxs = append(toCheckTxs, txPool.txs[key])
		}
		toRemoveTxs[hex.EncodeToString(currentTxNode.Value.ID)] = currentTxNode
	}

	return toRemoveTxs
}

// The param toRemoveTxs must be calculated by function getDependentTxs
func (txPool *TransactionPool) removeSelectedTransactions(toRemoveTxs map[string]*TransactionNode) {
	for _, txNode := range toRemoveTxs {
		txPool.removeTransactionNodeAndChildren(txNode.Value)
	}
}

//removeTransactionNodeAndChildren removes the txNode from tx pool and all its children.
//Note: this function does not remove the node from tipOrder!
func (txPool *TransactionPool) removeTransactionNodeAndChildren(tx *Transaction) {

	txStack := stack.New()
	txStack.Push(hex.EncodeToString(tx.ID))
	for txStack.Len() > 0 {
		txid := txStack.Pop().(string)
		currTxNode, ok := txPool.txs[txid]
		if !ok {
			continue
		}
		for _, child := range currTxNode.Children {
			txStack.Push(hex.EncodeToString(child.ID))
		}
		txPool.removeTransaction(currTxNode)
	}
}

//removeTransactionNodeAndChildren removes the txNode from tx pool.
//Note: this function does not remove the node from tipOrder!
func (txPool *TransactionPool) removeTransaction(txNode *TransactionNode) {
	txPool.disconnectFromParent(txNode.Value)
	txPool.EventBus.Publish(EvictTransactionTopic, txNode.Value)
	txPool.currSize -= uint32(txNode.Size)
	MetricsTransactionPoolSize.Dec(1)
	delete(txPool.txs, hex.EncodeToString(txNode.Value.ID))
}

//disconnectFromParent removes itself from its parent's node's children field
func (txPool *TransactionPool) disconnectFromParent(tx *Transaction) {
	for _, vin := range tx.Vin {
		if parentTx, exist := txPool.txs[hex.EncodeToString(vin.Txid)]; exist {
			delete(parentTx.Children, hex.EncodeToString(tx.ID))
		}
	}
}

func (txPool *TransactionPool) removeMinTipTx() {
	minTipTx := txPool.getMinTipTransaction()
	if minTipTx == nil {
		return
	}
	txPool.removeTransactionNodeAndChildren(minTipTx.Value)
	txPool.tipOrder = txPool.tipOrder[:len(txPool.tipOrder)-1]
}

func (txPool *TransactionPool) addTransactionAndSort(txNode *TransactionNode) {
	isDependentOnParent := false
	for _, vin := range txNode.Value.Vin {
		parentTx, exist := txPool.txs[hex.EncodeToString(vin.Txid)]
		if exist {
			parentTx.Children[hex.EncodeToString(txNode.Value.ID)] = txNode.Value
			isDependentOnParent = true
		}
	}

	txPool.addTransaction(txNode)

	txPool.EventBus.Publish(NewTransactionTopic, txNode.Value)

	//if it depends on another tx in txpool, the transaction will be not be included in the sorted list
	if isDependentOnParent {
		return
	}

	txPool.insertIntoTipOrder(txNode)
}

func (txPool *TransactionPool) addTransaction(txNode *TransactionNode) {
	txPool.txs[hex.EncodeToString(txNode.Value.ID)] = txNode
	txPool.currSize += uint32(txNode.Size)
	MetricsTransactionPoolSize.Inc(1)
}

func (txPool *TransactionPool) insertChildrenIntoSortedWaitlist(txNode *TransactionNode) {
	for _, child := range txNode.Children {
		parentTxidsInTxPool := txPool.GetParentTxidsInTxPool(child)
		if len(parentTxidsInTxPool) == 1 {
			txPool.insertIntoTipOrder(txPool.txs[hex.EncodeToString(child.ID)])
		}
	}
}

func (txPool *TransactionPool) GetParentTxidsInTxPool(tx *Transaction) []string {
	txids := []string{}
	for _, vin := range tx.Vin {
		txidStr := hex.EncodeToString(vin.Txid)
		if _, exist := txPool.txs[txidStr]; exist {
			txids = append(txids, txidStr)
		}
	}
	return txids
}

//insertIntoTipOrder insert a transaction into txSort based on tip.
//If the transaction is a child of another transaction, the transaction will NOT be inserted
func (txPool *TransactionPool) insertIntoTipOrder(txNode *TransactionNode) {
	index := sort.Search(len(txPool.tipOrder), func(i int) bool {
		if txPool.txs[txPool.tipOrder[i]] == nil {
			logger.WithFields(logger.Fields{
				"txid":             txPool.tipOrder[i],
				"len_of_tip_order": len(txPool.tipOrder),
				"len_of_txs":       len(txPool.txs),
			}).Warn("TransactionPool: the transaction in tip order does not exist in txs!")
			return false
		}
		if txPool.txs[txPool.tipOrder[i]].Value == nil {
			logger.WithFields(logger.Fields{
				"txid":             txPool.tipOrder[i],
				"len_of_tip_order": len(txPool.tipOrder),
				"len_of_txs":       len(txPool.txs),
			}).Warn("TransactionPool: the transaction in tip order does not exist in txs!")
		}
		return txPool.txs[txPool.tipOrder[i]].GetTipsPerByte().Cmp(txNode.GetTipsPerByte()) == -1
	})

	txPool.tipOrder = append(txPool.tipOrder, "")
	copy(txPool.tipOrder[index+1:], txPool.tipOrder[index:])
	txPool.tipOrder[index] = hex.EncodeToString(txNode.Value.ID)
}

func deserializeTxPool(d []byte) *TransactionPool {

	txPoolProto := &corepb.TransactionPool{}
	err := proto.Unmarshal(d, txPoolProto)
	if err != nil {
		println(err)
		logger.WithError(err).Panic("TxPool: failed to deserialize TxPool transactions.")
	}
	txPool := NewTransactionPool(nil, 1)
	txPool.FromProto(txPoolProto)

	return txPool
}

func LoadTxPoolFromDatabase(db storage.Storage, netService NetService, txPoolSize uint32) *TransactionPool {
	rawBytes, err := db.Get([]byte(TxPoolDbKey))
	if err != nil && err.Error() == storage.ErrKeyInvalid.Error() || len(rawBytes) == 0 {
		return NewTransactionPool(netService, txPoolSize)
	}
	txPool := deserializeTxPool(rawBytes)
	txPool.sizeLimit = txPoolSize
	txPool.netService = netService
	return txPool
}

func (txPool *TransactionPool) serialize() []byte {

	rawBytes, err := proto.Marshal(txPool.ToProto())
	if err != nil {
		logger.WithError(err).Panic("TxPool: failed to serialize TxPool transactions.")
	}
	return rawBytes
}

func (txPool *TransactionPool) SaveToDatabase(db storage.Storage) error {
	txPool.mutex.Lock()
	defer txPool.mutex.Unlock()
	return db.Put([]byte(TxPoolDbKey), txPool.serialize())
}

//getMinTipTransaction gets the transactionNode with minimum tip
func (txPool *TransactionPool) getMaxTipTransaction() *TransactionNode {
	txid := txPool.getMaxTipTxid()
	if txid == "" {
		return nil
	}
	for txPool.txs[txid] == nil {
		logger.WithFields(logger.Fields{
			"txid": txid,
		}).Warn("TransactionPool: max tip transaction is not found in pool")
		txPool.tipOrder = txPool.tipOrder[1:]
		txid = txPool.getMaxTipTxid()
		if txid == "" {
			return nil
		}
	}
	return txPool.txs[txid]
}

//getMinTipTransaction gets the transactionNode with minimum tip
func (txPool *TransactionPool) getMinTipTransaction() *TransactionNode {
	txid := txPool.getMinTipTxid()
	if txid == "" {
		return nil
	}
	return txPool.txs[txid]
}

//getMinTipTxid gets the txid of the transaction with minimum tip
func (txPool *TransactionPool) getMaxTipTxid() string {
	if len(txPool.tipOrder) == 0 {
		logger.Warn("TransactionPool: nothing in the tip order")
		return ""
	}
	return txPool.tipOrder[0]
}

//getMinTipTxid gets the txid of the transaction with minimum tip
func (txPool *TransactionPool) getMinTipTxid() string {
	if len(txPool.tipOrder) == 0 {
		return ""
	}
	return txPool.tipOrder[len(txPool.tipOrder)-1]
}

func (txPool *TransactionPool) ToProto() proto.Message {
	txs := make(map[string]*corepb.TransactionNode)
	for key, val := range txPool.txs {
		txs[key] = val.ToProto().(*corepb.TransactionNode)
	}
	return &corepb.TransactionPool{
		Txs:      txs,
		TipOrder: txPool.tipOrder,
		CurrSize: txPool.currSize,
	}
}

func (txPool *TransactionPool) FromProto(pb proto.Message) {
	for key, val := range pb.(*corepb.TransactionPool).Txs {
		txNode := NewTransactionNode(nil)
		txNode.FromProto(val)
		txPool.txs[key] = txNode
	}
	txPool.tipOrder = pb.(*corepb.TransactionPool).TipOrder
	txPool.currSize = pb.(*corepb.TransactionPool).CurrSize
	MetricsTransactionPoolSize.Clear()
	MetricsTransactionPoolSize.Inc(int64(len(txPool.txs)))
}

func (txPool *TransactionPool) BroadcastTx(tx *Transaction) {
	var broadcastPid peer.ID
	txPool.netService.SendCommand(BroadcastTx, tx.ToProto(), broadcastPid, network_model.Broadcast, network_model.NormalPriorityCommand)
}

func (txPool *TransactionPool) BroadcastTxHandler(command *network_model.DappRcvdCmdContext) {
	//TODO: Check if the blockchain state is ready
	txpb := &corepb.Transaction{}

	if err := proto.Unmarshal(command.GetData(), txpb); err != nil {
		logger.Warn(err)
	}

	tx := &Transaction{}
	tx.FromProto(txpb)
	//TODO: Check if the transaction is generated from running a smart contract
	//utxoIndex := core.NewUTXOIndex(n.GetBlockchain().GetUtxoCache())
	//if tx.IsFromContract(utxoIndex) {
	//	return
	//}

	txPool.Push(*tx)

	if command.IsBroadcast() {
		//relay the original command
		var broadcastPid peer.ID
		txPool.netService.Relay(command.GetCommand(), broadcastPid, network_model.NormalPriorityCommand)
	}
}

func (txPool *TransactionPool) BroadcastBatchTxs(txs []Transaction) {

	if len(txs) == 0 {
		return
	}

	transactions := NewTransactions(txs)

	var broadcastPid peer.ID
	txPool.netService.SendCommand(BroadcastBatchTxs, transactions.ToProto(), broadcastPid, network_model.Broadcast, network_model.NormalPriorityCommand)
}

func (txPool *TransactionPool) BroadcastBatchTxsHandler(command *network_model.DappRcvdCmdContext) {
	//TODO: Check if the blockchain state is ready
	txspb := &corepb.Transactions{}

	if err := proto.Unmarshal(command.GetData(), txspb); err != nil {
		logger.Warn(err)
	}

	txs := &Transactions{}

	//load the tx with proto
	txs.FromProto(txspb)

	for _, tx := range txs.GetTransactions() {
		//TODO: Check if the transaction is generated from running a smart contract
		//utxoIndex := core.NewUTXOIndex(n.GetBlockchain().GetUtxoCache())
		//if tx.IsFromContract(utxoIndex) {
		//	return
		//}
		txPool.Push(tx)
	}

	if command.IsBroadcast() {
		//relay the original command
		var broadcastPid peer.ID
		txPool.netService.Relay(command.GetCommand(), broadcastPid, network_model.NormalPriorityCommand)
	}

}
