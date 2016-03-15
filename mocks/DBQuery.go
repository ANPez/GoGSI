package mocks

import "github.com/anpez/gogsi/interfaces"
import "github.com/stretchr/testify/mock"

type DBQuery struct {
	mock.Mock
}

func (_m *DBQuery) Distinct(_a0 string, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *DBQuery) Iterator() interfaces.DBResultIterator {
	ret := _m.Called()

	var r0 interfaces.DBResultIterator
	if rf, ok := ret.Get(0).(func() interfaces.DBResultIterator); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(interfaces.DBResultIterator)
	}

	return r0
}
