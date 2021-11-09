package domain

import (
	"database/sql"
	"g/go/allsports/errs"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type PlayerRepositoryDb struct {
	client *sql.DB
}

func (pl PlayerRepositoryDb) ById(id string) (*Player, *errs.AppError) {
	findPlayerSql := "select * from players where player_id = ?"

	row := pl.client.QueryRow(findPlayerSql, id)
	var p Player
	err := row.Scan(&p.Id, &p.FirstName, &p.LastName, &p.DateofBirth, &p.Gender, &p.PhoneNumber, &p.EmailAddress, &p.JerseNumber,
		&p.Team, &p.Address.Address1, &p.Address.Address2, &p.Address.City, &p.Address.State, &p.Address.Zipcode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("player not found")
		} else {
			log.Println("Error while scanning players table " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error")
		}
	}

	return &p, nil
}

func (p PlayerRepositoryDb) FindAll() ([]Player, *errs.AppError) {

	findAllSql := "select * from players"

	rows, err := p.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying players table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	customers := make([]Player, 0)
	for rows.Next() {
		var p Player
		err := rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.DateofBirth, &p.Gender, &p.PhoneNumber, &p.EmailAddress, &p.JerseNumber,
			&p.Team, &p.Address.Address1, &p.Address.Address2, &p.Address.City, &p.Address.State, &p.Address.Zipcode)
		if err != nil {
			log.Println("Error while scanning players table " + err.Error())
		}
		customers = append(customers, p)
	}

	return customers, nil
}

func NewPlayerRepositoryDb() PlayerRepositoryDb {
	client, err := sql.Open("mysql", "root:dayaxQ!@tcp(localhost:3306)/soccer")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return PlayerRepositoryDb{client}
}
