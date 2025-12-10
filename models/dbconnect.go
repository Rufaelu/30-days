package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:root@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateDB(name string, conn *pgx.Conn) error {

	query := fmt.Sprintf("create table %s  (id Serial Primary key, name Text Not Null, age int not null);", name)

	_, err := conn.Exec(context.Background(), query)
	if err != nil {

		return err
	}

	return nil
}
