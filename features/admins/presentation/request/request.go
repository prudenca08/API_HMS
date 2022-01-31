package request

import "finalproject/features/admins"

type Admins struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Admins) ToDomain() *admins.Domain {
	return &admins.Domain{
		Username: req.Username,
		Password: req.Password,
	}
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
