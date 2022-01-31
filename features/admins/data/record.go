package data

import (
	"finalproject/features/admins"
	"time"

	"gorm.io/gorm"
)

type Admins struct {
	gorm.Model
	ID        int    `gorm:"primary_key"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(admin Admins) admins.Domain {
	return admins.Domain{
		ID:        admin.ID,
		Username:  admin.Username,
		Password:  admin.Password,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func fromDomain(domain admins.Domain) Admins {
	return Admins{
		ID:        domain.ID,
		Username:  domain.Username,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
