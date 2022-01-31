package presentation

import (
	"finalproject/features/docses"
	"finalproject/features/docses/presentation/request"
	"finalproject/features/docses/presentation/response"
	"finalproject/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DocsesHandler struct {
	docsesService docses.Service
}

func NewHandlerDocses(serv docses.Service) *DocsesHandler {
	return &DocsesHandler{
		docsesService: serv,
	}
}

func (ctrl *DocsesHandler) Create(c echo.Context) error {

	createReq := request.Docses{}

	if err := c.Bind(&createReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.docsesService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *DocsesHandler) AllDocses(c echo.Context) error {

	result, err := ctrl.docsesService.AllDocses()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDocsesListDomain(result))

}

func (ctrl *DocsesHandler) Update(c echo.Context) error {

	updateReq := request.Docses{}

	if err := c.Bind(&updateReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.docsesService.Update(jwtGetID.ID, id, updateReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainUpdateDocses(result))

}

func (ctrl *DocsesHandler) Delete(c echo.Context) error {

	dssID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.docsesService.Delete(dssID.ID, deletedId)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, result)

}

func (ctrl *DocsesHandler) DocsesByID(c echo.Context) error {

	docsesID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.docsesService.DocsesByID(docsesID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainAllDocses(result))
}
