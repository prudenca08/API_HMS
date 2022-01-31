package data

import (
	"finalproject/features/patient"

	"finalproject/features/patient/bussiness"

	"gorm.io/gorm"
)

type MysqlPatientRepository struct {
	Conn *gorm.DB
}

func NewMysqlPatientRepository(conn *gorm.DB) patient.Repository {
	return &MysqlPatientRepository{
		Conn: conn,
	}
}

func (rep *MysqlPatientRepository) Create(patID int, domain *patient.Domain) (patient.Domain, error) {

	pat := fromDomain(*domain)

	pat.AdminID = patID

	result := rep.Conn.Create(&pat)

	if result.Error != nil {
		return patient.Domain{}, result.Error
	}

	return ToDomain(pat), nil

}

func (rep *MysqlPatientRepository) AllPatient() ([]patient.Domain, error) {

	var pat []Patient

	result := rep.Conn.Find(&pat)

	if result.Error != nil {
		return []patient.Domain{}, result.Error
	}

	return toDomainList(pat), nil

}

func (rep *MysqlPatientRepository) Update(admID int, patID int, domain *patient.Domain) (patient.Domain, error) {
	patientUpdate := fromDomain(*domain)

	patientUpdate.ID = patID
	result := rep.Conn.Where("id = ?", patID).Updates(&patientUpdate)

	if result.Error != nil {
		return patient.Domain{}, bussiness.ErrNotFound
	}

	return toDomainUpdate(patientUpdate), nil
}

func (rep *MysqlPatientRepository) Delete(patID int, id int) (string, error) {
	rec := Patient{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", bussiness.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", bussiness.ErrNotFound
	}

	return "Patient has been delete", nil

}

func (rep *MysqlPatientRepository) PatientByID(id int) (patient.Domain, error) {

	var pat Patient

	result := rep.Conn.Where("id = ?", id).First(&pat)

	if result.Error != nil {
		return patient.Domain{}, result.Error
	}

	return ToDomain(pat), nil
}
