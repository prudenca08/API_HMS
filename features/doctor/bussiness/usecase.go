package bussiness

import (
	"finalproject/features/doctor"
	"finalproject/helpers/encrypt"
	"finalproject/middleware"
	"time"
)

type serviceDoctor struct {
	doctorRepository doctor.Repository
	contextTimeout   time.Duration
	jwtAuth          *middleware.ConfigJWT
}

func NewServiceDoctor(repoDoctor doctor.Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) doctor.Service {
	return &serviceDoctor{
		doctorRepository: repoDoctor,
		contextTimeout:   timeout,
		jwtAuth:          jwtauth,
	}
}

func (serv *serviceDoctor) Register(domain *doctor.Domain) (doctor.Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return doctor.Domain{}, ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.doctorRepository.Register(domain)

	if result == (doctor.Domain{}) {
		return doctor.Domain{}, ErrDuplicateData
	}

	if err != nil {
		return doctor.Domain{}, ErrInternalServer
	}
	return result, nil
}

func (serv *serviceDoctor) Login(email, password string) (doctor.Domain, error) {

	result, err := serv.doctorRepository.Login(email, password)

	if err != nil {
		return doctor.Domain{}, err
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return doctor.Domain{}, ErrPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "doctor")

	return result, nil
}
func (serv *serviceDoctor) AllDoctor() ([]doctor.Domain, error) {

	result, err := serv.doctorRepository.AllDoctor()

	if err != nil {
		return []doctor.Domain{}, err
	}

	return result, nil
}

func (serv *serviceDoctor) Update(docID int, domain *doctor.Domain) (doctor.Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)
	domain.Password = hashedPassword
	result, err := serv.doctorRepository.Update(docID, domain)

	if err != nil {
		return doctor.Domain{}, err
	}

	return result, nil
}

func (serv *serviceDoctor) ChangePass(docID int, domain *doctor.Domain) (doctor.Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)
	domain.Password = hashedPassword
	result, err := serv.doctorRepository.ChangePass(docID, domain)

	if err != nil {
		return doctor.Domain{}, err
	}

	return result, nil
}

func (serv *serviceDoctor) DoctorByID(id int) (doctor.Domain, error) {

	result, err := serv.doctorRepository.DoctorByID(id)

	if err != nil {
		return doctor.Domain{}, err
	}

	return result, nil
}

func (serv *serviceDoctor) Delete(docID int, id int) (string, error) {

	result, err := serv.doctorRepository.Delete(docID, id)

	if err != nil {
		return "", ErrNotFound
	}

	return result, nil
}
