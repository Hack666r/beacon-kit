// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// BeaconBlock is an autogenerated mock type for the BeaconBlock type
type BeaconBlock[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	mock.Mock
}

type BeaconBlock_Expecter[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	mock *mock.Mock
}

func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) EXPECT() *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]{mock: &_m.Mock}
}

// GetBody provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) GetBody() BeaconBlockBodyT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 BeaconBlockBodyT
	if rf, ok := ret.Get(0).(func() BeaconBlockBodyT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BeaconBlockBodyT)
	}

	return r0
}

// BeaconBlock_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type BeaconBlock_GetBody_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) GetBody() *BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("GetBody")}
}

func (_c *BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 BeaconBlockBodyT) *BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() BeaconBlockBodyT) *BeaconBlock_GetBody_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetHeader provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) GetHeader() BeaconBlockHeaderT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHeader")
	}

	var r0 BeaconBlockHeaderT
	if rf, ok := ret.Get(0).(func() BeaconBlockHeaderT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BeaconBlockHeaderT)
	}

	return r0
}

// BeaconBlock_GetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeader'
type BeaconBlock_GetHeader_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// GetHeader is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) GetHeader() *BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("GetHeader")}
}

func (_c *BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 BeaconBlockHeaderT) *BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() BeaconBlockHeaderT) *BeaconBlock_GetHeader_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetParentBlockRoot provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) GetParentBlockRoot() bytes.B32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetParentBlockRoot")
	}

	var r0 bytes.B32
	if rf, ok := ret.Get(0).(func() bytes.B32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.B32)
		}
	}

	return r0
}

// BeaconBlock_GetParentBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParentBlockRoot'
type BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// GetParentBlockRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) GetParentBlockRoot() *BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("GetParentBlockRoot")}
}

func (_c *BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 bytes.B32) *BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() bytes.B32) *BeaconBlock_GetParentBlockRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetProposerIndex provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) GetProposerIndex() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProposerIndex")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetProposerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposerIndex'
type BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// GetProposerIndex is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) GetProposerIndex() *BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("GetProposerIndex")}
}

func (_c *BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 math.U64) *BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() math.U64) *BeaconBlock_GetProposerIndex_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlock_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type BeaconBlock_GetSlot_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) GetSlot() *BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("GetSlot")}
}

func (_c *BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 math.U64) *BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() math.U64) *BeaconBlock_GetSlot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetStateRoot provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) GetStateRoot() bytes.B32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStateRoot")
	}

	var r0 bytes.B32
	if rf, ok := ret.Get(0).(func() bytes.B32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.B32)
		}
	}

	return r0
}

// BeaconBlock_GetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStateRoot'
type BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// GetStateRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) GetStateRoot() *BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("GetStateRoot")}
}

func (_c *BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 bytes.B32) *BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() bytes.B32) *BeaconBlock_GetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// HashTreeRoot provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) HashTreeRoot() ([32]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRoot")
	}

	var r0 [32]byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([32]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() [32]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlock_HashTreeRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRoot'
type BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// HashTreeRoot is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) HashTreeRoot() *BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("HashTreeRoot")}
}

func (_c *BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 [32]byte, _a1 error) *BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() ([32]byte, error)) *BeaconBlock_HashTreeRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// IsNil provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) IsNil() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsNil")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BeaconBlock_IsNil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNil'
type BeaconBlock_IsNil_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// IsNil is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) IsNil() *BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("IsNil")}
}

func (_c *BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 bool) *BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() bool) *BeaconBlock_IsNil_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZ provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) MarshalSSZ() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZ")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlock_MarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZ'
type BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// MarshalSSZ is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) MarshalSSZ() *BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("MarshalSSZ")}
}

func (_c *BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 []byte, _a1 error) *BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() ([]byte, error)) *BeaconBlock_MarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// SetStateRoot provides a mock function with given fields: _a0
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) SetStateRoot(_a0 bytes.B32) {
	_m.Called(_a0)
}

// BeaconBlock_SetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStateRoot'
type BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// SetStateRoot is a helper method to define mock.On call
//   - _a0 bytes.B32
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) SetStateRoot(_a0 interface{}) *BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("SetStateRoot", _a0)}
}

func (_c *BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func(_a0 bytes.B32)) *BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bytes.B32))
	})
	return _c
}

func (_c *BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return() *BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return()
	return _c
}

func (_c *BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func(bytes.B32)) *BeaconBlock_SetStateRoot_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// UnmarshalSSZ provides a mock function with given fields: _a0
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) UnmarshalSSZ(_a0 []byte) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalSSZ")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeaconBlock_UnmarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalSSZ'
type BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// UnmarshalSSZ is a helper method to define mock.On call
//   - _a0 []byte
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) UnmarshalSSZ(_a0 interface{}) *BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("UnmarshalSSZ", _a0)}
}

func (_c *BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func(_a0 []byte)) *BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 error) *BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func([]byte) error) *BeaconBlock_UnmarshalSSZ_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// Version provides a mock function with given fields:
func (_m *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]) Version() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// BeaconBlock_Version_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Version'
type BeaconBlock_Version_Call[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}] struct {
	*mock.Call
}

// Version is a helper method to define mock.On call
func (_e *BeaconBlock_Expecter[BeaconBlockBodyT, BeaconBlockHeaderT]) Version() *BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	return &BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT]{Call: _e.mock.On("Version")}
}

func (_c *BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Run(run func()) *BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) Return(_a0 uint32) *BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT]) RunAndReturn(run func() uint32) *BeaconBlock_Version_Call[BeaconBlockBodyT, BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// NewBeaconBlock creates a new instance of BeaconBlock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconBlock[BeaconBlockBodyT interface{}, BeaconBlockHeaderT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT] {
	mock := &BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}