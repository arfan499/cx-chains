// Code generated by mockery v1.0.0. DO NOT EDIT.

package visor

import (
	cipher "github.com/SkycoinProject/cx-chains/src/cipher"
	coin "github.com/SkycoinProject/cx-chains/src/coin"

	dbutil "github.com/SkycoinProject/cx-chains/src/visor/dbutil"

	mock "github.com/stretchr/testify/mock"

	params "github.com/SkycoinProject/cx-chains/src/params"
)

// MockUnconfirmedTransactionPooler is an autogenerated mock type for the UnconfirmedTransactionPooler type
type MockUnconfirmedTransactionPooler struct {
	mock.Mock
}

// AllRawTransactions provides a mock function with given fields: tx
func (_m *MockUnconfirmedTransactionPooler) AllRawTransactions(tx *dbutil.Tx) (coin.Transactions, error) {
	ret := _m.Called(tx)

	var r0 coin.Transactions
	if rf, ok := ret.Get(0).(func(*dbutil.Tx) coin.Transactions); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coin.Transactions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterKnown provides a mock function with given fields: tx, txns
func (_m *MockUnconfirmedTransactionPooler) FilterKnown(tx *dbutil.Tx, txns []cipher.SHA256) ([]cipher.SHA256, error) {
	ret := _m.Called(tx, txns)

	var r0 []cipher.SHA256
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, []cipher.SHA256) []cipher.SHA256); ok {
		r0 = rf(tx, txns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.SHA256)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, []cipher.SHA256) error); ok {
		r1 = rf(tx, txns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForEach provides a mock function with given fields: tx, f
func (_m *MockUnconfirmedTransactionPooler) ForEach(tx *dbutil.Tx, f func(cipher.SHA256, UnconfirmedTransaction) error) error {
	ret := _m.Called(tx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, func(cipher.SHA256, UnconfirmedTransaction) error) error); ok {
		r0 = rf(tx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: tx, hash
func (_m *MockUnconfirmedTransactionPooler) Get(tx *dbutil.Tx, hash cipher.SHA256) (*UnconfirmedTransaction, error) {
	ret := _m.Called(tx, hash)

	var r0 *UnconfirmedTransaction
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, cipher.SHA256) *UnconfirmedTransaction); ok {
		r0 = rf(tx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UnconfirmedTransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, cipher.SHA256) error); ok {
		r1 = rf(tx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFiltered provides a mock function with given fields: tx, filter
func (_m *MockUnconfirmedTransactionPooler) GetFiltered(tx *dbutil.Tx, filter func(UnconfirmedTransaction) bool) ([]UnconfirmedTransaction, error) {
	ret := _m.Called(tx, filter)

	var r0 []UnconfirmedTransaction
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, func(UnconfirmedTransaction) bool) []UnconfirmedTransaction); ok {
		r0 = rf(tx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]UnconfirmedTransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, func(UnconfirmedTransaction) bool) error); ok {
		r1 = rf(tx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHashes provides a mock function with given fields: tx, filter
func (_m *MockUnconfirmedTransactionPooler) GetHashes(tx *dbutil.Tx, filter func(UnconfirmedTransaction) bool) ([]cipher.SHA256, error) {
	ret := _m.Called(tx, filter)

	var r0 []cipher.SHA256
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, func(UnconfirmedTransaction) bool) []cipher.SHA256); ok {
		r0 = rf(tx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.SHA256)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, func(UnconfirmedTransaction) bool) error); ok {
		r1 = rf(tx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIncomingOutputs provides a mock function with given fields: tx, bh
func (_m *MockUnconfirmedTransactionPooler) GetIncomingOutputs(tx *dbutil.Tx, bh coin.BlockHeader) (coin.UxArray, error) {
	ret := _m.Called(tx, bh)

	var r0 coin.UxArray
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, coin.BlockHeader) coin.UxArray); ok {
		r0 = rf(tx, bh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coin.UxArray)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, coin.BlockHeader) error); ok {
		r1 = rf(tx, bh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKnown provides a mock function with given fields: tx, txns
func (_m *MockUnconfirmedTransactionPooler) GetKnown(tx *dbutil.Tx, txns []cipher.SHA256) (coin.Transactions, error) {
	ret := _m.Called(tx, txns)

	var r0 coin.Transactions
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, []cipher.SHA256) coin.Transactions); ok {
		r0 = rf(tx, txns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coin.Transactions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, []cipher.SHA256) error); ok {
		r1 = rf(tx, txns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnspentsOfAddr provides a mock function with given fields: tx, addr
func (_m *MockUnconfirmedTransactionPooler) GetUnspentsOfAddr(tx *dbutil.Tx, addr cipher.Address) (coin.UxArray, error) {
	ret := _m.Called(tx, addr)

	var r0 coin.UxArray
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, cipher.Address) coin.UxArray); ok {
		r0 = rf(tx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coin.UxArray)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, cipher.Address) error); ok {
		r1 = rf(tx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InjectTransaction provides a mock function with given fields: tx, bc, t, distParams, verifyParams
func (_m *MockUnconfirmedTransactionPooler) InjectTransaction(tx *dbutil.Tx, bc Blockchainer, t coin.Transaction, distParams params.Distribution, verifyParams params.VerifyTxn) (bool, *ErrTxnViolatesSoftConstraint, error) {
	ret := _m.Called(tx, bc, t, distParams, verifyParams)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, Blockchainer, coin.Transaction, params.Distribution, params.VerifyTxn) bool); ok {
		r0 = rf(tx, bc, t, distParams, verifyParams)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *ErrTxnViolatesSoftConstraint
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, Blockchainer, coin.Transaction, params.Distribution, params.VerifyTxn) *ErrTxnViolatesSoftConstraint); ok {
		r1 = rf(tx, bc, t, distParams, verifyParams)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ErrTxnViolatesSoftConstraint)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*dbutil.Tx, Blockchainer, coin.Transaction, params.Distribution, params.VerifyTxn) error); ok {
		r2 = rf(tx, bc, t, distParams, verifyParams)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Len provides a mock function with given fields: tx
func (_m *MockUnconfirmedTransactionPooler) Len(tx *dbutil.Tx) (uint64, error) {
	ret := _m.Called(tx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*dbutil.Tx) uint64); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx) error); ok {
		r1 = rf(tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecvOfAddresses provides a mock function with given fields: tx, bh, addrs
func (_m *MockUnconfirmedTransactionPooler) RecvOfAddresses(tx *dbutil.Tx, bh coin.BlockHeader, addrs []cipher.Address) (coin.AddressUxOuts, error) {
	ret := _m.Called(tx, bh, addrs)

	var r0 coin.AddressUxOuts
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, coin.BlockHeader, []cipher.Address) coin.AddressUxOuts); ok {
		r0 = rf(tx, bh, addrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coin.AddressUxOuts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, coin.BlockHeader, []cipher.Address) error); ok {
		r1 = rf(tx, bh, addrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: tx, bc, distParams, verifyParams
func (_m *MockUnconfirmedTransactionPooler) Refresh(tx *dbutil.Tx, bc Blockchainer, distParams params.Distribution, verifyParams params.VerifyTxn) ([]cipher.SHA256, error) {
	ret := _m.Called(tx, bc, distParams, verifyParams)

	var r0 []cipher.SHA256
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, Blockchainer, params.Distribution, params.VerifyTxn) []cipher.SHA256); ok {
		r0 = rf(tx, bc, distParams, verifyParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.SHA256)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, Blockchainer, params.Distribution, params.VerifyTxn) error); ok {
		r1 = rf(tx, bc, distParams, verifyParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveInvalid provides a mock function with given fields: tx, bc
func (_m *MockUnconfirmedTransactionPooler) RemoveInvalid(tx *dbutil.Tx, bc Blockchainer) ([]cipher.SHA256, error) {
	ret := _m.Called(tx, bc)

	var r0 []cipher.SHA256
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, Blockchainer) []cipher.SHA256); ok {
		r0 = rf(tx, bc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.SHA256)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dbutil.Tx, Blockchainer) error); ok {
		r1 = rf(tx, bc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTransactions provides a mock function with given fields: tx, txns
func (_m *MockUnconfirmedTransactionPooler) RemoveTransactions(tx *dbutil.Tx, txns []cipher.SHA256) error {
	ret := _m.Called(tx, txns)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, []cipher.SHA256) error); ok {
		r0 = rf(tx, txns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTransactionsAnnounced provides a mock function with given fields: tx, hashes
func (_m *MockUnconfirmedTransactionPooler) SetTransactionsAnnounced(tx *dbutil.Tx, hashes map[cipher.SHA256]int64) error {
	ret := _m.Called(tx, hashes)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbutil.Tx, map[cipher.SHA256]int64) error); ok {
		r0 = rf(tx, hashes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
