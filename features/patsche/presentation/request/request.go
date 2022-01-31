package request

import "finalproject/features/patsche"

type Patsche struct {
	Day  string `json:"day"`
	Time string `json:"time"`
}

func (req *Patsche) ToDomain() *patsche.Domain {
	return &patsche.Domain{
		Day:  req.Day,
		Time: req.Time,
	}
}
