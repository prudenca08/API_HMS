package patient

import "time"

type Domain struct {
	ID        int
	AdminID   int
	Name      string
	NIK       string
	DOB       string
	Gender    string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AllPatient() ([]Domain, error)
	Create(patID int, domain *Domain) (Domain, error)
	Update(admID int, patID int, domain *Domain) (Domain, error)
	Delete(patID, id int) (string, error)
	PatientByID(id int) (Domain, error)
}

type Repository interface {
	AllPatient() ([]Domain, error)
	Create(patID int, domain *Domain) (Domain, error)
	Update(admID int, patID int, domain *Domain) (Domain, error)
	Delete(patID, id int) (string, error)
	PatientByID(id int) (Domain, error)
}

