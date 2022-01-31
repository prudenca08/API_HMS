package main

import (
	"log"

	_routes "finalproject/routes"

	_adminService "finalproject/features/admins/bussiness"
	_adminRepo "finalproject/features/admins/data"
	_adminController "finalproject/features/admins/presentation"

	_doctorService "finalproject/features/doctor/bussiness"
	_doctorRepo "finalproject/features/doctor/data"
	_doctorController "finalproject/features/doctor/presentation"

	_docsesService "finalproject/features/docses/bussiness"
	_docsesRepo "finalproject/features/docses/data"
	_docsesController "finalproject/features/docses/presentation"

	_patientsesService "finalproject/features/patientses/bussiness"
	_patientsesRepo "finalproject/features/patientses/data"
	_patientsesController "finalproject/features/patientses/presentation"

	_patscheService "finalproject/features/patsche/bussiness"
	_patscheRepo "finalproject/features/patsche/data"
	_patscheController "finalproject/features/patsche/presentation"

	_patientService "finalproject/features/patient/bussiness"
	_patientRepo "finalproject/features/patient/data"
	_patientController "finalproject/features/patient/presentation"

	_dbDriver "finalproject/config"

	_driverFactory "finalproject/drivers"

	_middleware "finalproject/middleware"

	corsm "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_adminRepo.Admins{},
		&_docsesRepo.Docses{},
	)
	db.AutoMigrate(
		&_doctorRepo.Doctor{},
	)
	db.AutoMigrate(
		&_patientRepo.Patient{},
		&_patscheRepo.Patsche{},
	)
	db.AutoMigrate(
		&_patientsesRepo.Patientses{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()
	e.Use(corsm.CORSWithConfig(corsm.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderAuthorization, echo.HeaderAccessControlAllowMethods},
	}))

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminService := _adminService.NewServiceAdmin(adminRepo, 10, &configJWT)
	adminCtrl := _adminController.NewHandlerAdmin(adminService)

	doctorRepo := _driverFactory.NewDoctorRepository(db)
	doctorService := _doctorService.NewServiceDoctor(doctorRepo, 10, &configJWT)
	doctorCtrl := _doctorController.NewHandlerDoctor(doctorService)

	docsesRepo := _driverFactory.NewDocsesRepository(db)
	docsesService := _docsesService.NewServiceDocses(docsesRepo)
	docsesCtrl := _docsesController.NewHandlerDocses(docsesService)

	patscheRepo := _driverFactory.NewPatscheRepository(db)
	patscheService := _patscheService.NewServicePatsche(patscheRepo)
	patscheCtrl := _patscheController.NewHandlerPatsche(patscheService)

	patientRepo := _driverFactory.NewPatientRepository(db)
	patientService := _patientService.NewServicePatient(patientRepo)
	patientCtrl := _patientController.NewHandlerPatient(patientService)

	patientsesRepo := _driverFactory.NewPatientsesRepository(db)
	patientsesService := _patientsesService.NewServicePatientses(patientsesRepo)
	patientsesCtrl := _patientsesController.NewHandlerPatientses(patientsesService)

	routesInit := _routes.RouteList{
		JWTMiddleware:    configJWT.Init(),
		AdminRouter:      *adminCtrl,
		DoctorRouter:     *doctorCtrl,
		DocsesRouter:     *docsesCtrl,
		PatientRouter:    *patientCtrl,
		PatientsesRouter: *patientsesCtrl,
		PatscheRouter:    *patscheCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
