package services

import (
	"log"

	"vanjiii/kook-book-server/pkg/auth"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Env struct {
	AuthHandler *auth.Handler
}

func NewEnv() (*Env, error) {
	db, err := gorm.Open("postgres", "host=localhost user=vanjiii dbname=kookbook sslmode=disable")
	if err != nil {
		log.Fatalf("fail to create a db connection: %v", err)
	}

	ah := auth.NewHandler(db)

	return &Env{
		AuthHandler: ah,
	}, err
}
