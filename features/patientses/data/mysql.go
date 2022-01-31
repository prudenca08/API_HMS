package data

import (
	"finalproject/features/patientses"
	"finalproject/features/patientses/bussiness"

	"gorm.io/gorm"
)

type MysqlPatientsesRepository struct {
	Conn *gorm.DB
}

func NewMysqlPatientsesRepository(conn *gorm.DB) patientses.Repository {
	return &MysqlPatientsesRepository{
		Conn: conn,
	}
}

func (rep *MysqlPatientsesRepository) AllPatientses() ([]patientses.Domain, error){
	
	var pss []Patientses
	
	result := rep.Conn.Preload("Doctor").Preload("Patient").Preload("Patsche").Find(&pss)

	
	
	if result.Error != nil {
		return []patientses.Domain{}, result.Error
	}
	return toDomainList(pss), nil
}

func (rep *MysqlPatientsesRepository) Create(pssID int, domain *patientses.Domain) (patientses.Domain,error){
	dss := fromDomain(*domain)

	dss.AdminID = pssID
	result := rep.Conn.Create(&dss)
	if result.Error != nil {
		return patientses.Domain{}, result.Error
	}
	return ToDomain(dss), nil
}



func (rep *MysqlPatientsesRepository) Update(admID int, pssID int, domain *patientses.Domain)(patientses.Domain, error){
	patientsesUpdate := fromDomain(*domain)

	patientsesUpdate.ID = pssID
	result := rep.Conn.Where("id = ?", pssID).Updates(&patientsesUpdate)

	if result.Error != nil {
		return patientses.Domain{}, bussiness.ErrNotFound
	}
	
	return toDomainUpdate(patientsesUpdate), nil
}

func (rep *MysqlPatientsesRepository) Delete(pssID int, id int) (string, error){
	rec := Patientses{}
	find := rep.Conn.Where("id = ?", id).First(&rec)
	if find.Error != nil {
		return "", bussiness.ErrUnathorized		
	}
	err := rep.Conn.Delete(&rec, "id = ?", id). Error

	if err != nil {
		return "", bussiness.ErrNotFound
	}
	return "Patient Sesion has been delete", nil
}

func (rep *MysqlPatientsesRepository) PatientsesByID(id int) (patientses.Domain, error){
	var pss Patientses
	result := rep.Conn.Where("id = ?", id).First(&pss)
	if result.Error != nil {
		return patientses.Domain{}, result.Error
	}
	return ToDomain(pss), nil
}