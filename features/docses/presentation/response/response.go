package response

import (
	"finalproject/features/docses"
	"time"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

type CreateDocsesResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Day       string    `json:"day"`
	Time      string    `json:"time"`
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

func FromDomainCreate(domain docses.Domain) CreateDocsesResponse {
	return CreateDocsesResponse{
		Message:   "Create Doctor Session Success",
		ID:        domain.ID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type DocsesResponse struct {
	ID        int       `json:"id:"`
	Day       string    `json:"day"`
	Time      string    `json:"time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainAllDocses(domain docses.Domain) DocsesResponse {
	return DocsesResponse{
		ID:        domain.ID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdateDocses(domain docses.Domain) CreateDocsesResponse {
	return CreateDocsesResponse{
		Message:   "Update Doctor Session Success",
		ID:        domain.ID,
		Day:       domain.Day,
		Time:      domain.Time,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDocsesListDomain(domain []docses.Domain) []DocsesResponse {
	var response []DocsesResponse
	for _, value := range domain {
		response = append(response, FromDomainAllDocses(value))
	}
	return response
}
