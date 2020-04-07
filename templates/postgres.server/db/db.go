package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"project-name/model"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
}

type PostgresDB struct {
	conn *pgx.Conn
}

func NewDB(conn *pgx.Conn) DB {
	return PostgresDB{conn: conn}
}

func (d PostgresDB) GetTechnologies() ([]*model.Technology, error) {
	rows, err := d.conn.Query(
		context.Background(),
		"select name, details from technologies",
	)
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
