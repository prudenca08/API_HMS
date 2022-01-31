package bussiness_test

import (
	"finalproject/features/patsche"
	"finalproject/features/patsche/bussiness"
	_patscheMock "finalproject/features/patsche/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockPatscheRepository _patscheMock.Repository
	patscheService        patsche.Service
	patscheDomain         patsche.Domain
)

func TestMain(m *testing.M) {
	patscheService = bussiness.NewServicePatsche(&mockPatscheRepository)
	patscheDomain = patsche.Domain{
		ID:          1,
		AdminID:   2,
		Day:        "Senin",
		Time: "12/12/2021",
	}
}

func TestAllPatsche(t *testing.T) {
	t.Run("test case 1, valid all patsche", func(t *testing.T) {
		expectedReturn := []patsche.Domain{
			{
				ID: 1,
				AdminID: 2,
				Day: "Senin",
				Time: "12/12/2021",
			},
			{
				ID: 3,
				AdminID: 4,
				Day: "Selasa",
				Time: "13/12/2021",
			},
		}
		mockPatscheRepository.On("AllPatsche").Return(expectedReturn, nil).Once()
		_, err := patscheService.AllPatsche()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all patsche", func(t *testing.T) {
		expectedReturn := []patsche.Domain{}
		mockPatscheRepository.On("AllPatsche").Return(expectedReturn, assert.AnError).Once()
		_, err := patscheService.AllPatsche()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestCreate(t *testing.T) {
	t.Run("test case 1, valid create patsche", func(t *testing.T) {
		outputDomain := patsche.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		inputService := patsche.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		mockPatscheRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := patscheService.Create(inputService.AdminID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.ID, resp.ID)
	})

	t.Run("test case 2, invalid create patsche", func(t *testing.T) {
		outputDomain := patsche.Domain{}
		inputService := patsche.Domain{
			ID: 3,
			AdminID: 4,
			Day: "Selasa",
			Time: "13/12/2021",
		}
		mockPatscheRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()

		resp, err := patscheService.Create(inputService.AdminID, &inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid update patsche", func(t *testing.T) {
		outputDomain := patsche.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		inputService := patsche.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		mockPatscheRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := patscheService.Update(inputService.AdminID, inputService.ID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.ID, resp.ID)
	})

	t.Run("test case 2, invalid update patsche", func(t *testing.T) {
		outputDomain := patsche.Domain{}
		inputService := patsche.Domain{
			ID: 3,
			AdminID: 4,
			Day: "Selasa",
			Time: "13/12/2021",
		}
		mockPatscheRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := patscheService.Update(inputService.AdminID, inputService.ID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}


func TestDelete(t *testing.T) {
	t.Run("test case 1, valid delete patsche", func(t *testing.T) {
		mockPatscheRepository.On("Delete", mock.Anything, mock.Anything).Return("Patsche has been delete", nil).Once()
		resp, err := patscheService.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Patsche has been delete", resp)
	})

	t.Run("test case 2, invalid delete patsche", func(t *testing.T) {
		mockPatscheRepository.On("Delete", mock.Anything, mock.Anything).Return("", bussiness.ErrNotFound).Once()
		resp, err := patscheService.Delete(2, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestPatscheByID(t *testing.T) {
	t.Run("test case 1, valid all patsche", func(t *testing.T) {
		expectedReturn := patsche.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		mockPatscheRepository.On("PatscheByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := patscheService.PatscheByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all Patsche", func(t *testing.T) {
		expectedReturn := patsche.Domain{}
		mockPatscheRepository.On("PatscheByID", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := patscheService.PatscheByID(1)
		assert.Equal(t, err, assert.AnError)

	})
}
