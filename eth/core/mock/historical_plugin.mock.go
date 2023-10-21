// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	ethereumcoretypes "github.com/ethereum/go-ethereum/core/types"
	"pkg.berachain.dev/polaris/eth/core"
	ethcoretypes "pkg.berachain.dev/polaris/eth/core/types"
	"sync"
)

// Ensure, that HistoricalPluginMock does implement core.HistoricalPlugin.
// If this is not the case, regenerate this file with moq.
var _ core.HistoricalPlugin = &HistoricalPluginMock{}

// HistoricalPluginMock is a mock implementation of core.HistoricalPlugin.
//
//	func TestSomethingThatUsesHistoricalPlugin(t *testing.T) {
//
//		// make and configure a mocked core.HistoricalPlugin
//		mockedHistoricalPlugin := &HistoricalPluginMock{
//			GetBlockByHashFunc: func(hash common.Hash) (*ethereumcoretypes.Block, error) {
//				panic("mock out the GetBlockByHash method")
//			},
//			GetBlockByNumberFunc: func(v uint64) (*ethereumcoretypes.Block, error) {
//				panic("mock out the GetBlockByNumber method")
//			},
//			GetReceiptsByHashFunc: func(hash common.Hash) (ethereumcoretypes.Receipts, error) {
//				panic("mock out the GetReceiptsByHash method")
//			},
//			GetTransactionByHashFunc: func(hash common.Hash) (*ethcoretypes.TxLookupEntry, error) {
//				panic("mock out the GetTransactionByHash method")
//			},
//			PrepareFunc: func(contextMoqParam context.Context)  {
//				panic("mock out the Prepare method")
//			},
//			StoreBlockFunc: func(block *ethereumcoretypes.Block) error {
//				panic("mock out the StoreBlock method")
//			},
//			StoreReceiptsFunc: func(hash common.Hash, receipts ethereumcoretypes.Receipts) error {
//				panic("mock out the StoreReceipts method")
//			},
//			StoreTransactionsFunc: func(v uint64, hash common.Hash, transactions ethereumcoretypes.Transactions) error {
//				panic("mock out the StoreTransactions method")
//			},
//		}
//
//		// use mockedHistoricalPlugin in code that requires core.HistoricalPlugin
//		// and then make assertions.
//
//	}
type HistoricalPluginMock struct {
	// GetBlockByHashFunc mocks the GetBlockByHash method.
	GetBlockByHashFunc func(hash common.Hash) (*ethereumcoretypes.Block, error)

	// GetBlockByNumberFunc mocks the GetBlockByNumber method.
	GetBlockByNumberFunc func(v uint64) (*ethereumcoretypes.Block, error)

	// GetReceiptsByHashFunc mocks the GetReceiptsByHash method.
	GetReceiptsByHashFunc func(hash common.Hash) (ethereumcoretypes.Receipts, error)

	// GetTransactionByHashFunc mocks the GetTransactionByHash method.
	GetTransactionByHashFunc func(hash common.Hash) (*ethcoretypes.TxLookupEntry, error)

	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(contextMoqParam context.Context)

	// StoreBlockFunc mocks the StoreBlock method.
	StoreBlockFunc func(block *ethereumcoretypes.Block) error

	// StoreReceiptsFunc mocks the StoreReceipts method.
	StoreReceiptsFunc func(hash common.Hash, receipts ethereumcoretypes.Receipts) error

	// StoreTransactionsFunc mocks the StoreTransactions method.
	StoreTransactionsFunc func(v uint64, hash common.Hash, transactions ethereumcoretypes.Transactions) error

	// calls tracks calls to the methods.
	calls struct {
		// GetBlockByHash holds details about calls to the GetBlockByHash method.
		GetBlockByHash []struct {
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// GetBlockByNumber holds details about calls to the GetBlockByNumber method.
		GetBlockByNumber []struct {
			// V is the v argument value.
			V uint64
		}
		// GetReceiptsByHash holds details about calls to the GetReceiptsByHash method.
		GetReceiptsByHash []struct {
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// GetTransactionByHash holds details about calls to the GetTransactionByHash method.
		GetTransactionByHash []struct {
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// StoreBlock holds details about calls to the StoreBlock method.
		StoreBlock []struct {
			// Block is the block argument value.
			Block *ethereumcoretypes.Block
		}
		// StoreReceipts holds details about calls to the StoreReceipts method.
		StoreReceipts []struct {
			// Hash is the hash argument value.
			Hash common.Hash
			// Receipts is the receipts argument value.
			Receipts ethereumcoretypes.Receipts
		}
		// StoreTransactions holds details about calls to the StoreTransactions method.
		StoreTransactions []struct {
			// V is the v argument value.
			V uint64
			// Hash is the hash argument value.
			Hash common.Hash
			// Transactions is the transactions argument value.
			Transactions ethereumcoretypes.Transactions
		}
	}
	lockGetBlockByHash       sync.RWMutex
	lockGetBlockByNumber     sync.RWMutex
	lockGetReceiptsByHash    sync.RWMutex
	lockGetTransactionByHash sync.RWMutex
	lockPrepare              sync.RWMutex
	lockStoreBlock           sync.RWMutex
	lockStoreReceipts        sync.RWMutex
	lockStoreTransactions    sync.RWMutex
}

// GetBlockByHash calls GetBlockByHashFunc.
func (mock *HistoricalPluginMock) GetBlockByHash(hash common.Hash) (*ethereumcoretypes.Block, error) {
	if mock.GetBlockByHashFunc == nil {
		panic("HistoricalPluginMock.GetBlockByHashFunc: method is nil but HistoricalPlugin.GetBlockByHash was just called")
	}
	callInfo := struct {
		Hash common.Hash
	}{
		Hash: hash,
	}
	mock.lockGetBlockByHash.Lock()
	mock.calls.GetBlockByHash = append(mock.calls.GetBlockByHash, callInfo)
	mock.lockGetBlockByHash.Unlock()
	return mock.GetBlockByHashFunc(hash)
}

// GetBlockByHashCalls gets all the calls that were made to GetBlockByHash.
// Check the length with:
//
//	len(mockedHistoricalPlugin.GetBlockByHashCalls())
func (mock *HistoricalPluginMock) GetBlockByHashCalls() []struct {
	Hash common.Hash
} {
	var calls []struct {
		Hash common.Hash
	}
	mock.lockGetBlockByHash.RLock()
	calls = mock.calls.GetBlockByHash
	mock.lockGetBlockByHash.RUnlock()
	return calls
}

// GetBlockByNumber calls GetBlockByNumberFunc.
func (mock *HistoricalPluginMock) GetBlockByNumber(v uint64) (*ethereumcoretypes.Block, error) {
	if mock.GetBlockByNumberFunc == nil {
		panic("HistoricalPluginMock.GetBlockByNumberFunc: method is nil but HistoricalPlugin.GetBlockByNumber was just called")
	}
	callInfo := struct {
		V uint64
	}{
		V: v,
	}
	mock.lockGetBlockByNumber.Lock()
	mock.calls.GetBlockByNumber = append(mock.calls.GetBlockByNumber, callInfo)
	mock.lockGetBlockByNumber.Unlock()
	return mock.GetBlockByNumberFunc(v)
}

// GetBlockByNumberCalls gets all the calls that were made to GetBlockByNumber.
// Check the length with:
//
//	len(mockedHistoricalPlugin.GetBlockByNumberCalls())
func (mock *HistoricalPluginMock) GetBlockByNumberCalls() []struct {
	V uint64
} {
	var calls []struct {
		V uint64
	}
	mock.lockGetBlockByNumber.RLock()
	calls = mock.calls.GetBlockByNumber
	mock.lockGetBlockByNumber.RUnlock()
	return calls
}

// GetReceiptsByHash calls GetReceiptsByHashFunc.
func (mock *HistoricalPluginMock) GetReceiptsByHash(hash common.Hash) (ethereumcoretypes.Receipts, error) {
	if mock.GetReceiptsByHashFunc == nil {
		panic("HistoricalPluginMock.GetReceiptsByHashFunc: method is nil but HistoricalPlugin.GetReceiptsByHash was just called")
	}
	callInfo := struct {
		Hash common.Hash
	}{
		Hash: hash,
	}
	mock.lockGetReceiptsByHash.Lock()
	mock.calls.GetReceiptsByHash = append(mock.calls.GetReceiptsByHash, callInfo)
	mock.lockGetReceiptsByHash.Unlock()
	return mock.GetReceiptsByHashFunc(hash)
}

// GetReceiptsByHashCalls gets all the calls that were made to GetReceiptsByHash.
// Check the length with:
//
//	len(mockedHistoricalPlugin.GetReceiptsByHashCalls())
func (mock *HistoricalPluginMock) GetReceiptsByHashCalls() []struct {
	Hash common.Hash
} {
	var calls []struct {
		Hash common.Hash
	}
	mock.lockGetReceiptsByHash.RLock()
	calls = mock.calls.GetReceiptsByHash
	mock.lockGetReceiptsByHash.RUnlock()
	return calls
}

// GetTransactionByHash calls GetTransactionByHashFunc.
func (mock *HistoricalPluginMock) GetTransactionByHash(hash common.Hash) (*ethcoretypes.TxLookupEntry, error) {
	if mock.GetTransactionByHashFunc == nil {
		panic("HistoricalPluginMock.GetTransactionByHashFunc: method is nil but HistoricalPlugin.GetTransactionByHash was just called")
	}
	callInfo := struct {
		Hash common.Hash
	}{
		Hash: hash,
	}
	mock.lockGetTransactionByHash.Lock()
	mock.calls.GetTransactionByHash = append(mock.calls.GetTransactionByHash, callInfo)
	mock.lockGetTransactionByHash.Unlock()
	return mock.GetTransactionByHashFunc(hash)
}

// GetTransactionByHashCalls gets all the calls that were made to GetTransactionByHash.
// Check the length with:
//
//	len(mockedHistoricalPlugin.GetTransactionByHashCalls())
func (mock *HistoricalPluginMock) GetTransactionByHashCalls() []struct {
	Hash common.Hash
} {
	var calls []struct {
		Hash common.Hash
	}
	mock.lockGetTransactionByHash.RLock()
	calls = mock.calls.GetTransactionByHash
	mock.lockGetTransactionByHash.RUnlock()
	return calls
}

// Prepare calls PrepareFunc.
func (mock *HistoricalPluginMock) Prepare(contextMoqParam context.Context) {
	if mock.PrepareFunc == nil {
		panic("HistoricalPluginMock.PrepareFunc: method is nil but HistoricalPlugin.Prepare was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockPrepare.Lock()
	mock.calls.Prepare = append(mock.calls.Prepare, callInfo)
	mock.lockPrepare.Unlock()
	mock.PrepareFunc(contextMoqParam)
}

// PrepareCalls gets all the calls that were made to Prepare.
// Check the length with:
//
//	len(mockedHistoricalPlugin.PrepareCalls())
func (mock *HistoricalPluginMock) PrepareCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockPrepare.RLock()
	calls = mock.calls.Prepare
	mock.lockPrepare.RUnlock()
	return calls
}

// StoreBlock calls StoreBlockFunc.
func (mock *HistoricalPluginMock) StoreBlock(block *ethereumcoretypes.Block) error {
	if mock.StoreBlockFunc == nil {
		panic("HistoricalPluginMock.StoreBlockFunc: method is nil but HistoricalPlugin.StoreBlock was just called")
	}
	callInfo := struct {
		Block *ethereumcoretypes.Block
	}{
		Block: block,
	}
	mock.lockStoreBlock.Lock()
	mock.calls.StoreBlock = append(mock.calls.StoreBlock, callInfo)
	mock.lockStoreBlock.Unlock()
	return mock.StoreBlockFunc(block)
}

// StoreBlockCalls gets all the calls that were made to StoreBlock.
// Check the length with:
//
//	len(mockedHistoricalPlugin.StoreBlockCalls())
func (mock *HistoricalPluginMock) StoreBlockCalls() []struct {
	Block *ethereumcoretypes.Block
} {
	var calls []struct {
		Block *ethereumcoretypes.Block
	}
	mock.lockStoreBlock.RLock()
	calls = mock.calls.StoreBlock
	mock.lockStoreBlock.RUnlock()
	return calls
}

// StoreReceipts calls StoreReceiptsFunc.
func (mock *HistoricalPluginMock) StoreReceipts(hash common.Hash, receipts ethereumcoretypes.Receipts) error {
	if mock.StoreReceiptsFunc == nil {
		panic("HistoricalPluginMock.StoreReceiptsFunc: method is nil but HistoricalPlugin.StoreReceipts was just called")
	}
	callInfo := struct {
		Hash     common.Hash
		Receipts ethereumcoretypes.Receipts
	}{
		Hash:     hash,
		Receipts: receipts,
	}
	mock.lockStoreReceipts.Lock()
	mock.calls.StoreReceipts = append(mock.calls.StoreReceipts, callInfo)
	mock.lockStoreReceipts.Unlock()
	return mock.StoreReceiptsFunc(hash, receipts)
}

// StoreReceiptsCalls gets all the calls that were made to StoreReceipts.
// Check the length with:
//
//	len(mockedHistoricalPlugin.StoreReceiptsCalls())
func (mock *HistoricalPluginMock) StoreReceiptsCalls() []struct {
	Hash     common.Hash
	Receipts ethereumcoretypes.Receipts
} {
	var calls []struct {
		Hash     common.Hash
		Receipts ethereumcoretypes.Receipts
	}
	mock.lockStoreReceipts.RLock()
	calls = mock.calls.StoreReceipts
	mock.lockStoreReceipts.RUnlock()
	return calls
}

// StoreTransactions calls StoreTransactionsFunc.
func (mock *HistoricalPluginMock) StoreTransactions(v uint64, hash common.Hash, transactions ethereumcoretypes.Transactions) error {
	if mock.StoreTransactionsFunc == nil {
		panic("HistoricalPluginMock.StoreTransactionsFunc: method is nil but HistoricalPlugin.StoreTransactions was just called")
	}
	callInfo := struct {
		V            uint64
		Hash         common.Hash
		Transactions ethereumcoretypes.Transactions
	}{
		V:            v,
		Hash:         hash,
		Transactions: transactions,
	}
	mock.lockStoreTransactions.Lock()
	mock.calls.StoreTransactions = append(mock.calls.StoreTransactions, callInfo)
	mock.lockStoreTransactions.Unlock()
	return mock.StoreTransactionsFunc(v, hash, transactions)
}

// StoreTransactionsCalls gets all the calls that were made to StoreTransactions.
// Check the length with:
//
//	len(mockedHistoricalPlugin.StoreTransactionsCalls())
func (mock *HistoricalPluginMock) StoreTransactionsCalls() []struct {
	V            uint64
	Hash         common.Hash
	Transactions ethereumcoretypes.Transactions
} {
	var calls []struct {
		V            uint64
		Hash         common.Hash
		Transactions ethereumcoretypes.Transactions
	}
	mock.lockStoreTransactions.RLock()
	calls = mock.calls.StoreTransactions
	mock.lockStoreTransactions.RUnlock()
	return calls
}