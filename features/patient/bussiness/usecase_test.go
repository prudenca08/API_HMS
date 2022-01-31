package bussiness_test

import (
	"finalproject/features/patient"
	bussiness "finalproject/features/patient/bussiness"
	_patientMock "finalproject/features/patient/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockPatientRepository _patientMock.Repository
	patientService        patient.Service
	patientDomain         patient.Domain
)

func TestMain(m *testing.M) {
	patientService = bussiness.NewServicePatient(&mockPatientRepository)
	patientDomain = patient.Domain{
		ID        : 1,
		AdminID   : 1,
		Name      : "dapoganteng",
		NIK       : "12312313",
		DOB       : "12/12/2004",
		Gender    : "pria",
		Phone     : "0987654321",
		Address   : "yogyakarta",
	
	}
}

func TestAllPatient(t *testing.T) {
	t.Run("test case 1, valid all patient", func(t *testing.T) {
		expectedReturn := []patient.Domain{
			{
				ID        : 1,
				AdminID   : 3,
				Name      : "dapoganteng",
				NIK       : "12312313",
				DOB       : "12/12/2004",
				Gender    : "pria",
				Phone     : "0987654321",
				Address   : "yogyakarta",
				
			},
			{
				ID        : 6,
				AdminID   : 8,
				Name      : "ikhsanoke",
				NIK       : "3435",
				DOB       : "11/11/2001",
				Gender    : "pria",
				Phone     : "3254365465",
				Address   : "yogyakarta",
				
			},
		}
		mockPatientRepository.On("AllPatient").Return(expectedReturn, nil).Once()
		_, err := patientService.AllPatient()
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all patient", func(t *testing.T) {
		expectedReturn := []patient.Domain{}
		mockPatientRepository.On("AllPatient").Return(expectedReturn, assert.AnError).Once()
		_, err := patientService.AllPatient()
		assert.Equal(t, err, assert.AnError)

	})
}

func TestCreate(t *testing.T) {
	t.Run("test case 1, valid create patient", func(t *testing.T) {
		outputDomain := patient.Domain{
			ID        : 1,
			AdminID   : 3,
			Name      : "dapoganteng",
			NIK       : "12312313",
			DOB       : "12/12/2004",
			Gender    : "pria",
			Phone     : "0987654321",
			Address   : "yogyakarta",
			
		}
		inputService := patient.Domain{
			ID        : 1,
			AdminID   : 3,
			Name      : "dapoganteng",
			NIK       : "12312313",
			DOB       : "12/12/2004",
			Gender    : "pria",
			Phone     : "0987654321",
			Address   : "yogyakarta",
			
		}
		mockPatientRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := patientService.Create(inputService.AdminID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid create patient", func(t *testing.T) {
		outputDomain := patient.Domain{}
		inputService := patient.Domain{
			ID        : 1,
			AdminID   : 3,
			Name      : "dapoganteng",
			NIK       : "12312313",
			DOB       : "12/12/2004",
			Gender    : "pria",
			Phone     : "0987654321",
			Address   : "yogyakarta",
			 
		}
		mockPatientRepository.On("Create", mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()

		resp, err := patientService.Create(inputService.AdminID, &inputService)
		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid update patient", func(t *testing.T) {
		outputDomain := patient.Domain{
			ID        : 1,
			AdminID   : 3,
			Name      : "dapoganteng",
			NIK       : "12312313",
			DOB       : "12/12/2004",
			Gender    : "pria",
			Phone     : "0987654321",
			Address   : "yogyakarta",
			
		}
		inputService := patient.Domain{
			ID        : 1,
			AdminID   : 3,
			Name      : "dapoganteng",
			NIK       : "12312313",
			DOB       : "12/12/2004",
			Gender    : "pria",
			Phone     : "0987654321",
			Address   : "yogyakarta",
			
		}
		mockPatientRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, nil).Once()
		resp, err := patientService.Update(inputService.AdminID, inputService.ID, &inputService)
		assert.Nil(t, err)
		assert.Equal(t, inputService.Name, resp.Name)
	})

	t.Run("test case 2, invalid update patient", func(t *testing.T) {
		outputDomain := patient.Domain{}
		inputService := patient.Domain{
			ID        : 1,
			AdminID   : 3,
			Name      : "dapoganteng",
			NIK       : "12312313",
			DOB       : "12/12/2004",
			Gender    : "pria",
			Phone     : "0987654321",
			Address   : "yogyakarta",
			
		}
		mockPatientRepository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(outputDomain, assert.AnError).Once()
		resp, err := patientService.Update(inputService.AdminID, inputService.ID, &inputService)

		assert.Empty(t, resp)
		assert.Equal(t, err, assert.AnError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid delete patient", func(t *testing.T) {
		mockPatientRepository.On("Delete", mock.Anything, mock.Anything).Return("Patient has been delete", nil).Once()
		resp, err := patientService.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Patient has been delete", resp)
	})

	t.Run("test case 2, invalid delete patient", func(t *testing.T) {
		mockPatientRepository.On("Delete", mock.Anything, mock.Anything).Return("", bussiness.ErrNotFound).Once()
		resp, err := patientService.Delete(2, 2)
		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestPatientByID(t *testing.T) {
	t.Run("test case 1, valid all patient", func(t *testing.T) {
		expectedReturn := patient.Domain{
			ID        : 1,
			AdminID   : 3,
			Name      : "dapoganteng",
			NIK       : "12312313",
			DOB       : "12/12/2004",
			Gender    : "pria",
			Phone     : "0987654321",
			Address   : "yogyakarta",
			
		}
		mockPatientRepository.On("PatientByID", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := patientService.PatientByID(1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, invalid all patient", func(t *testing.T) {
		expectedReturn := patient.Domain{}
		mockPatientRepository.On("PatientByID", mock.Anything).Return(expectedReturn, assert.AnError).Once()
		_, err := patientService.PatientByID(1)
		assert.Equal(t, err, assert.AnError)

	})
}