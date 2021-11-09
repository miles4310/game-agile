package domain

import "g/go/allsports/errs"

type Player struct {
	Id           string  `json:"player_id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	DateofBirth  string  `json:"DOB"`
	Gender       string  `json:"gender"`
	PhoneNumber  string  `json:"phone_number"`
	EmailAddress string  `json:"email_address"`
	JerseNumber  string  `json:"jersey_number"`
	Team         string  `json:"team"`
	Address      Address `json:"address"`
}

type PlayerRepository interface {
	FindAll() ([]Player, *errs.AppError)
	ById(string) (*Player, *errs.AppError)
}
