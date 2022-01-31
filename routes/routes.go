package routes

import (
	business "finalproject/features/admins/bussiness"
	admins "finalproject/features/admins/presentation"
	controller "finalproject/features/admins/presentation/response"

	bussiness "finalproject/features/doctor/bussiness"
	doctor "finalproject/features/doctor/presentation"
	response "finalproject/features/doctor/presentation/response"

	docses "finalproject/features/docses/presentation"

	patsche "finalproject/features/patsche/presentation"

	patient "finalproject/features/patient/presentation"

	patientses "finalproject/features/patientses/presentation"

	middlewareApp "finalproject/middleware"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteList struct {
	JWTMiddleware    middleware.JWTConfig
	AdminRouter      admins.AdminHandler
	DoctorRouter     doctor.DoctorHandler
	DocsesRouter     docses.DocsesHandler
	PatientRouter    patient.PatientHandler
	PatientsesRouter patientses.PatientsesHandler
	PatscheRouter    patsche.PatscheHandler
}

func (cl *RouteList) RouteRegister(e *echo.Echo) {
	// Admins
	admins := e.Group("admins")
	admins.POST("/register", cl.AdminRouter.Register)
	admins.POST("/login", cl.AdminRouter.Login)
	admins.PUT("/update-doctor/:id", cl.DoctorRouter.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.DELETE("/delete-doctor/:id", cl.DoctorRouter.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.PUT("/change-password-doctor/:id", cl.DoctorRouter.ChangePass, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	admins.POST("/create-patsche", cl.PatscheRouter.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.PUT("/update-patsche/:id", cl.PatscheRouter.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.DELETE("/delete-patsche/:id", cl.PatscheRouter.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	admins.POST("/create-docses", cl.DocsesRouter.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.PUT("/update-docses/:id", cl.DocsesRouter.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.DELETE("/delete-docses/:id", cl.DocsesRouter.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	admins.POST("/create-patientses", cl.PatientsesRouter.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.PUT("/update-patientses/:id", cl.PatientsesRouter.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.DELETE("/delete-patientses/:id", cl.PatientsesRouter.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	admins.POST("/create-patient", cl.PatientRouter.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.PUT("/update-patient/:id", cl.PatientRouter.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	admins.DELETE("/delete-patient/:id", cl.PatientRouter.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())

	// Doctor
	doctor := e.Group("doctor")
	doctor.POST("/register", cl.DoctorRouter.Register, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationAdmin())
	doctor.POST("/login", cl.DoctorRouter.Login)
	doctor.PUT("/update-doctor/:id", cl.DoctorRouter.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctor())
	doctor.PUT("/change-password/:id", cl.DoctorRouter.ChangePass, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctor())
	doctor.PUT("/update-patientses/:id", cl.PatientsesRouter.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationDoctor())


	//Doctor
	e.GET("/doctor", cl.DoctorRouter.AllDoctor)
	e.GET("/doctor/:id", cl.DoctorRouter.DoctorByID)

	//Patients
	e.GET("/patient", cl.PatientRouter.AllPatient)
	e.GET("/patient/:id", cl.PatientRouter.PatientByID)

	//Patientsche
	e.GET("/patsche", cl.PatscheRouter.AllPatsche)
	e.GET("/patsche/:id", cl.PatscheRouter.PatscheByID)

	//Docses
	e.GET("/docses", cl.DocsesRouter.AllDocses)
	e.GET("/docses/:id", cl.DocsesRouter.DocsesByID)

	//Patientses
	e.GET("/patientses", cl.PatientsesRouter.AllPatientses)
	e.GET("/patientses/:id", cl.PatientsesRouter.PatientsesByID)

}

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "admin" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}
func RoleValidationDoctor() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "doctor" {
				return hf(c)
			} else {
				return response.NewErrorResponse(c, http.StatusForbidden, bussiness.ErrUnathorized)
			}
		}
	}
}
