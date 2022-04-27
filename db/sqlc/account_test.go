package db

import (
	"context"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	ctx := context.Background()
	_, err := queries.CreateAccount(ctx, CreateAccountParams{
		Owner:    "Test",
		Balance:  100,
		Currency: "USD",
	})
	if err != nil {
		t.Errorf("could not create account: %s", err)
	}
}
