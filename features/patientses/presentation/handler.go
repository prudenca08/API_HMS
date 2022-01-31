package presentation

import (
	"finalproject/features/patientses"
	"finalproject/features/patientses/presentation/request"
	"finalproject/features/patientses/presentation/response"
	"finalproject/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PatientsesHandler struct {
	patientsesService patientses.Service
}

func NewHandlerPatientses(serv patientses.Service) *PatientsesHandler{
	return &PatientsesHandler{
		patientsesService: serv,
	}
}

func (ctrl *PatientsesHandler) Create(c echo.Context) error {
	createReq := request.Patientses{}

	if err := c.Bind(&createReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	jwtGetID := middleware.GetUser(c)
	result, err := ctrl.patientsesService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadGateway, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainCreate(result))
}

func (ctrl *PatientsesHandler) AllPatientses(c echo.Context) error {
	
	result, err := ctrl.patientsesService.AllPatientses()
	
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
		
	}
	
	return response.NewSuccessResponse(c, response.FromPatientsesListDomain(result))
}
func (ctrl *PatientsesHandler) Update(c echo.Context) error{

	updateReq := request.Patientses{}

	if err := c.Bind(&updateReq); err !=nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	id,_ := strconv.Atoi(c. Param("id"))
	jwtGetID := middleware.GetUser(c)
	result, err := ctrl.patientsesService.Update(jwtGetID.ID, id, updateReq.ToDomain())
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadGateway, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainUpdatePatientses(result))
}

func (ctrl *PatientsesHandler) Delete(c echo.Context) error {
	dssID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))
	result, err := ctrl.patientsesService.Delete(dssID.ID, deletedId)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, result)
}

func (ctrl *PatientsesHandler) PatientsesByID(c echo.Context) error {
	patientsesID, _ := strconv.Atoi(c.Param("id"))
	result, err := ctrl.patientsesService.PatientsesByID(patientsesID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainAllPatientses(result))
}