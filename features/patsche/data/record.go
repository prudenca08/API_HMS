package data

import (
	"finalproject/features/patsche"
	"time"

	"gorm.io/gorm"
)

type Patsche struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	AdminID   int
	Day       string
	Time      string
	CreatedAt time.Time
	UpdatedAt time.Time
	
}

func ToDomain(ds Patsche) patsche.Domain {
	return patsche.Domain{
		ID:        ds.ID,
		AdminID:   ds.AdminID,
		Day:       ds.Day,
		Time:      ds.Time,
		CreatedAt: ds.CreatedAt,
		UpdatedAt: ds.UpdatedAt,
	}
}

func fromDomain(domain patsche.Domain) Patsche {
	return Patsche{
		ID:        domain.ID,
		AdminID:   domain.AdminID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(ds Patsche) patsche.Domain {
	return patsche.Domain{
		ID:        ds.ID,
		AdminID:   ds.AdminID,
		Day:       ds.Day,
		Time:      ds.Time,
		CreatedAt: ds.CreatedAt,
		UpdatedAt: ds.UpdatedAt,
	}
}
func toDomainList(data []Patsche) []patsche.Domain {
	result := []patsche.Domain{}

	for _, ds := range data {
		result = append(result, ToDomain(ds))
	}
	return result
}
