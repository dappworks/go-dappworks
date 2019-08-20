// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//
package rpc

import (
	"context"
	"encoding/hex"

	"github.com/dappley/go-dappley/core/transaction"
	"github.com/dappley/go-dappley/logic/lblockchain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dappley/go-dappley/common"
	"github.com/dappley/go-dappley/core/account"
	"github.com/dappley/go-dappley/logic"
	"github.com/dappley/go-dappley/wallet"
	"github.com/dappley/go-dappley/network"
	networkpb "github.com/dappley/go-dappley/network/pb"
	rpcpb "github.com/dappley/go-dappley/rpc/pb"
)

type AdminRpcService struct {
	bm   *lblockchain.BlockchainManager
	node *network.Node
}

func (adminRpcService *AdminRpcService) RpcAddPeer(ctx context.Context, in *rpcpb.AddPeerRequest) (*rpcpb.AddPeerResponse, error) {
	err := adminRpcService.node.GetNetwork().ConnectToSeedByString(in.GetFullAddress())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &rpcpb.AddPeerResponse{}, nil
}

func (adminRpcService *AdminRpcService) RpcAddProducer(ctx context.Context, in *rpcpb.AddProducerRequest) (*rpcpb.AddProducerResponse, error) {
	if len(in.GetAddress()) == 0 || !account.NewAddress(in.GetAddress()).IsValid() {
		return nil, status.Error(codes.InvalidArgument, account.ErrInvalidAddress.Error())
	}
	err := adminRpcService.bm.Getblockchain().GetConsensus().AddProducer(in.GetAddress())
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &rpcpb.AddProducerResponse{}, nil
}

func (adminRpcService *AdminRpcService) RpcGetPeerInfo(ctx context.Context, in *rpcpb.GetPeerInfoRequest) (*rpcpb.GetPeerInfoResponse, error) {

	return &rpcpb.GetPeerInfoResponse{
		PeerList: getPeerInfo(adminRpcService.node),
	}, nil
}

//unlock the account through rpc service
func (adminRpcService *AdminRpcService) RpcUnlockAccount(ctx context.Context, in *rpcpb.UnlockAccountRequest) (*rpcpb.UnlockAccountResponse, error) {
	err := logic.SetUnLockAccount()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &rpcpb.UnlockAccountResponse{}, nil
}

func (adminRpcService *AdminRpcService) RpcSendFromMiner(ctx context.Context, in *rpcpb.SendFromMinerRequest) (*rpcpb.SendFromMinerResponse, error) {
	sendToAddress := account.NewAddress(in.GetTo())
	sendAmount := common.NewAmountFromBytes(in.GetAmount())
	if sendAmount.Validate() != nil || sendAmount.IsZero() {
		return nil, status.Error(codes.InvalidArgument, logic.ErrInvalidAmount.Error())
	}

	_, _, err := logic.SendFromMiner(sendToAddress, sendAmount, adminRpcService.bm.Getblockchain())
	if err != nil {
		switch err {
		case logic.ErrInvalidSenderAddress, logic.ErrInvalidRcverAddress, logic.ErrInvalidAmount:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case transaction.ErrInsufficientFund:
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}
	return &rpcpb.SendFromMinerResponse{}, nil
}

func (adminRpcService *AdminRpcService) RpcSend(ctx context.Context, in *rpcpb.SendRequest) (*rpcpb.SendResponse, error) {
	sendFromAddress := account.NewAddress(in.GetFrom())
	sendToAddress := account.NewAddress(in.GetTo())
	sendAmount := common.NewAmountFromBytes(in.GetAmount())
	tip := common.NewAmountFromBytes(in.GetTip())
	gasLimit := common.NewAmountFromBytes(in.GetGasLimit())
	gasPrice := common.NewAmountFromBytes(in.GetGasPrice())

	if sendAmount.Validate() != nil || sendAmount.IsZero() {
		return nil, status.Error(codes.InvalidArgument, transaction.ErrInvalidAmount.Error())
	}
	path := in.GetAccountPath()
	if len(path) == 0 {
		path = wallet.GetAccountFilePath()
	}

	am, err := logic.GetAccountManager(path)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	senderAccount := am.GetAccountByAddress(sendFromAddress)
	if senderAccount == nil || senderAccount.GetKeyPair() == nil {
		return nil, status.Error(codes.NotFound, wallet.ErrAddressNotFound.Error())
	}

	txHash, scAddress, err := logic.Send(senderAccount, sendToAddress, sendAmount, tip, gasLimit, gasPrice, in.GetData(),
		adminRpcService.bm.Getblockchain())

	txHashStr := hex.EncodeToString(txHash)
	if err != nil {
		switch err {
		case logic.ErrInvalidSenderAddress, logic.ErrInvalidRcverAddress, logic.ErrInvalidAmount:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case transaction.ErrInsufficientFund:
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	resp := &rpcpb.SendResponse{Txid: txHashStr}
	if scAddress != "" {
		resp.ContractAddress = scAddress
	}

	return resp, nil
}

func (adminRpcService *AdminRpcService) IsPrivate() bool { return true }

func getPeerInfo(node *network.Node) []*networkpb.PeerInfo {
	peers := node.GetNetwork().GetConnectedPeers()

	var peerPbs []*networkpb.PeerInfo
	for _, peerInfo := range peers {
		peerPbs = append(peerPbs, peerInfo.ToProto().(*networkpb.PeerInfo))
	}

	return peerPbs
}
