// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	instruments "project/features/instruments"

	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"
)

// InstrumentService is an autogenerated mock type for the InstrumentService type
type InstrumentService struct {
	mock.Mock
}

// AddInstrument provides a mock function with given fields: newData, token
func (_m *InstrumentService) AddInstrument(newData instruments.Instruments, token *jwt.Token) (*instruments.Instruments, error) {
	ret := _m.Called(newData, token)

	var r0 *instruments.Instruments
	var r1 error
	if rf, ok := ret.Get(0).(func(instruments.Instruments, *jwt.Token) (*instruments.Instruments, error)); ok {
		return rf(newData, token)
	}
	if rf, ok := ret.Get(0).(func(instruments.Instruments, *jwt.Token) *instruments.Instruments); ok {
		r0 = rf(newData, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*instruments.Instruments)
		}
	}

	if rf, ok := ret.Get(1).(func(instruments.Instruments, *jwt.Token) error); ok {
		r1 = rf(newData, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInstrument provides a mock function with given fields: id, token
func (_m *InstrumentService) DeleteInstrument(id string, token *jwt.Token) error {
	ret := _m.Called(id, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *jwt.Token) error); ok {
		r0 = rf(id, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllInstruments provides a mock function with given fields: page, pageSize
func (_m *InstrumentService) GetAllInstruments(page int, pageSize int) ([]map[string]interface{}, error) {
	ret := _m.Called(page, pageSize)

	var r0 []map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]map[string]interface{}, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []map[string]interface{}); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstrumentByID provides a mock function with given fields: id
func (_m *InstrumentService) GetInstrumentByID(id string) (*instruments.Instruments, error) {
	ret := _m.Called(id)

	var r0 *instruments.Instruments
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*instruments.Instruments, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *instruments.Instruments); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*instruments.Instruments)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateInstrument provides a mock function with given fields: id, newData, token
func (_m *InstrumentService) UpdateInstrument(id string, newData instruments.Instruments, token *jwt.Token) (*instruments.Instruments, error) {
	ret := _m.Called(id, newData, token)

	var r0 *instruments.Instruments
	var r1 error
	if rf, ok := ret.Get(0).(func(string, instruments.Instruments, *jwt.Token) (*instruments.Instruments, error)); ok {
		return rf(id, newData, token)
	}
	if rf, ok := ret.Get(0).(func(string, instruments.Instruments, *jwt.Token) *instruments.Instruments); ok {
		r0 = rf(id, newData, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*instruments.Instruments)
		}
	}

	if rf, ok := ret.Get(1).(func(string, instruments.Instruments, *jwt.Token) error); ok {
		r1 = rf(id, newData, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInstrumentService creates a new instance of InstrumentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInstrumentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *InstrumentService {
	mock := &InstrumentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
