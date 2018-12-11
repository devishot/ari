// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import ari "github.com/devishot/ari"
import mock "github.com/stretchr/testify/mock"

// Playback is an autogenerated mock type for the Playback type
type Playback struct {
	mock.Mock
}

// Control provides a mock function with given fields: key, op
func (_m *Playback) Control(key *ari.Key, op string) error {
	ret := _m.Called(key, op)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string) error); ok {
		r0 = rf(key, op)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Data provides a mock function with given fields: key
func (_m *Playback) Data(key *ari.Key) (*ari.PlaybackData, error) {
	ret := _m.Called(key)

	var r0 *ari.PlaybackData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.PlaybackData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.PlaybackData)
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

// Get provides a mock function with given fields: key
func (_m *Playback) Get(key *ari.Key) *ari.PlaybackHandle {
	ret := _m.Called(key)

	var r0 *ari.PlaybackHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.PlaybackHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.PlaybackHandle)
		}
	}

	return r0
}

// Stop provides a mock function with given fields: key
func (_m *Playback) Stop(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: key, n
func (_m *Playback) Subscribe(key *ari.Key, n ...string) ari.Subscription {
	_va := make([]interface{}, len(n))
	for _i := range n {
		_va[_i] = n[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ari.Subscription
	if rf, ok := ret.Get(0).(func(*ari.Key, ...string) ari.Subscription); ok {
		r0 = rf(key, n...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Subscription)
		}
	}

	return r0
}
