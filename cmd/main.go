package main

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/tobidalhof/sqlc-example/db/sqlc"

	_ "github.com/lib/pq"
)

func run() error {
	conn, err := sql.Open("postgres", "host=postgres user=default password=secret dbname=default sslmode=disable")
	if err != nil {
		return err
	}

	ctx := context.Background()
	queries := db.New(conn)
	account, err := queries.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    "Test",
		Balance:  0,
		Currency: "USD",
	})
	if err != nil {
		return err
	}
	fmt.Println(account)

	account, err = queries.AddAccountBalance(ctx, db.AddAccountBalanceParams{
		ID:     account.ID,
		Amount: 100,
	})
	if err != nil {
		return err
	}
	fmt.Println(account)

	return nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
