package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateUser(pool *pgxpool.Pool, firstname string, surname string, age int) {
	pool.QueryRow(
		context.Background(),
		"INSERT INTO users(firstname, surname, age) VALUES($1, $2, $3) RETURNING id, firstname, surname, age",
		firstname, surname, age,
	)
}

func GetUser(pool *pgxpool.Pool, id int32) map[string]interface{} {
	res := map[string]interface{}{}

	pool.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE id=$1 RETURNING id, firstname, surname, age",
		id,
	).Scan(res)

	return res
}

func DeleteUser(pool *pgxpool.Pool, id int32) {
	pool.QueryRow(
		context.Background(),
		"DELETE FROM users WHERE id=$1 RETURNING id, firstname, surname, age",
		id,
	)
}
