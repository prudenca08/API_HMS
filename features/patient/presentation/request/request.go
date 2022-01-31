package request

import "finalproject/features/patient"

type Patient struct {
	Name     string `json:"name"`
	NIK      string `json:"nik"`
	DOB      string `json:"dob"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	
}

func (req *Patient) ToDomain() *patient.Domain {
	return &patient.Domain{
		Name:     req.Name,
		NIK:      req.NIK,
		DOB:      req.DOB,
		Gender:   req.Gender,
		Phone:    req.Phone,
		Address:  req.Address,
		
	}
}
