package bussiness

import (
	"finalproject/features/patsche"
)

type servicePatsche struct {
	patscheRepository patsche.Repository
}

func NewServicePatsche(repoPatsche patsche.Repository) patsche.Service {
	return &servicePatsche{
		patscheRepository: repoPatsche,
	}
}

func (serv *servicePatsche) AllPatsche() ([]patsche.Domain, error) {

	result, err := serv.patscheRepository.AllPatsche()

	if err != nil {
		return []patsche.Domain{}, err
	}

	return result, nil
}

func (serv *servicePatsche) Create(dsID int, domain *patsche.Domain) (patsche.Domain, error) {

	result, err := serv.patscheRepository.Create(dsID, domain)

	if err != nil {
		return patsche.Domain{}, err
	}

	return result, nil
}

func (serv *servicePatsche) Update(admID int, dsID int, domain *patsche.Domain) (patsche.Domain, error) {

	result, err := serv.patscheRepository.Update(admID, dsID, domain)

	if err != nil {
		return patsche.Domain{}, err
	}

	return result, nil
}

func (serv *servicePatsche) Delete(dsID int, id int) (string, error) {

	result, err := serv.patscheRepository.Delete(dsID, id)

	if err != nil {
		return "", ErrNotFound
	}

	return result, nil
}

func (serv *servicePatsche) PatscheByID(id int) (patsche.Domain, error) {

	result, err := serv.patscheRepository.PatscheByID(id)

	if err != nil {
		return patsche.Domain{}, err
	}

	return result, nil
}
