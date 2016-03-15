package mocks

import "github.com/anpez/gogsi/interfaces"
import "github.com/stretchr/testify/mock"

type Database struct {
	mock.Mock
}

func (_m *Database) C(_a0 string) interfaces.DBCollection {
	ret := _m.Called(_a0)

	var r0 interfaces.DBCollection
	if rf, ok := ret.Get(0).(func(string) interfaces.DBCollection); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(interfaces.DBCollection)
	}

	return r0
}
