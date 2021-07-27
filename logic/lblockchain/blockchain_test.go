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

package lblockchain

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dappley/go-dappley/core/transactionbase"
	blockchainMock "github.com/dappley/go-dappley/logic/lblockchain/mocks"
	"github.com/dappley/go-dappley/util"
	"os"
	"sync"
	"testing"

	"github.com/dappley/go-dappley/core/scState"
	"github.com/dappley/go-dappley/core/transaction"
	"github.com/dappley/go-dappley/core/utxo"
	"github.com/dappley/go-dappley/logic/ltransaction"
	"github.com/dappley/go-dappley/logic/lutxo"
	"github.com/dappley/go-dappley/logic/transactionpool"

	"github.com/dappley/go-dappley/common/hash"
	"github.com/dappley/go-dappley/core"
	"github.com/dappley/go-dappley/core/block"
	"github.com/dappley/go-dappley/core/blockchain"
	"github.com/dappley/go-dappley/logic/lblock"

	"github.com/dappley/go-dappley/common"
	"github.com/dappley/go-dappley/core/account"
	"github.com/dappley/go-dappley/storage"
	"github.com/dappley/go-dappley/storage/mocks"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	logger.SetLevel(logger.WarnLevel)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestCreateBlockchain(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)

	//find next block. This block should be the genesis block and its prev hash should be empty
	blk, err := bc.Next()
	assert.Nil(t, err)
	assert.Empty(t, blk.GetPrevHash())
}

func TestBlockchain_SetTailBlockHash(t *testing.T) {
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)

	tailHash := hash.Hash("TestHash")
	bc.SetTailBlockHash(tailHash)
	assert.Equal(t, tailHash, bc.GetTailBlockHash())

	newTailHash := hash.Hash("NewTestHash")
	bc.SetTailBlockHash(newTailHash)
	assert.NotEqual(t, tailHash, bc.GetTailBlockHash())
}

func TestBlockchain_HigherThanBlockchainTestHigher(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)
	blk := block.GenerateMockBlock()
	blk.SetHeight(1)
	assert.True(t, bc.IsHigherThanBlockchain(blk))
}

func TestBlockchain_HigherThanBlockchainTestLower(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)
	tailblk, _ := bc.GetTailBlock()
	blk := ltransaction.GenerateBlockWithCbtx(addr, tailblk)
	blk.SetHeight(1)
	bc.AddBlockContextToTail(PrepareBlockContext(bc, blk))

	assert.False(t, bc.IsHigherThanBlockchain(blk))

}

func TestBlockchain_IsInBlockchain(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 100000)

	blk := core.GenerateUtxoMockBlockWithoutInputs()
	bc.AddBlockContextToTail(PrepareBlockContext(bc, blk))

	isFound := bc.IsFoundBeforeLib([]byte("hash"))
	assert.True(t, isFound)

	isFound = bc.IsFoundBeforeLib([]byte("hash2"))
	assert.False(t, isFound)
}

func TestBlockchain_RollbackToABlock(t *testing.T) {
	//create a mock blockchain with max height of 5
	bc := GenerateMockBlockchainWithCoinbaseTxOnly(5)
	defer bc.db.Close()

	//find the hash at height 3
	blk, err := bc.GetBlockByHeight(3)
	assert.Nil(t, err)

	//rollback to height 3
	bc.Rollback(lutxo.NewUTXOIndex(bc.GetUtxoCache()), blk.GetHash(), scState.NewScState(bc.GetUtxoCache()))

	//the height 3 block should be the new tail block
	newTailBlk, err := bc.GetTailBlock()
	assert.Nil(t, err)
	assert.Equal(t, blk.GetHash(), newTailBlk.GetHash())

}

func TestBlockchain_AddBlockToTail(t *testing.T) {

	// Serialized data of an empty block (generated using `utx := NewGenesisBlock(Address{}) hex.EncodeToString(utx.Serialize())`)
	serializedBlk, _ := hex.DecodeString(`0a280a205e2d1835dd623d81317b6d896b2b541d4ccf4fd5000547f2466cd1492fe6ef4f20e0ebd9da0512430a20ba33bb7be2181496cbba9e426505e9fc4ea6f0e4c55fff708697d9c5ed9ff7bd121810ffffffffffffffffff01220b48656c6c6f20776f726c641a050a03989680`)
	db := new(mocks.Storage)

	// Create a blockchain for testing
	addr := account.NewAddress("dGDrVKjCG3sdXtDUgWZ7Fp3Q97tLhqWivf")
	bc := &Blockchain{blockchain.NewBlockchain(hash.Hash{}, hash.Hash{}), db, utxo.NewUTXOCache(db), nil, transactionpool.NewTransactionPool(nil, 128), nil, 1000000, &sync.Mutex{}}
	bc.SetState(blockchain.BlockchainInit)

	// Add genesis block
	genesis := NewGenesisBlock(addr, common.NewAmount(0))

	// Storage will allow blockchain creation to succeed
	db.On("Put", mock.Anything, mock.Anything).Return(nil)
	db.On("Get", []byte("utxo")).Return([]byte{}, nil)
	db.On("Get", []byte("scState")).Return([]byte{}, nil)
	db.On("Get", []byte("scState")).Return([]byte{}, nil)
	db.On("Get", mock.Anything).Return(serializedBlk, nil)
	db.On("EnableBatch").Return()
	db.On("DisableBatch").Return()
	// Flush invoked in AddBlockToTail twice
	db.On("Flush").Return(nil).Twice()

	err := bc.AddBlockContextToTail(PrepareBlockContext(bc, genesis))

	// Expect batch write was used
	//todo:to test Batch, if it's efficient than use it to save utxo
	//db.AssertCalled(t, "EnableBatch")
	//db.AssertCalled(t, "Flush")
	//db.AssertCalled(t, "DisableBatch")

	// Expect no error when adding genesis block
	assert.Nil(t, err)
	// Expect that blockchain tail is genesis block
	assert.Equal(t, genesis.GetHash(), hash.Hash(bc.GetTailBlockHash()))

	// Simulate a failure when flushing new block to storage
	simulatedFailure := errors.New("simulated storage failure")
	db.On("Flush").Return(simulatedFailure)

	// Add new block
	blk := block.NewBlock([]*transaction.Transaction{}, genesis, "")
	blk.SetHash([]byte("hash1"))

	blk.SetHeight(1)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk))

	// Expect the coinbase tx to go through
	assert.Equal(t, nil, err)
	// Expect that the block added is the blockchain tail
	assert.Equal(t, blk.GetHash(), hash.Hash(bc.GetTailBlockHash()))
}

func TestBlockchain_GetMaxHeight(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)
	assert.Equal(t, uint64(0), bc.GetMaxHeight())

	// Add new block
	genesis, err := bc.GetTailBlock()
	assert.Nil(t, err)
	blk1 := block.NewBlock([]*transaction.Transaction{}, genesis, "")
	blk1.SetHash([]byte("hash1"))
	blk1.SetHeight(5)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk1))
	assert.Nil(t, err)
	assert.Equal(t, uint64(5), bc.GetMaxHeight())

	blk2 := block.NewBlock([]*transaction.Transaction{}, blk1, "")
	blk2.SetHash([]byte("hash2"))
	blk2.SetHeight(1)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk2))
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), bc.GetMaxHeight())
}

func TestBlockchain_GetLIBHeight(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)
	assert.Equal(t, uint64(0), bc.GetLIBHeight())
	genesis, err := bc.GetTailBlock()
	assert.Nil(t, err)

	blk1 := block.NewBlock([]*transaction.Transaction{}, genesis, "")
	blk1.SetHash([]byte("hash1"))
	blk1.SetHeight(5)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk1))
	assert.Equal(t, uint64(0), bc.GetLIBHeight())
	assert.Nil(t, err)
	bc.SetLIBHash(blk1.GetHash())
	assert.Equal(t, uint64(5), bc.GetLIBHeight())
	bc.SetLIBHash(genesis.GetHash())
	assert.Equal(t, uint64(0), bc.GetLIBHeight())
}

func TestBlockchain_GetBlockByHash(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)
	genesis, err := bc.GetTailBlock()
	assert.Nil(t, err)

	// Add new blocks
	blk1 := block.NewBlock([]*transaction.Transaction{}, genesis, "")
	blk1.SetHash([]byte("hash1"))
	blk1.SetHeight(1)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk1))
	assert.Nil(t, err)

	blk2 := block.NewBlock([]*transaction.Transaction{}, blk1, "")
	blk2.SetHash([]byte("hash2"))
	blk2.SetHeight(2)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk2))
	assert.Nil(t, err)

	result, err := bc.GetBlockByHash(genesis.GetHash())
	assert.Equal(t, genesis.GetHash(), result.GetHash())
	assert.Nil(t, err)
	result, err = bc.GetBlockByHash(blk1.GetHash())
	assert.Equal(t, blk1.GetHash(), result.GetHash())
	assert.Nil(t, err)
	result, err = bc.GetBlockByHash(blk2.GetHash())
	assert.Equal(t, blk2.GetHash(), result.GetHash())
	assert.Nil(t, err)
}

func TestBlockchain_Iterator(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	libPolicy := &blockchainMock.LIBPolicy{}
	libPolicy.On("GetProducers").Return(nil)
	libPolicy.On("GetMinConfirmationNum").Return(6)
	libPolicy.On("IsBypassingLibCheck").Return(true)
	bc := CreateBlockchain(addr, s, libPolicy, transactionpool.NewTransactionPool(nil, 128), 1000000)
	genesis, err := bc.GetTailBlock()
	assert.Nil(t, err)

	// Add new block
	blk1 := block.NewBlock([]*transaction.Transaction{}, genesis, "")
	blk1.SetHash([]byte("hash1"))
	blk1.SetHeight(1)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk1))

	expected := &Blockchain{
		blockchain.NewBlockchain(blk1.GetHash(), genesis.GetHash()),
		s,
		bc.utxoCache,
		libPolicy,
		nil,
		nil,
		1000000,
		bc.mutex,
	}
	assert.Equal(t, expected, bc.Iterator())
}

func TestBlockchain_String(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)
	genesis, err := bc.GetTailBlock()
	assert.Nil(t, err)

	txs := []*transaction.Transaction{
		{
			ID: []byte("test1"),
			Vin: []transactionbase.TXInput{
				{
					Txid:      []byte("vin1"),
					Vout:      0,
					Signature: []byte("signature"),
					PubKey:    account.PubKeyHash([]byte{0xde, 0xad}),
				},
			},
			Vout: []transactionbase.TXOutput{
				{
					Value:      common.NewAmount(10),
					PubKeyHash: account.PubKeyHash([]byte{0xbe, 0xef}),
					Contract:   "testcontract",
				},
			},
			Tip:        common.NewAmount(1),
			GasLimit:   common.NewAmount(30000),
			GasPrice:   common.NewAmount(2),
			CreateTime: 99,
			Type:       transaction.TxTypeNormal,
		},
	}

	// Add new blocks
	blk1 := block.NewBlock(txs, genesis, "hello")
	blk1.SetHash([]byte("hash1"))
	blk1.SetHeight(1)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk1))
	assert.Nil(t, err)

	expected := "============ Block 36383631373336383331 ============\n" +
		"Height: 1\nPrev. block: 61653333363636666433316666366664333461653837633863333161316661643030653266353962643537373736396332356466653632653033393934623166\n\n--- Transaction: 7465737431\n     " +
		"Input 0:\n       TXID:      76696e31\n       Out:       0\n       Signature: 7369676e6174757265\n       PubKey:    dead\n     Output: 0\n       Value:  10\n       Script: beef\n       Contract: testcontract\n     GasLimit: 30000\n     GasPrice: 2\n     Type: 1\n\n\n\n" +
		"============ Block 61653333363636666433316666366664333461653837633863333161316661643030653266353962643537373736396332356466653632653033393934623166 ============\n" +
		"Height: 0\nPrev. block: \n\n--- Transaction: 1e3cb55cd3ae308b894a8a85869fb940e2c91e8bd115502f90f66e5abc1da287\n" +
		"     Input 0:\n       TXID:      \n       Out:       -1\n       Signature: \n       PubKey:    48656c6c6f20776f726c64\n     Output: 0\n       Value:  10000000000\n       Script: 003b21529355fe27c192eeda99a8330caaf16f5f0a\n       Contract: \n     GasLimit: 0\n     GasPrice: 0\n     Type: 3\n\n\n\n"
	assert.Equal(t, expected, bc.String())
	fmt.Println(bc.String())
}

func TestBlockchain_AddBlockToDb(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)
	genesis, err := bc.GetTailBlock()
	assert.Nil(t, err)

	txs := []*transaction.Transaction{
		{
			ID:         []byte("test1"),
			Vin:        transactionbase.MockTxInputs(),
			Vout:       transactionbase.MockTxOutputs(),
			Tip:        common.NewAmount(1),
			GasLimit:   common.NewAmount(30000),
			GasPrice:   common.NewAmount(2),
			CreateTime: 0,
			Type:       transaction.TxTypeNormal,
		},
		{
			ID:         []byte("test2"),
			Vin:        transactionbase.MockTxInputs(),
			Vout:       transactionbase.MockTxOutputs(),
			Tip:        common.NewAmount(15),
			GasLimit:   common.NewAmount(200),
			GasPrice:   common.NewAmount(1),
			CreateTime: 100000,
			Type:       transaction.TxTypeReward,
		},
	}

	blk := block.NewBlock(txs, genesis, "")
	blk.SetHash([]byte("hash1"))
	blk.SetHeight(1)
	err = bc.AddBlockContextToTail(PrepareBlockContext(bc, blk))
	assert.Nil(t, err)

	bc.AddBlockToDb(blk)

	// check that blk is stored in db
	result, err := bc.db.Get(blk.GetHash())
	assert.Equal(t, blk.Serialize(), result)
	assert.Nil(t, err)

	// check that blk hash is stored in db
	result, err = bc.db.Get(util.UintToHex(blk.GetHeight()))
	assert.Equal(t, []uint8("hash1"), result)
	assert.Nil(t, err)

	// check that blk's txs are stored in db
	txJournal1 := transaction.NewTxJournal(txs[0].ID, txs[0].Vout)
	expected, err := txJournal1.SerializeJournal()
	assert.Nil(t, err)
	result, err = bc.db.Get([]byte("tx_journal_test1"))
	assert.Equal(t, expected, result)
	assert.Nil(t, err)

	txJournal2 := transaction.NewTxJournal(txs[1].ID, txs[1].Vout)
	expected, err = txJournal2.SerializeJournal()
	assert.Nil(t, err)
	result, err = bc.db.Get([]byte("tx_journal_test2"))
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestBlockchain_setTailBlockHash(t *testing.T) {
	//create a new block chain
	s := storage.NewRamStorage()
	defer s.Close()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128), 1000000)

	testHash := []byte("test")
	err := bc.setTailBlockHash(testHash)
	assert.Nil(t, err)
	// check db
	storedHash, err := bc.db.Get(tailBlockHash)
	assert.Equal(t, testHash, storedHash)
	assert.Nil(t, err)
	// check bc.tailBlockHash
	assert.Equal(t, hash.Hash(testHash), bc.GetTailBlockHash())
}

func BenchmarkBlockchain_AddBlockToTail(b *testing.B) {
	//create a new block chain

	s := storage.NewRamStorage()
	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")

	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 1280000), 100000)
	var accounts []*account.Account
	for i := 0; i < 10; i++ {
		acc := account.NewAccount()
		accounts = append(accounts, acc)
	}

	for i := 0; i < b.N; i++ {

		tailBlk, _ := bc.GetTailBlock()
		txs := []*transaction.Transaction{}
		utxo := lutxo.NewUTXOIndex(bc.GetUtxoCache())
		cbtx := ltransaction.NewCoinbaseTX(accounts[0].GetAddress(), "", uint64(i+1), common.NewAmount(0))
		utxo.UpdateUtxo(&cbtx)
		txs = append(txs, &cbtx)
		for j := 0; j < 10; j++ {
			sendTxParam := transaction.NewSendTxParam(accounts[0].GetAddress(), accounts[0].GetKeyPair(), accounts[i%10].GetAddress(), common.NewAmount(1), common.NewAmount(0), common.NewAmount(0), common.NewAmount(0), "")
			tx, _ := ltransaction.NewNormalUTXOTransaction(utxo.GetAllUTXOsByPubKeyHash(accounts[0].GetPubKeyHash()).GetAllUtxos(), sendTxParam)
			utxo.UpdateUtxo(&tx)
			txs = append(txs, &tx)
		}

		b := block.NewBlock(txs, tailBlk, "")
		b.SetHash(lblock.CalculateHash(b))
		state := scState.NewScState(bc.GetUtxoCache())
		bc.AddBlockContextToTail(&BlockContext{Block: b, UtxoIndex: utxo, State: state})
	}
}

func GenerateMockBlockchain(size int) *Blockchain {
	//create a new block chain
	s := storage.NewRamStorage()

	addr := account.NewAddress("16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
	bc := CreateBlockchain(addr, s, nil, transactionpool.NewTransactionPool(nil, 128000), 100000)

	for i := 0; i < size; i++ {
		tailBlk, _ := bc.GetTailBlock()
		b := block.NewBlock([]*transaction.Transaction{core.MockTransaction()}, tailBlk, "16PencPNnF8CiSx2EBGEd1axhf7vuHCouj")
		b.SetHash(lblock.CalculateHash(b))
		bc.AddBlockContextToTail(PrepareBlockContext(bc, b))
	}
	return bc
}
