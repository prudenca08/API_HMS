package data

import (
	doctorrecord "finalproject/features/doctor/data"
	patientrecord "finalproject/features/patient/data"
	"finalproject/features/patientses"
	patscherecord "finalproject/features/patsche/data"

	"time"

	"gorm.io/gorm"
)

type Patientses struct {
	gorm.Model
	ID                int `gorm:"primary_key"`
	AdminID           int
	DoctorID          int
	PatientID         int
	PatientScheduleID int
	Date              string
	Symptoms          string
	Title             string
	DetailRecipe      string
	Status            string
	Patsche           patscherecord.Patsche `gorm:"foreignKey:ID;references:PatientScheduleID"`
	Doctor            doctorrecord.Doctor   `gorm:"foreignKey:ID;references:DoctorID"`
	Patient           patientrecord.Patient `gorm:"foreignKey:ID;references:PatientID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func ToDomain(pss Patientses) patientses.Domain {
	return patientses.Domain{
		ID:                pss.ID,
		AdminID:           pss.AdminID,
		DoctorID:          pss.DoctorID,
		PatientID:         pss.PatientID,
		PatientScheduleID: pss.PatientScheduleID,
		Date:              pss.Date,
		Symptoms:          pss.Symptoms,
		Title:             pss.Title,
		DetailRecipe:      pss.DetailRecipe,
		Status:            pss.Status,
		Patient:           patientrecord.ToDomain(pss.Patient),
		Doctor:            doctorrecord.ToDomain(pss.Doctor),
		Patsche:           patscherecord.ToDomain(pss.Patsche),
		CreatedAt:         pss.CreatedAt,
		UpdatedAt:         pss.UpdatedAt,
	}
}
func fromDomain(domain patientses.Domain) Patientses {
	return Patientses{
		ID:                domain.ID,
		AdminID:           domain.AdminID,
		DoctorID:          domain.DoctorID,
		PatientID:         domain.PatientID,
		PatientScheduleID: domain.PatientScheduleID,
		Date:              domain.Date,
		Symptoms:          domain.Symptoms,
		Title:             domain.Title,
		DetailRecipe:      domain.DetailRecipe,
		Status:            domain.Status,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
func toDomainUpdate(pss Patientses) patientses.Domain {
	return patientses.Domain{
		ID:                pss.ID,
		AdminID:           pss.AdminID,
		DoctorID:          pss.DoctorID,
		PatientID:         pss.PatientID,
		PatientScheduleID: pss.PatientScheduleID,
		Date:              pss.Date,
		Symptoms:          pss.Symptoms,
		Title:             pss.Title,
		DetailRecipe:      pss.DetailRecipe,
		Status:            pss.Status,
		CreatedAt:         pss.CreatedAt,
		UpdatedAt:         pss.UpdatedAt,
	}
}
func toDomainList(data []Patientses) []patientses.Domain {
	result := []patientses.Domain{}
	for _, pss := range data {
		result = append(result, ToDomain(pss))
	}
	return result
}
