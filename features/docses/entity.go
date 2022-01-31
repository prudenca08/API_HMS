package docses

import "time"

type Domain struct {
	ID      int
	AdminID int
	Day     string
	Time    string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AllDocses() ([]Domain, error)
	Create(dsID int, domain *Domain) (Domain, error)
	Update(admID int, dsID int, domain *Domain) (Domain, error)
	Delete(dsID, id int) (string, error)
	DocsesByID(id int) (Domain, error)
}

type Repository interface {
	AllDocses() ([]Domain, error)
	Create(dsID int, domain *Domain) (Domain, error)
	Update(admID, dsID int, domain *Domain) (Domain, error)
	Delete(dsID, id int) (string, error)
	DocsesByID(id int) (Domain, error)
}
