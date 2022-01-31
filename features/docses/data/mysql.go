package data

import (
	"finalproject/features/docses"

	"finalproject/features/docses/bussiness"

	"gorm.io/gorm"
)

type MysqlDocsesRepository struct {
	Conn *gorm.DB
}

func NewMysqlDocsesRepository(conn *gorm.DB) docses.Repository {
	return &MysqlDocsesRepository{
		Conn: conn,
	}
}

func (rep *MysqlDocsesRepository) Create(dsID int, domain *docses.Domain) (docses.Domain, error) {

	dss := fromDomain(*domain)

	dss.AdminID = dsID

	result := rep.Conn.Create(&dss)

	if result.Error != nil {
		return docses.Domain{}, result.Error
	}

	return ToDomain(dss), nil

}

func (rep *MysqlDocsesRepository) AllDocses() ([]docses.Domain, error) {

	var ds []Docses

	result := rep.Conn.Find(&ds)

	if result.Error != nil {
		return []docses.Domain{}, result.Error
	}

	return toDomainList(ds), nil

}

func (rep *MysqlDocsesRepository) Update(admID int, dsID int, domain *docses.Domain) (docses.Domain, error) {
	docsesUpdate := fromDomain(*domain)

	docsesUpdate.ID = dsID
	result := rep.Conn.Where("id = ?", dsID).Updates(&docsesUpdate)

	if result.Error != nil {
		return docses.Domain{}, bussiness.ErrNotFound
	}

	return toDomainUpdate(docsesUpdate), nil
}

func (rep *MysqlDocsesRepository) Delete(dsID int, id int) (string, error) {
	rec := Docses{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", bussiness.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", bussiness.ErrNotFound
	}

	return "Doctor Session has been delete", nil

}

func (rep *MysqlDocsesRepository) DocsesByID(id int) (docses.Domain, error) {

	var ds Docses

	result := rep.Conn.Where("id = ?", id).First(&ds)

	if result.Error != nil {
		return docses.Domain{}, result.Error
	}

	return ToDomain(ds), nil
}
