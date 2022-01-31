package request

import "finalproject/features/patientses"

type Patientses struct {

	DoctorID          int    `json:"doctorid"`
	PatientID         int    `json:"patientid"`
	PatientScheduleID int    `json:"patientscheduleid"`
	Date              string `json:"date"`
	Symptoms          string `json:"symptoms"`
	Title             string `json:"title"`
	DetailRecipe      string `json:"detailrecipe"`
	Status            string `json:"status"`

}

func (req *Patientses) ToDomain() *patientses.Domain {
	return &patientses.Domain{

		DoctorID:          req.DoctorID,
		PatientID:         req.PatientID,
		PatientScheduleID: req.PatientScheduleID,
		Date:              req.Date,
		Symptoms:          req.Symptoms,
		Title:             req.Title,
		DetailRecipe:      req.DetailRecipe,
		Status:            req.Status,

	}
}
