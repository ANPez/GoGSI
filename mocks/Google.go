package mocks

import "github.com/stretchr/testify/mock"

import "github.com/anpez/gogsi/types"

type Google struct {
	mock.Mock
}

func (_m *Google) VerifyToken(_a0 string) (*types.User, error) {
	ret := _m.Called(_a0)

	var r0 *types.User
	if rf, ok := ret.Get(0).(func(string) *types.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
