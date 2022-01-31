package presentation

import (
	"finalproject/features/patient"
	"finalproject/features/patient/presentation/request"
	"finalproject/features/patient/presentation/response"
	"finalproject/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PatientHandler struct {
	patientService patient.Service
}

func NewHandlerPatient(serv patient.Service) *PatientHandler {
	return &PatientHandler{
		patientService: serv,
	}
}

func (ctrl *PatientHandler) Create(c echo.Context) error {

	createReq := request.Patient{}

	if err := c.Bind(&createReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.patientService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *PatientHandler) AllPatient(c echo.Context) error {

	result, err := ctrl.patientService.AllPatient()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromPatientListDomain(result))

}

func (ctrl *PatientHandler) Update(c echo.Context) error {

	updateReq := request.Patient{}

	if err := c.Bind(&updateReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.patientService.Update(jwtGetID.ID, id, updateReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainUpdatePatient(result))

}

func (ctrl *PatientHandler) Delete(c echo.Context) error {

	patsID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.patientService.Delete(patsID.ID, deletedId)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, result)

}

func (ctrl *PatientHandler) PatientByID(c echo.Context) error {

	patientID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.patientService.PatientByID(patientID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainAllPatient(result))
}
