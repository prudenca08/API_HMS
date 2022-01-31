package data

import (
	"finalproject/features/patsche"

	"finalproject/features/patsche/bussiness"

	"gorm.io/gorm"
)

type MysqlPatscheRepository struct {
	Conn *gorm.DB
}

func NewMysqlPatscheRepository(conn *gorm.DB) patsche.Repository {
	return &MysqlPatscheRepository{
		Conn: conn,
	}
}

func (rep *MysqlPatscheRepository) Create(dsID int, domain *patsche.Domain) (patsche.Domain, error) {

	dss := fromDomain(*domain)

	dss.AdminID = dsID

	result := rep.Conn.Create(&dss)

	if result.Error != nil {
		return patsche.Domain{}, result.Error
	}

	return ToDomain(dss), nil

}

func (rep *MysqlPatscheRepository) AllPatsche() ([]patsche.Domain, error) {

	var ds []Patsche

	result := rep.Conn.Find(&ds)

	if result.Error != nil {
		return []patsche.Domain{}, result.Error
	}

	return toDomainList(ds), nil

}

func (rep *MysqlPatscheRepository) Update(admID int, dsID int, domain *patsche.Domain) (patsche.Domain, error) {
	patscheUpdate := fromDomain(*domain)

	patscheUpdate.ID = dsID
	result := rep.Conn.Where("id = ?", dsID).Updates(&patscheUpdate)

	if result.Error != nil {
		return patsche.Domain{}, bussiness.ErrNotFound
	}

	return toDomainUpdate(patscheUpdate), nil
}

func (rep *MysqlPatscheRepository) Delete(dsID int, id int) (string, error) {
	rec := Patsche{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", bussiness.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", bussiness.ErrNotFound
	}

	return "Patient Schedule has been delete", nil

}

func (rep *MysqlPatscheRepository) PatscheByID(id int) (patsche.Domain, error) {

	var ds Patsche

	result := rep.Conn.Where("id = ?", id).First(&ds)

	if result.Error != nil {
		return patsche.Domain{}, result.Error
	}

	return ToDomain(ds), nil
}
