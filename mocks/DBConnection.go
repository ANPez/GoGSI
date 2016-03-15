package mocks

import "github.com/anpez/gogsi/interfaces"
import "github.com/stretchr/testify/mock"

type DBConnection struct {
	mock.Mock
}

func (_m *DBConnection) Clone() interfaces.DBConnection {
	ret := _m.Called()

	var r0 interfaces.DBConnection
	if rf, ok := ret.Get(0).(func() interfaces.DBConnection); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(interfaces.DBConnection)
	}

	return r0
}
func (_m *DBConnection) DB(_a0 string) interfaces.Database {
	ret := _m.Called(_a0)

	var r0 interfaces.Database
	if rf, ok := ret.Get(0).(func(string) interfaces.Database); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(interfaces.Database)
	}

	return r0
}
func (_m *DBConnection) Close() {
	_m.Called()
}
