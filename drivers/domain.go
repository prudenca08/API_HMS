package drivers

import (
	"finalproject/features/admins"
	adminDB "finalproject/features/admins/data"
	"finalproject/features/docses"
	docsesDB "finalproject/features/docses/data"
	"finalproject/features/doctor"
	doctorDB "finalproject/features/doctor/data"
	"finalproject/features/patient"
	patientDB "finalproject/features/patient/data"
	"finalproject/features/patientses"
	patientsesDB "finalproject/features/patientses/data"
	"finalproject/features/patsche"
	patscheDB "finalproject/features/patsche/data"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) admins.Repository {
	return adminDB.NewMysqlAdminRepository(conn)

}
func NewDoctorRepository(conn *gorm.DB) doctor.Repository {
	return doctorDB.NewMysqlDoctorRepository(conn)
}
func NewDocsesRepository(conn *gorm.DB) docses.Repository {
	return docsesDB.NewMysqlDocsesRepository(conn)
}
func NewPatientRepository(conn *gorm.DB) patient.Repository {
	return patientDB.NewMysqlPatientRepository(conn)
}

func NewPatientsesRepository(conn *gorm.DB) patientses.Repository{
	return patientsesDB.NewMysqlPatientsesRepository(conn)
}
func NewPatscheRepository(conn *gorm.DB) patsche.Repository {
	return patscheDB.NewMysqlPatscheRepository(conn)
}