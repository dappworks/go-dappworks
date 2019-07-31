// Code generated by mockery v1.0.0. DO NOT EDIT.

package core

import (
	"errors"

	"github.com/dappley/go-dappley/core/account"
	mock "github.com/stretchr/testify/mock"
)

// MockScEngine is an autogenerated mock type for the ScEngine type
type MockScEngine struct {
	mock.Mock
}

// DestroyEngine provides a mock function with given fields:
func (_m *MockScEngine) DestroyEngine() {
	_m.Called()
}

// Execute provides a mock function with given fields: function, args
func (_m *MockScEngine) Execute(function string, args string) (string, error) {
	ret := _m.Called(function, args)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(function, args)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0, nil
}

// GetGeneratedTXs provides a mock function with given fields:
func (_m *MockScEngine) GetGeneratedTXs() []*Transaction {
	ret := _m.Called()

	var r0 []*Transaction
	if rf, ok := ret.Get(0).(func() []*Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Transaction)
		}
	}

	return r0
}

// ImportContractAddr provides a mock function with given fields: contractAddr
func (_m *MockScEngine) ImportContractAddr(contractAddr account.Address) {
	_m.Called(contractAddr)
}

// ImportCurrBlockHeight provides a mock function with given fields: currBlkHeight
func (_m *MockScEngine) ImportCurrBlockHeight(currBlkHeight uint64) {
	_m.Called(currBlkHeight)
}

// ImportLocalStorage provides a mock function with given fields: state
func (_m *MockScEngine) ImportLocalStorage(state *ScState) {
	_m.Called(state)
}

// ImportNodeAddress provides a mock function with given fields: addr
func (_m *MockScEngine) ImportNodeAddress(addr account.Address) {
	_m.Called(addr)
}

// ImportPrevUtxos provides a mock function with given fields: utxos
func (_m *MockScEngine) ImportPrevUtxos(utxos []*UTXO) {
	_m.Called(utxos)
}

// ImportRewardStorage provides a mock function with given fields: rewards
func (_m *MockScEngine) ImportRewardStorage(rewards map[string]string) {
	_m.Called(rewards)
}

// ImportSeed provides a mock function with given fields: seed
func (_m *MockScEngine) ImportSeed(seed int64) {
	_m.Called(seed)
}

// ImportSourceCode provides a mock function with given fields: source
func (_m *MockScEngine) ImportSourceCode(source string) {
	_m.Called(source)
}

// ImportSourceTXID provides a mock function with given fields: txid
func (_m *MockScEngine) ImportSourceTXID(txid []byte) {
	_m.Called(txid)
}

// ImportTransaction provides a mock function with given fields: tx
func (_m *MockScEngine) ImportTransaction(tx *Transaction) {
	_m.Called(tx)
}

func (_m *MockScEngine) ImportContractCreateUTXO(utxo *UTXO) {
	_m.Called(utxo)
}

// ImportUTXOs provides a mock function with given fields: utxos
func (_m *MockScEngine) ImportUTXOs(utxos []*UTXO) {
	_m.Called(utxos)
}

// CheckContactSyntax provides a mock function with given fields: sourece
func (_m *MockScEngine) CheckContactSyntax(sourece string) error {
	ret := _m.Called(sourece)

	var r0 string
	r0 = ret.Get(0).(string)

	if r0 == "" {
		return errors.New("contract error syntax")
	}
	return nil
}

// SetExecutionLimits set execution limits of V8 Engine, prevent Halting Problem.
func (_m *MockScEngine) SetExecutionLimits(limitsOfExecutionInstructions, limitsOfTotalMemorySize uint64) error {
	return nil
}

// ExecutionInstructions returns the execution instructions
func (_m *MockScEngine) ExecutionInstructions() uint64 {
	return uint64(100)
}
