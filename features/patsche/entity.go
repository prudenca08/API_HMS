package patsche

import (
	"time"
)

type Domain struct {
	ID      int
	AdminID int
	Day     string
	Time    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AllPatsche() ([]Domain, error)
	Create(dsID int, domain *Domain) (Domain, error)
	Update(admID int, dsID int, domain *Domain) (Domain, error)
	Delete(dsID, id int) (string, error)
	PatscheByID(id int) (Domain, error)
}

type Repository interface {
	AllPatsche() ([]Domain, error)
	Create(dsID int, domain *Domain) (Domain, error)
	Update(admID, dsID int, domain *Domain) (Domain, error)
	Delete(dsID, id int) (string, error)
	PatscheByID(id int) (Domain, error)
}
