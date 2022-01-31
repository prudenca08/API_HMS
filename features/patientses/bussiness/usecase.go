package bussiness

import (
	"finalproject/features/patientses"
)

type servicePatientses struct {
	patientsesRepository patientses.Repository
}

func NewServicePatientses(repoPatientses patientses.Repository) patientses.Service {
	return &servicePatientses {
		patientsesRepository: repoPatientses,
	}
}

func (serv *servicePatientses) AllPatientses() ([]patientses.Domain, error) {
	result, err := serv.patientsesRepository.AllPatientses()
	if err != nil {
		return []patientses.Domain{}, err

	}
	return result, nil
}

func (serv *servicePatientses) Create(orgID int, domain *patientses.Domain) (patientses.Domain, error){
	result, err := serv.patientsesRepository.Create(orgID, domain)
	if err != nil {
		return patientses.Domain{}, err
	}
	return result, nil
}
func (serv *servicePatientses) Update(orgID int, prodID int, domain *patientses.Domain) (patientses.Domain, error){
	result , err := serv.patientsesRepository.Update(orgID,prodID, domain)
	if err != nil {
		return patientses.Domain{}, err
	}
	return result, nil
}
func (serv *servicePatientses) Delete(orgID int, id int)(string, error){
	result, err := serv.patientsesRepository.Delete(orgID, id)
	if err != nil {
		return "", ErrNotFound
	}
	return result, nil
}

func (serv *servicePatientses) PatientsesByID(id int) (patientses.Domain, error) {
	result, err := serv.patientsesRepository.PatientsesByID(id)
	if err != nil {
		return patientses.Domain{}, err
	}
	return result, nil
}
