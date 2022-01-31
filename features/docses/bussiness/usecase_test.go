package bussiness_test

import (
	"finalproject/features/docses"
	bussiness "finalproject/features/docses/bussiness"
	_docsesMock "finalproject/features/docses/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
var (
	mockDocsesRepository _docsesMock.Repository
	docsesService        docses.Service
	docsesDomain         docses.Domain
)

func TestMain(m *testing.M) {
	docsesService = bussiness.NewServiceDocses(&mockDocsesRepository)
	docsesDomain = docses.Domain{
		ID:          1,
		AdminID:   2,
		Day:        "Senin",
		Time: "12/12/2021",
	}
}

func TestAllDocses(t *testing.T) {
	t.Run("test case 1, valid all docses", func(t *testing.T) {
		expectedReturn := []docses.Domain{
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
		mockDocsesRepository.On("AllDocses").Return(expectedReturn, nil).Once()
		_, err := docsesService.AllDocses()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all  docses", func(t *testing.T) {
		expectedReturn := []docses.Domain{}
		mockDocsesRepository.On("AllDocses").Return(expectedReturn, assert.AnError).Once()
		_, err := docsesService.AllDocses()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestCreate(t *testing.T) {
	t.Run("test case 1, valid create docses", func(t *testing.T) {
		outputDomain := docses.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		inputService := docses.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		mockDocsesRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := docsesService.Create(inputService.AdminID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.ID, resp.ID)
	})

	t.Run("test case 2, invalid create docses", func(t *testing.T) {
		outputDomain := docses.Domain{}
		inputService := docses.Domain{
			ID: 3,
			AdminID: 4,
			Day: "Selasa",
			Time: "13/12/2021",
		}
		mockDocsesRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()

		resp, err := docsesService.Create(inputService.AdminID, &inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid update docses", func(t *testing.T) {
		outputDomain := docses.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		inputService := docses.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		mockDocsesRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := docsesService.Update(inputService.AdminID, inputService.ID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.ID, resp.ID)
	})

	t.Run("test case 2, invalid update docses", func(t *testing.T) {
		outputDomain := docses.Domain{}
		inputService := docses.Domain{
			ID: 3,
			AdminID: 4,
			Day: "Selasa",
			Time: "13/12/2021",
		}
		mockDocsesRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := docsesService.Update(inputService.AdminID, inputService.ID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid delete docses", func(t *testing.T) {
		mockDocsesRepository.On("Delete", mock.Anything, mock.Anything).Return("Docses has been delete", nil).Once()
		resp, err := docsesService.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Docses has been delete", resp)
	})

	t.Run("test case 2, invalid delete docses", func(t *testing.T) {
		mockDocsesRepository.On("Delete", mock.Anything, mock.Anything).Return("", bussiness.ErrNotFound).Once()
		resp, err := docsesService.Delete(2, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestDocsesByID(t *testing.T) {
	t.Run("test case 1, valid all docses", func(t *testing.T) {
		expectedReturn := docses.Domain{
			ID: 1,
			AdminID: 2,
			Day: "Senin",
			Time: "12/12/2021",
		}
		mockDocsesRepository.On("DocsesByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := docsesService.DocsesByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all Docses", func(t *testing.T) {
		expectedReturn := docses.Domain{}
		mockDocsesRepository.On("DocsesByID", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := docsesService.DocsesByID(1)
		assert.Equal(t, err, assert.AnError)

	})
}