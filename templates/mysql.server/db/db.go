package db

import (
	"database/sql"
	"project-name/model"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
}

type MySQLDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return MySQLDB{db: db}
}

func (d MySQLDB) GetTechnologies() ([]*model.Technology, error) {
	rows, err := d.db.Query("select name, details from technologies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tech []*model.Technology
	for rows.Next() {
		t := new(model.Technology)
		err = rows.Scan(&t.Name, &t.Details)
		if err != nil {
			return nil, err
		}
		tech = append(tech, t)
	}
	return tech, nil
}
