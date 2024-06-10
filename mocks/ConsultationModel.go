// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	consultation "PetPalApp/features/consultation"

	mock "github.com/stretchr/testify/mock"
)

// ConsultationModel is an autogenerated mock type for the ConsultationModel type
type ConsultationModel struct {
	mock.Mock
}

// CreateConsultation provides a mock function with given fields: _a0
func (_m *ConsultationModel) CreateConsultation(_a0 consultation.ConsultationCore) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateConsultation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(consultation.ConsultationCore) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConsultations provides a mock function with given fields: currentID
func (_m *ConsultationModel) GetConsultations(currentID uint) ([]consultation.ConsultationCore, error) {
	ret := _m.Called(currentID)

	if len(ret) == 0 {
		panic("no return value specified for GetConsultations")
	}

	var r0 []consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]consultation.ConsultationCore, error)); ok {
		return rf(currentID)
	}
	if rf, ok := ret.Get(0).(func(uint) []consultation.ConsultationCore); ok {
		r0 = rf(currentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(currentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsultationsByDoctorID provides a mock function with given fields: doctorID
func (_m *ConsultationModel) GetConsultationsByDoctorID(doctorID uint) ([]consultation.ConsultationCore, error) {
	ret := _m.Called(doctorID)

	if len(ret) == 0 {
		panic("no return value specified for GetConsultationsByDoctorID")
	}

	var r0 []consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]consultation.ConsultationCore, error)); ok {
		return rf(doctorID)
	}
	if rf, ok := ret.Get(0).(func(uint) []consultation.ConsultationCore); ok {
		r0 = rf(doctorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(doctorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsultationsByUserID provides a mock function with given fields: userID
func (_m *ConsultationModel) GetConsultationsByUserID(userID uint) ([]consultation.ConsultationCore, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetConsultationsByUserID")
	}

	var r0 []consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]consultation.ConsultationCore, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []consultation.ConsultationCore); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCuntationsById provides a mock function with given fields: id
func (_m *ConsultationModel) GetCuntationsById(id uint) (*consultation.ConsultationCore, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetCuntationsById")
	}

	var r0 *consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*consultation.ConsultationCore, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *consultation.ConsultationCore); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConsultation provides a mock function with given fields: consultationID, core
func (_m *ConsultationModel) UpdateConsultation(consultationID uint, core consultation.ConsultationCore) error {
	ret := _m.Called(consultationID, core)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConsultation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, consultation.ConsultationCore) error); ok {
		r0 = rf(consultationID, core)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerAdmin provides a mock function with given fields: userID, doctorID, roomchatID
func (_m *ConsultationModel) VerAdmin(userID uint, doctorID uint, roomchatID uint) (*consultation.ConsultationCore, error) {
	ret := _m.Called(userID, doctorID, roomchatID)

	if len(ret) == 0 {
		panic("no return value specified for VerAdmin")
	}

	var r0 *consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint) (*consultation.ConsultationCore, error)); ok {
		return rf(userID, doctorID, roomchatID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, uint) *consultation.ConsultationCore); ok {
		r0 = rf(userID, doctorID, roomchatID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, uint) error); ok {
		r1 = rf(userID, doctorID, roomchatID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerAvailConcul provides a mock function with given fields: currentUserId, id
func (_m *ConsultationModel) VerAvailConcul(currentUserId uint, id uint) (*consultation.ConsultationCore, error) {
	ret := _m.Called(currentUserId, id)

	if len(ret) == 0 {
		panic("no return value specified for VerAvailConcul")
	}

	var r0 *consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (*consultation.ConsultationCore, error)); ok {
		return rf(currentUserId, id)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) *consultation.ConsultationCore); ok {
		r0 = rf(currentUserId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(currentUserId, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerIsDoctor provides a mock function with given fields: userid, id
func (_m *ConsultationModel) VerIsDoctor(userid uint, id uint) (*consultation.ConsultationCore, error) {
	ret := _m.Called(userid, id)

	if len(ret) == 0 {
		panic("no return value specified for VerIsDoctor")
	}

	var r0 *consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (*consultation.ConsultationCore, error)); ok {
		return rf(userid, id)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) *consultation.ConsultationCore); ok {
		r0 = rf(userid, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userid, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerUser provides a mock function with given fields: userID, doctorID, roomchatID
func (_m *ConsultationModel) VerUser(userID uint, doctorID uint, roomchatID uint) (*consultation.ConsultationCore, error) {
	ret := _m.Called(userID, doctorID, roomchatID)

	if len(ret) == 0 {
		panic("no return value specified for VerUser")
	}

	var r0 *consultation.ConsultationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint) (*consultation.ConsultationCore, error)); ok {
		return rf(userID, doctorID, roomchatID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, uint) *consultation.ConsultationCore); ok {
		r0 = rf(userID, doctorID, roomchatID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consultation.ConsultationCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, uint) error); ok {
		r1 = rf(userID, doctorID, roomchatID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewConsultationModel creates a new instance of ConsultationModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsultationModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConsultationModel {
	mock := &ConsultationModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
