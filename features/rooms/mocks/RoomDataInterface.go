// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	instruments "project/features/instruments"

	mock "github.com/stretchr/testify/mock"

	rooms "project/features/rooms"
)

// RoomDataInterface is an autogenerated mock type for the RoomDataInterface type
type RoomDataInterface struct {
	mock.Mock
}

// AddRoom provides a mock function with given fields: newRoom
func (_m *RoomDataInterface) AddRoom(newRoom rooms.Rooms) (*rooms.Rooms, error) {
	ret := _m.Called(newRoom)

	var r0 *rooms.Rooms
	var r1 error
	if rf, ok := ret.Get(0).(func(rooms.Rooms) (*rooms.Rooms, error)); ok {
		return rf(newRoom)
	}
	if rf, ok := ret.Get(0).(func(rooms.Rooms) *rooms.Rooms); ok {
		r0 = rf(newRoom)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rooms.Rooms)
		}
	}

	if rf, ok := ret.Get(1).(func(rooms.Rooms) error); ok {
		r1 = rf(newRoom)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddRoomInstrument provides a mock function with given fields: roomId, instrumentData
func (_m *RoomDataInterface) AddRoomInstrument(roomId string, instrumentData instruments.RoomInstrument) (interface{}, error) {
	ret := _m.Called(roomId, instrumentData)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, instruments.RoomInstrument) (interface{}, error)); ok {
		return rf(roomId, instrumentData)
	}
	if rf, ok := ret.Get(0).(func(string, instruments.RoomInstrument) interface{}); ok {
		r0 = rf(roomId, instrumentData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, instruments.RoomInstrument) error); ok {
		r1 = rf(roomId, instrumentData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRoom provides a mock function with given fields: roomID
func (_m *RoomDataInterface) DeleteRoom(roomID string) error {
	ret := _m.Called(roomID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(roomID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterRoomByPrice provides a mock function with given fields: price
func (_m *RoomDataInterface) FilterRoomByPrice(price int) (map[string]map[string]interface{}, error) {
	ret := _m.Called(price)

	var r0 map[string]map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (map[string]map[string]interface{}, error)); ok {
		return rf(price)
	}
	if rf, ok := ret.Get(0).(func(int) map[string]map[string]interface{}); ok {
		r0 = rf(price)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllRooms provides a mock function with given fields:
func (_m *RoomDataInterface) GetAllRooms() (map[string]map[string]interface{}, error) {
	ret := _m.Called()

	var r0 map[string]map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookedRooms provides a mock function with given fields:
func (_m *RoomDataInterface) GetBookedRooms() (map[string]map[string]interface{}, error) {
	ret := _m.Called()

	var r0 map[string]map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]map[string]interface{}, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecommendation provides a mock function with given fields: genre1, genre2, message
func (_m *RoomDataInterface) GetRecommendation(genre1 string, genre2 string, message string) (string, error) {
	ret := _m.Called(genre1, genre2, message)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(genre1, genre2, message)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(genre1, genre2, message)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(genre1, genre2, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoomByID provides a mock function with given fields: roomID
func (_m *RoomDataInterface) GetRoomByID(roomID string) (*rooms.Rooms, error) {
	ret := _m.Called(roomID)

	var r0 *rooms.Rooms
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*rooms.Rooms, error)); ok {
		return rf(roomID)
	}
	if rf, ok := ret.Get(0).(func(string) *rooms.Rooms); ok {
		r0 = rf(roomID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rooms.Rooms)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roomID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRoomExist provides a mock function with given fields: roomID
func (_m *RoomDataInterface) IsRoomExist(roomID string) bool {
	ret := _m.Called(roomID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(roomID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UpdateRoom provides a mock function with given fields: roomID, updatedRoom
func (_m *RoomDataInterface) UpdateRoom(roomID string, updatedRoom rooms.Rooms) (*rooms.Rooms, error) {
	ret := _m.Called(roomID, updatedRoom)

	var r0 *rooms.Rooms
	var r1 error
	if rf, ok := ret.Get(0).(func(string, rooms.Rooms) (*rooms.Rooms, error)); ok {
		return rf(roomID, updatedRoom)
	}
	if rf, ok := ret.Get(0).(func(string, rooms.Rooms) *rooms.Rooms); ok {
		r0 = rf(roomID, updatedRoom)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rooms.Rooms)
		}
	}

	if rf, ok := ret.Get(1).(func(string, rooms.Rooms) error); ok {
		r1 = rf(roomID, updatedRoom)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRoomDataInterface creates a new instance of RoomDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoomDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RoomDataInterface {
	mock := &RoomDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
