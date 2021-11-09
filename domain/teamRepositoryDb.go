package domain

import (
	"database/sql"
	"g/go/allsports/errs"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TeamRepositoryDb struct {
	client *sql.DB
}

func (team TeamRepositoryDb) ByName(name string) (*Team, *errs.AppError) {

	findTeamSql := "select * from teams t join coaches c join players p where t.name = ? and t.name = p.team and c.team = t.name"

	row := team.client.QueryRow(findTeamSql, name)
	var t Team
	err := row.Scan(&t.Id, &t.Name, &t.Address.Address1, &t.Address.Address2, &t.Address.City, &t.Address.State, &t.Address.Zipcode,
		&t.Coach.Id, &t.Coach.FirstName, &t.Coach.LastName, &t.Coach.Gender, &t.Coach.PhoneNumber, &t.Coach.EmailAddress,
		&t.Coach.Address.Address1, &t.Coach.Address.Address2, &t.Coach.Address.City, &t.Coach.Address.State, &t.Coach.Address.Zipcode, &t.Coach.Role, &t.Coach.Team,
		&t.Player.Id, &t.Player.FirstName, &t.Player.LastName, &t.Player.DateofBirth, &t.Player.Gender, &t.Player.PhoneNumber, &t.Player.EmailAddress, &t.Player.JerseNumber,
		&t.Player.Team, &t.Player.Address.Address1, &t.Player.Address.Address2, &t.Player.Address.City, &t.Player.Address.State, &t.Player.Address.Zipcode)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Team Not Found")
		} else {
			log.Println("Error while scanning team table " + err.Error())

			return nil, errs.NewUnexpectedError("Unexpected Error")
		}
	}

	return &t, nil
}

func (team TeamRepositoryDb) FindAll() ([]Team, *errs.AppError) {

	findAllSql := "select * from teams t join coaches c join players p where t.name = p.team and c.team = t.name"

	rows, err := team.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying team table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error")

	}

	teams := make([]Team, 0)
	for rows.Next() {
		var t Team
		err := rows.Scan(&t.Id, &t.Name, &t.Address.Address1, &t.Address.Address2, &t.Address.City, &t.Address.State, &t.Address.Zipcode,
			&t.Coach.Id, &t.Coach.FirstName, &t.Coach.LastName, &t.Coach.Gender, &t.Coach.PhoneNumber, &t.Coach.EmailAddress,
			&t.Coach.Address.Address1, &t.Coach.Address.Address2, &t.Coach.Address.City, &t.Coach.Address.State, &t.Coach.Address.Zipcode, &t.Coach.Role, &t.Coach.Team,
			&t.Player.Id, &t.Player.FirstName, &t.Player.LastName, &t.Player.DateofBirth, &t.Player.Gender, &t.Player.PhoneNumber, &t.Player.EmailAddress, &t.Player.JerseNumber,
			&t.Player.Team, &t.Player.Address.Address1, &t.Player.Address.Address2, &t.Player.Address.City, &t.Player.Address.State, &t.Player.Address.Zipcode)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Team Not Found")
			} else {
				log.Println("Error while scanning team table " + err.Error())
				return nil, errs.NewUnexpectedError("Unexpected Error")
			}
		}

		teams = append(teams, t)
	}

	return teams, nil
}

func NewTeamRepositoryDb() TeamRepositoryDb {
	client, err := sql.Open("mysql", "root:dayaxQ!@tcp(localhost:3306)/soccer")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return TeamRepositoryDb{client}
}
