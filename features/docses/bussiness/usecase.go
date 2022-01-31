package bussiness

import (
	"finalproject/features/docses"
)

type serviceDocses struct {
	docsesRepository docses.Repository
}

func NewServiceDocses(repoDocses docses.Repository) docses.Service {
	return &serviceDocses{
		docsesRepository: repoDocses,
	}
}

func (serv *serviceDocses) AllDocses() ([]docses.Domain, error) {

	result, err := serv.docsesRepository.AllDocses()

	if err != nil {
		return []docses.Domain{}, err
	}

	return result, nil
}

func (serv *serviceDocses) Create(dsID int, domain *docses.Domain) (docses.Domain, error) {

	result, err := serv.docsesRepository.Create(dsID, domain)

	if err != nil {
		return docses.Domain{}, err
	}

	return result, nil
}

func (serv *serviceDocses) Update(admID int, dsID int, domain *docses.Domain) (docses.Domain, error) {

	result, err := serv.docsesRepository.Update(admID, dsID, domain)

	if err != nil {
		return docses.Domain{}, err
	}

	return result, nil
}

func (serv *serviceDocses) Delete(dsID int, id int) (string, error) {

	result, err := serv.docsesRepository.Delete(dsID, id)

	if err != nil {
		return "", ErrNotFound
	}

	return result, nil
}

func (serv *serviceDocses) DocsesByID(id int) (docses.Domain, error) {

	result, err := serv.docsesRepository.DocsesByID(id)

	if err != nil {
		return docses.Domain{}, err
	}

	return result, nil
}
