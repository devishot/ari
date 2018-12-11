// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import ari "github.com/devishot/ari"
import mock "github.com/stretchr/testify/mock"

// DeviceState is an autogenerated mock type for the DeviceState type
type DeviceState struct {
	mock.Mock
}

// Data provides a mock function with given fields: key
func (_m *DeviceState) Data(key *ari.Key) (*ari.DeviceStateData, error) {
	ret := _m.Called(key)

	var r0 *ari.DeviceStateData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.DeviceStateData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.DeviceStateData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: key
func (_m *DeviceState) Delete(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *DeviceState) Get(key *ari.Key) *ari.DeviceStateHandle {
	ret := _m.Called(key)

	var r0 *ari.DeviceStateHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.DeviceStateHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.DeviceStateHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields: filter
func (_m *DeviceState) List(filter *ari.Key) ([]*ari.Key, error) {
	ret := _m.Called(filter)

	var r0 []*ari.Key
	if rf, ok := ret.Get(0).(func(*ari.Key) []*ari.Key); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ari.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: key, state
func (_m *DeviceState) Update(key *ari.Key, state string) error {
	ret := _m.Called(key, state)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string) error); ok {
		r0 = rf(key, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
