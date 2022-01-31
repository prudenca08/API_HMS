package bussiness_test

import (
	"finalproject/features/admins"
	bussiness "finalproject/features/admins/bussiness"
	_adminMock "finalproject/features/admins/mocks"
	"finalproject/helpers/encrypt"
	"finalproject/middleware"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockAdminRepository _adminMock.Repository
	adminService        admins.Service
	adminDomain         admins.Domain
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "finalproject",
		ExpiresDuration: 1,
	}
	adminService = bussiness.NewServiceAdmin(&mockAdminRepository, 1, jwtAuth)
	adminDomain = admins.Domain{
		ID:       13,
		Username: "admin",
		Password: "admin",
	}
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("admin")
		outputDomain := admins.Domain{
			Username: "admin",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "admin",
			Password: "admin",
		}
		mockAdminRepository.On("Register", mock.Anything).Return(outputDomain, nil).Once()

		resp, err := adminService.Register(&inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Username, resp.Username)
	})

	t.Run("test case 2, invalid test for register wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("admin")
		outputDomain := admins.Domain{
			Username: "admin",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "admin",
			Password: "temon",
		}
		mockAdminRepository.On("Register", mock.Anything).Return(outputDomain, bussiness.ErrInternalServer).Once()

		resp, err := adminService.Register(&inputService)
		assert.Equal(t, err, bussiness.ErrInternalServer)
		assert.Empty(t, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid test for login", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("admin")
		outputDomain := admins.Domain{
			Username: "admin",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "admin",
			Password: "admin",
		}

		mockAdminRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, nil).Once()

		resp, err := adminService.Login(inputService.Username, inputService.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("test case 2, invalid test for login wrong password", func(t *testing.T) {
		password, _ := encrypt.HashingPassword("admin")
		outputDomain := admins.Domain{
			Username: "admin",
			Password: password,
		}
		inputService := admins.Domain{
			Username: "admin",
			Password: "temon",
		}
		mockAdminRepository.On("Register", mock.Anything).Return(outputDomain, bussiness.ErrInternalServer).Once()

		mockAdminRepository.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(outputDomain, bussiness.ErrEmailorPass).Once()

		resp, err := adminService.Login(inputService.Username, inputService.Password)
		assert.Equal(t, err, bussiness.ErrEmailorPass)
		assert.Empty(t, resp)
	})
}