package bussiness

import (
	"finalproject/features/patient"
)

type servicePatient struct {
	patientRepository patient.Repository
}

func NewServicePatient(repoPatient patient.Repository) patient.Service {
	return &servicePatient{
		patientRepository: repoPatient,
	}
}

func (serv *servicePatient) AllPatient() ([]patient.Domain, error) {

	result, err := serv.patientRepository.AllPatient()

	if err != nil {
		return []patient.Domain{}, err
	}

	return result, nil
}

func (serv *servicePatient) Create(orgID int, domain *patient.Domain) (patient.Domain, error) {

	result, err := serv.patientRepository.Create(orgID, domain)

	if err != nil {
		return patient.Domain{}, err
	}

	return result, nil
}

func (serv *servicePatient) Update(orgID int, prodID int, domain *patient.Domain) (patient.Domain, error) {

	result, err := serv.patientRepository.Update(orgID, prodID, domain)

	if err != nil {
		return patient.Domain{}, err
	}

	return result, nil
}

func (serv *servicePatient) Delete(orgID int, id int) (string, error) {

	result, err := serv.patientRepository.Delete(orgID, id)

	if err != nil {
		return "", ErrNotFound
	}

	return result, nil
}

func (serv *servicePatient) PatientByID(id int) (patient.Domain, error) {

	result, err := serv.patientRepository.PatientByID(id)

	if err != nil {
		return patient.Domain{}, err
	}

	return result, nil
}
