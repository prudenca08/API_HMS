package bussiness_test

import (
	"finalproject/features/patientses"
	bussiness "finalproject/features/patientses/bussiness"
	_patientsesMock "finalproject/features/patientses/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
var (
	mockPatientsesRepository _patientsesMock.Repository
	patientsesService        patientses.Service
	patientsesDomain         patientses.Domain
)

func TestMain(m *testing.M) {
	patientsesService = bussiness.NewServicePatientses(&mockPatientsesRepository)
	patientsesDomain = patientses.Domain{
		ID: 1,
		AdminID: 2,
		DoctorID: 3,
		PatientID: 1,
		PatientScheduleID: 1,
		Date: "12/12/2021",
		Title: "Meriang",
		DetailRecipe: "Perbanyak tidur",
		Status: "Active",
	}
}


func TestAllPatientses(t *testing.T) {
	t.Run("test case 1, valid all patientses", func(t *testing.T) {
		expectedReturn := []patientses.Domain{
			{
			ID: 1,
			AdminID: 2,
			DoctorID: 3,
			PatientID: 1,
			PatientScheduleID: 1,
			Date: "12/12/2021",
			Title: "Meriang",
			DetailRecipe: "Perbanyak tidur",
			Status: "Active",
			},
			{
			ID: 3,
			AdminID: 1,
			DoctorID: 2,
			PatientID: 5,
			PatientScheduleID: 4,
			Date: "11/10/2021",
			Title: "Batuk",
			DetailRecipe: "Ngopi terus",
			Status: "NonActive",
			},
		}
		mockPatientsesRepository.On("AllPatientses").Return(expectedReturn, nil).Once()
		_, err := patientsesService.AllPatientses()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all patientses", func(t *testing.T) {
		expectedReturn := []patientses.Domain{}
		mockPatientsesRepository.On("AllPatientses").Return(expectedReturn, assert.AnError).Once()
		_, err := patientsesService.AllPatientses()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestCreate(t *testing.T) {
	t.Run("test case 1, valid create patientses", func(t *testing.T) {
		outputDomain := patientses.Domain{
			ID: 1,
			AdminID: 2,
			DoctorID: 3,
			PatientID: 1,
			PatientScheduleID: 1,
			Date: "12/12/2021",
			Title: "Meriang",
			DetailRecipe: "Perbanyak tidur",
			Status: "Active",
		}
		inputService := patientses.Domain{
			ID: 1,
			AdminID: 2,
			DoctorID: 3,
			PatientID: 1,
			PatientScheduleID: 1,
			Date: "12/12/2021",
			Title: "Meriang",
			DetailRecipe: "Perbanyak tidur",
			Status: "Active",
		}
		mockPatientsesRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := patientsesService.Create(inputService.DoctorID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Status, resp.Status)
	})

	t.Run("test case 2, invalid create patientses", func(t *testing.T) {
		outputDomain := patientses.Domain{}
		inputService := patientses.Domain{
			ID: 3,
			AdminID: 1,
			DoctorID: 2,
			PatientID: 5,
			PatientScheduleID: 4,
			Date: "11/10/2021",
			Title: "Batuk",
			DetailRecipe: "Ngopi terus",
			Status: "NonActive",
		}
		mockPatientsesRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()

		resp, err := patientsesService.Create(inputService.DoctorID, &inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid update patientses", func(t *testing.T) {
		outputDomain := patientses.Domain{
			ID: 1,
			AdminID: 2,
			DoctorID: 3,
			PatientID: 1,
			PatientScheduleID: 1,
			Date: "12/12/2021",
			Title: "Meriang",
			DetailRecipe: "Perbanyak tidur",
			Status: "Active",
		}
		inputService := patientses.Domain{
			ID: 1,
			AdminID: 2,
			DoctorID: 3,
			PatientID: 1,
			PatientScheduleID: 1,
			Date: "12/12/2021",
			Title: "Meriang",
			DetailRecipe: "Perbanyak tidur",
			Status: "Active",
		}
		mockPatientsesRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := patientsesService.Update(inputService.DoctorID, inputService.ID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Status, resp.Status)
	})

	t.Run("test case 2, invalid update patientses", func(t *testing.T) {
		outputDomain := patientses.Domain{}
		inputService := patientses.Domain{
			ID: 3,
			AdminID: 1,
			DoctorID: 2,
			PatientID: 5,
			PatientScheduleID: 4,
			Date: "11/10/2021",
			Title: "Batuk",
			DetailRecipe: "Ngopi terus",
			Status: "NonActive",
		}
		mockPatientsesRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := patientsesService.Update(inputService.DoctorID, inputService.ID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid delete patientses", func(t *testing.T) {
		mockPatientsesRepository.On("Delete", mock.Anything, mock.Anything).Return("Patientses has been delete", nil).Once()
		resp, err := patientsesService.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Patientses has been delete", resp)
	})

	t.Run("test case 2, invalid delete patientses", func(t *testing.T) {
		mockPatientsesRepository.On("Delete", mock.Anything, mock.Anything).Return("", bussiness.ErrNotFound).Once()
		resp, err := patientsesService.Delete(2, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestPatientsesByID(t *testing.T) {
	t.Run("test case 1, valid all patientses", func(t *testing.T) {
		expectedReturn := patientses.Domain{
		ID: 1,
		AdminID: 2,
		DoctorID: 3,
		PatientID: 1,
		PatientScheduleID: 1,
		Date: "12/12/2021",
		Title: "Meriang",
		DetailRecipe: "Perbanyak tidur",
		Status: "Active",
		}
		mockPatientsesRepository.On("PatientsesByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := patientsesService.PatientsesByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all Product", func(t *testing.T) {
		expectedReturn := patientses.Domain{}
		mockPatientsesRepository.On("PatientsesByID", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := patientsesService.PatientsesByID(1)
		assert.Equal(t, err, assert.AnError)

	})
}
