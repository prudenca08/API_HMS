package data

import (
	"finalproject/features/patient"
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	AdminID   int
	Name      string
	NIK       string
	DOB       string
	Gender    string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDomain(pat Patient) patient.Domain {
	return patient.Domain{
		ID:        pat.ID,
		AdminID:   pat.AdminID,
		Name:      pat.Name,
		NIK:       pat.NIK,
		DOB:       pat.DOB,
		Gender:    pat.Gender,
		Phone:     pat.Phone,
		Address:   pat.Address,
		CreatedAt: pat.CreatedAt,
		UpdatedAt: pat.UpdatedAt,
	}
}

func fromDomain(domain patient.Domain) Patient {
	return Patient{
		ID:        domain.ID,
		AdminID:   domain.AdminID,
		Name:      domain.Name,
		NIK:       domain.NIK,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(pat Patient) patient.Domain {
	return patient.Domain{
		ID:        pat.ID,
		AdminID:   pat.AdminID,
		Name:      pat.Name,
		NIK:       pat.NIK,
		DOB:       pat.DOB,
		Gender:    pat.Gender,
		Phone:     pat.Phone,
		Address:   pat.Address,
		CreatedAt: pat.CreatedAt,
		UpdatedAt: pat.UpdatedAt,
	}
}
func toDomainList(data []Patient) []patient.Domain {
	result := []patient.Domain{}

	for _, pat := range data {
		result = append(result, ToDomain(pat))
	}
	return result
}
