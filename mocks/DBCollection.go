package mocks

import "github.com/ANPez/gogsi/interfaces"
import "github.com/stretchr/testify/mock"

type DBCollection struct {
	mock.Mock
}

func (_m *DBCollection) AggregationIterator(_a0 interface{}) interfaces.DBResultIterator {
	ret := _m.Called(_a0)

	var r0 interfaces.DBResultIterator
	if rf, ok := ret.Get(0).(func(interface{}) interfaces.DBResultIterator); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(interfaces.DBResultIterator)
	}

	return r0
}
func (_m *DBCollection) AggregationOne(_a0 interface{}, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *DBCollection) Find(_a0 interface{}) interfaces.DBQuery {
	ret := _m.Called(_a0)

	var r0 interfaces.DBQuery
	if rf, ok := ret.Get(0).(func(interface{}) interfaces.DBQuery); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(interfaces.DBQuery)
	}

	return r0
}
func (_m *DBCollection) FindOne(_a0 interface{}, _a1 interface{}) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DBCollection) Insert(object interface{}) error {
	ret := _m.Called(object)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(object)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
