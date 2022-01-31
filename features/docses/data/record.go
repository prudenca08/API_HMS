package data

import (
	"finalproject/features/docses"
	"time"

	"gorm.io/gorm"
)

type Docses struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	AdminID   int
	Day       string
	Time      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDomain(ds Docses) docses.Domain {
	return docses.Domain{
		ID:        ds.ID,
		AdminID:   ds.AdminID,
		Day:       ds.Day,
		Time:      ds.Time,
		CreatedAt: ds.CreatedAt,
		UpdatedAt: ds.UpdatedAt,
	}
}

func fromDomain(domain docses.Domain) Docses {
	return Docses{
		ID:        domain.ID,
		AdminID:   domain.AdminID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(ds Docses) docses.Domain {
	return docses.Domain{
		ID:        ds.ID,
		AdminID:   ds.AdminID,
		Day:       ds.Day,
		Time:      ds.Time,
		CreatedAt: ds.CreatedAt,
		UpdatedAt: ds.UpdatedAt,
	}
}
func toDomainList(data []Docses) []docses.Domain {
	result := []docses.Domain{}

	for _, ds := range data {
		result = append(result, ToDomain(ds))
	}
	return result
}
