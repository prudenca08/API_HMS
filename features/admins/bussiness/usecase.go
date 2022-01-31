package bussiness

import (
	"finalproject/features/admins"
	"finalproject/helpers/encrypt"
	"finalproject/middleware"
	"time"
)

type serviceAdmin struct {
	adminRepository admins.Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewServiceAdmin(repoAdmin admins.Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) admins.Service {
	return &serviceAdmin{
		adminRepository: repoAdmin,
		contextTimeout:  timeout,
		jwtAuth:         jwtauth,
	}
}

func (serv *serviceAdmin) Register(domain *admins.Domain) (admins.Domain, error) {

	hashedPassword, err := encrypt.HashingPassword(domain.Password)

	if err != nil {
		return admins.Domain{}, ErrInternalServer
	}

	domain.Password = hashedPassword

	result, err := serv.adminRepository.Register(domain)

	if err != nil {
		return admins.Domain{}, ErrInternalServer
	}
	return result, nil
}

func (serv *serviceAdmin) Login(username, password string) (admins.Domain, error) {

	result, err := serv.adminRepository.Login(username, password)

	if err != nil {
		return admins.Domain{}, err
	}

	checkPass := encrypt.CheckPasswordHash(password, result.Password)

	if !checkPass {
		return admins.Domain{}, ErrPass
	}

	result.Token = serv.jwtAuth.GenerateToken(result.ID, "admin")

	return result, nil
}
