package domain

import "g/go/allsports/errs"

type Team struct {
	Id      string  `json:"team_id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Coach   Coach   `json:"coach"`
	Player  Player  `json:"players"`
}

type Address struct {
	Address1 string `json:"adddress_line1"`
	Address2 string `json:"adddress_line2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string `json:"zip_code"`
}

type Coach struct {
	Id           string  `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Gender       string  `json:"gender"`
	PhoneNumber  string  `json:"phone_number"`
	EmailAddress string  `json:"email_address"`
	Address      Address `json:"address"`
	Team         string  `json:"team_name"`
	Role         string  `json:"role"`
}

type Assistant struct {
	Id          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gender      string  `json:"gender"`
	PhoneNumber string  `json:"phone_number"`
	Address     Address `json:"address"`
}

type TeamRepository interface {
	FindAll() ([]Team, *errs.AppError)
	ByName(name string) (*Team, *errs.AppError)
}
