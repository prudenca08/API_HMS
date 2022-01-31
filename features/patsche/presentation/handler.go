package presentation

import (
	"finalproject/features/patsche"
	"finalproject/features/patsche/presentation/request"
	"finalproject/features/patsche/presentation/response"
	"finalproject/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PatscheHandler struct {
	patscheService patsche.Service
}

func NewHandlerPatsche(serv patsche.Service) *PatscheHandler {
	return &PatscheHandler{
		patscheService: serv,
	}
}

func (ctrl *PatscheHandler) Create(c echo.Context) error {

	createReq := request.Patsche{}

	if err := c.Bind(&createReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.patscheService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *PatscheHandler) AllPatsche(c echo.Context) error {

	result, err := ctrl.patscheService.AllPatsche()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromPatscheListDomain(result))

}

func (ctrl *PatscheHandler) Update(c echo.Context) error {

	updateReq := request.Patsche{}

	if err := c.Bind(&updateReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.patscheService.Update(jwtGetID.ID, id, updateReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainUpdatePatsche(result))

}

func (ctrl *PatscheHandler) Delete(c echo.Context) error {

	dssID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.patscheService.Delete(dssID.ID, deletedId)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, result)

}

func (ctrl *PatscheHandler) PatscheByID(c echo.Context) error {

	patscheID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.patscheService.PatscheByID(patscheID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainAllPatsche(result))
}
