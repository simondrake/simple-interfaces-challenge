package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"cd.splunkdev.com/sdrake/simple-interface-test/database"
)

func main() {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "admin", "password", "127.0.0.1", "5432", "users")

	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	database.CreateUser(pool, "wibble", "wobble", 42)
}

func ProcessUserRequest(pool *pgxpool.Pool, firstname string, surname string, age int) {
	database.CreateUser(pool, "wibble", "wobble", 42)

	user := database.GetUser(pool, 1)

	fmt.Println(user)

	database.DeleteUser(pool, 1)
}
