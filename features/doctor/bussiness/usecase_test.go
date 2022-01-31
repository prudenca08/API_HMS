package bussiness_test

import (
	"finalproject/features/doctor"
	bussiness "finalproject/features/doctor/bussiness"
	_doctorMock "finalproject/features/doctor/mocks"
	"finalproject/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockDoctorRepository _doctorMock.Repository
	doctorService        doctor.Service
	doctorDomain         doctor.Domain
)

func TestMain(m *testing.M) {
	doctorDomain = doctor.Domain{
		ID                : 13,
		DoctorSessionID   : 13,
		Username          : "dr_dapo",
		Password          : "dapo123",
		Name              : "Dapo",
		NIP               : "123456",
		Experience        : "hjasfdgjas",
		Specialist        : "asdsadsasa",
		Room              : "13",
		Phone_Number      : "0875675767656",
		Status            : "waiting",
	}
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctor.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctor.Domain{
			Username: "doctor",
			Password: "doctor",
		}
		mockDoctorRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := doctorService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})

	t.Run("test case 2, invalid test for register wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctor.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctor.Domain{
			Username: "doctor",
			Password: "temon",
		}
		mockDoctorRepository.On("Register", mock.Anything).Return(outputDomain, bussiness.ErrInternalServer).Once()

		resp, err := doctorService.Register(&inputService)
		assert.Equal(t, err, bussiness.ErrInternalServer)
		assert.Empty(t, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid test for login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctor.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctor.Domain{
			Username: "doctor",
			Password: "doctor",
		}

		mockDoctorRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := doctorService.Login(inputService.Username, inputService.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2, invalid test for login wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("doctor")
		outputDomain := doctor.Domain{
			Username: "doctor",
			Password: password,
		}
		inputService := doctor.Domain{
			Username: "doctor",
			Password: "temon",
		}
		mockDoctorRepository.On("Register", mock.Anything).Return(outputDomain, bussiness.ErrInternalServer).Once()

		mockDoctorRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, bussiness.ErrEmailorPass).Once()

		resp, err := doctorService.Login(inputService.Username, inputService.Password)
		assert.Equal(t, err, bussiness.ErrEmailorPass)
		assert.Empty(t, resp)
	})
}


func TestAllDoctor(t *testing.T) {
	t.Run("test case 1, valid all doctor", func(t *testing.T) {
		expectedReturn := []doctor.Domain{
			{
				ID                : 13,
				DoctorSessionID   : 13,
				Username          : "dr_dapo",
				Password          : "dapo123",
				Name              : "Dapo",
				NIP               : "123456",
				Experience        : "hjasfdgjas",
				Specialist        : "asdsadsasa",
				Room              : "13",
				Phone_Number      : "0875675767656",
				Status            : "waiting",
			},
			{
				ID                : 11,
				DoctorSessionID   : 12,
				Username          : "dr_arya",
				Password          : "arya123",
				Name              : "Arya",
				NIP               : "324324",
				Experience        : "cxvcbv",
				Specialist        : "fdgfdg",
				Room              : "10",
				Phone_Number      : "085612324234",
				Status            : "waiting",
			},
		}
		mockDoctorRepository.On("AllDoctor").Return(expectedReturn, nil).Once()
		_, err := doctorService.AllDoctor()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all doctor", func(t *testing.T) {
		expectedReturn := []doctor.Domain{}
		mockDoctorRepository.On("AllDoctor").Return(expectedReturn, assert.AnError).Once()
		_, err := doctorService.AllDoctor()
		assert.Equal(t, err, assert.AnError)

	})
}


func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid update doctor", func(t *testing.T) {
		outputDomain := doctor.Domain{
			ID                : 13,
			DoctorSessionID   : 13,
			Username          : "dr_dapo",
			Password          : "dapo123",
			Name              : "Dapo",
			NIP               : "123456",
			Experience        : "hjasfdgjas",
			Specialist        : "asdsadsasa",
			Room              : "13",
			Phone_Number      : "0875675767656",
			Status            : "waiting",
		}
		inputService := doctor.Domain{
			ID                : 13,
			DoctorSessionID   : 13,
			Username          : "dr_dapo",
			Password          : "dapo123",
			Name              : "Dapo",
			NIP               : "123456",
			Experience        : "hjasfdgjas",
			Specialist        : "asdsadsasa",
			Room              : "13",
			Phone_Number      : "0875675767656",
			Status            : "waiting",
		}
		mockDoctorRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := doctorService.Update(inputService.DoctorSessionID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid update doctor", func(t *testing.T) {
		outputDomain := doctor.Domain{}
		inputService := doctor.Domain{
			ID                : 11,
			DoctorSessionID   : 12,
			Username          : "dr_arya",
			Password          : "arya123",
			Name              : "Arya",
			NIP               : "324324",
			Experience        : "cxvcbv",
			Specialist        : "fdgfdg",
			Room              : "10",
			Phone_Number      : "085612324234",
			Status            : "waiting",
		}
		mockDoctorRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := doctorService.Update(inputService.DoctorSessionID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}
func TestChangePass(t *testing.T) {
	t.Run("test case 1, valid update changepass", func(t *testing.T) {
		outputDomain := doctor.Domain{
			ID                : 13,
			Password          : "dapo123",
	
		}
		inputService := doctor.Domain{
			ID                : 13,
			Password          : "dapo123",

		}
		mockDoctorRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := doctorService.Update(inputService.DoctorSessionID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid update doctor", func(t *testing.T) {
		outputDomain := doctor.Domain{}
		inputService := doctor.Domain{
			ID                : 11,
			Password          : "arya123",

		}
		mockDoctorRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := doctorService.Update(inputService.DoctorSessionID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid delete doctor", func(t *testing.T) {
		mockDoctorRepository.On("Delete", mock.Anything, mock.Anything).Return("Doctor has been delete", nil).Once()
		resp, err := doctorService.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Doctor has been delete", resp)
	})

	t.Run("test case 2, invalid delete doctor", func(t *testing.T) {
		mockDoctorRepository.On("Delete", mock.Anything, mock.Anything).Return("", bussiness.ErrNotFound).Once()
		resp, err := doctorService.Delete(2, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestPatientByID(t *testing.T) {
	t.Run("test case 1, valid all doctor", func(t *testing.T) {
		expectedReturn := doctor.Domain{
			ID                : 13,
			DoctorSessionID   : 13,
			Username          : "dr_dapo",
			Password          : "dapo123",
			Name              : "Dapo",
			NIP               : "123456",
			Experience        : "hjasfdgjas",
			Specialist        : "asdsadsasa",
			Room              : "13",
			Phone_Number      : "0875675767656",
			Status            : "waiting",
		}
		mockDoctorRepository.On("DoctorByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := doctorService.DoctorByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all doctor", func(t *testing.T) {
		expectedReturn := doctor.Domain{}
		mockDoctorRepository.On("DoctorByID", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := doctorService.DoctorByID(1)
		assert.Equal(t, err, assert.AnError)

	})
}