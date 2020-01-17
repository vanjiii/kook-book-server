package auth

import (
	"context"
	"testing"

	"vanjiii/kook-book-server/pkg"
	"vanjiii/kook-book-server/pkg/test"
)

var ctx = context.Background()

func TestCreateUser(t *testing.T) {
	db, cleanup := test.NewDB(t)
	defer cleanup()

	h := NewHandler(db)

	email := "ivan@dim.org"

	err := h.CreateUser(ctx, email, "1234")
	if err != nil {
		t.Fatalf("create user fails with: %v", err)
	}

	var got pkg.User
	db.Where("email = ?", email).First(&got)

	if got.Email != email {
		t.Fatalf("expected: %q, got: %q", email, got.Email)
	}
}
