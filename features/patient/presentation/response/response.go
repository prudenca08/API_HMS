package response

import (
	"finalproject/features/patient"
	"time"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

type CreatePatientResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	NIK       string    `json:"nik"`
	DOB       string    `json:"dob"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}

func FromDomainCreate(domain patient.Domain) CreatePatientResponse {
	return CreatePatientResponse{
		Message:   "Create Patient Success",
		ID:        domain.ID,
		Name:      domain.Name,
		NIK:       domain.NIK,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type PatientResponse struct {
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	NIK       string    `json:"nik"`
	DOB       string    `json:"dob"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainAllPatient(domain patient.Domain) PatientResponse {
	return PatientResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		NIK:       domain.NIK,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdatePatient(domain patient.Domain) CreatePatientResponse {
	return CreatePatientResponse{
		Message:   "Update Patient Success",
		ID:        domain.ID,
		Name:      domain.Name,
		NIK:       domain.NIK,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		Address:   domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromPatientListDomain(domain []patient.Domain) []PatientResponse {
	var response []PatientResponse
	for _, value := range domain {
		response = append(response, FromDomainAllPatient(value))
	}
	return response
}
