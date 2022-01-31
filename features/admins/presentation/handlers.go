package presentation

import (
	"finalproject/features/admins"
	"finalproject/features/admins/presentation/request"
	"finalproject/features/admins/presentation/response"
	"net/http"
	"finalproject/features/admins/bussiness"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminHand admins.Service
}

func NewHandlerAdmin(serv admins.Service) *AdminHandler {
	return &AdminHandler{
		adminHand: serv,
	}
}

func (ctrl *AdminHandler) Register(c echo.Context) error {

	registerReq := request.Admins{}

	if err := c.Bind(&registerReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.adminHand.Register(registerReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *AdminHandler) Login(c echo.Context) error {

	loginReq := request.AdminLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.adminHand.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		if err == bussiness.ErrEmail{
			return response.NewErrorResponse(c, http.StatusNotFound, err)
		}else if err == bussiness.ErrPass{
			return response.NewErrorResponse(c, http.StatusUnauthorized, err)
		}
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, response.FromDomainLogin(result))
}
