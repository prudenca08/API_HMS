package response

import (
	"finalproject/features/doctor"
	"time"

	"net/http"

	docsesResp "finalproject/features/docses/presentation/response"

	echo "github.com/labstack/echo/v4"
)

type DoctorRegisterResponse struct {
	Message      string    `json:"message"`
	ID           int       `json:"id:"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	NIP          string    `json:"nip"`
	Experience   string    `json:"experience"`
	Spesialist   string    `json:"specialist"`
	Room         string    `json:"room"`
	Phone_Number string    `json:"phone_number"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type DoctorResponse struct {
	ID            int                       `json:"id:"`
	Username      string                    `json:"username"`
	Password      string                    `json:"password"`
	Name          string                    `json:"name"`
	NIP           string                    `json:"nip"`
	Experience    string                    `json:"experience"`
	Spesialist    string                    `json:"specialist"`
	Room          string                    `json:"room"`
	Phone_Number  string                    `json:"phone_number"`
	Status        string                    `json:"status"`
	CreatedAt     time.Time                 `json:"created_at"`
	UpdatedAt     time.Time                 `json:"updated_at"`
	DoctorSession docsesResp.DocsesResponse `json:"doctor_session"`
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

func NewErrorResponse1(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Password Not Same"

	return c.JSON(status, response)
}

func FromDomainRegister(domain doctor.Domain) DoctorRegisterResponse {
	return DoctorRegisterResponse{
		Message:      "Register Doctor Success",
		ID:           domain.ID,
		Username:     domain.Username,
		Password:     domain.Password,
		Name:         domain.Name,
		NIP:          domain.NIP,
		Experience:   domain.Experience,
		Spesialist:   domain.Specialist,
		Room:         domain.Room,
		Phone_Number: domain.Phone_Number,
		Status:       domain.Status,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

type DoctorLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	ID      int    `json:"id"`
}

type DoctorCPResponse struct {
	ID        int       `json:"id:"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainLogin(domain doctor.Domain) DoctorLoginResponse {
	return DoctorLoginResponse{
		Message: "Doctor Login Success",
		Token:   domain.Token,
		ID:      domain.ID,
	}
}
func FromDomainUpdateDoctor(domain doctor.Domain) DoctorRegisterResponse {
	return DoctorRegisterResponse{
		Message:      "Update Doctor Success",
		ID:           domain.ID,
		Username:     domain.Username,
		Name:         domain.Name,
		NIP:          domain.NIP,
		Experience:   domain.Experience,
		Spesialist:   domain.Specialist,
		Room:         domain.Room,
		Phone_Number: domain.Phone_Number,
		Status:       domain.Status,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromDomainUpdatePassword(domain doctor.Domain) DoctorCPResponse {
	return DoctorCPResponse{
		ID:        domain.ID,
		Message:   "Update Password Success",
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainAllDoctor(domain doctor.Domain) DoctorResponse {
	return DoctorResponse{
		ID:            domain.ID,
		Username:      domain.Username,
		Password:      domain.Password,
		Name:          domain.Name,
		NIP:           domain.NIP,
		Experience:    domain.Experience,
		Spesialist:    domain.Specialist,
		Room:          domain.Room,
		Phone_Number:  domain.Phone_Number,
		Status:        domain.Status,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DoctorSession: docsesResp.FromDomainAllDocses(domain.DoctorSession),
	}
}

func FromDoctorListDomain(domain []doctor.Domain) []DoctorResponse {
	var response []DoctorResponse
	for _, value := range domain {
		response = append(response, FromDomainAllDoctor(value))
	}
	return response
}
