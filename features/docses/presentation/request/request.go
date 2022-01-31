package request

import "finalproject/features/docses"

type Docses struct {
	Day  string `json:"day"`
	Time string `json:"time"`
}

func (req *Docses) ToDomain() *docses.Domain {
	return &docses.Domain{
		Day:  req.Day,
		Time: req.Time,
	}
}
