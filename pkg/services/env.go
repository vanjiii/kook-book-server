package services

import (
	"fmt"
	"log"

	"vanjiii/kook-book-server/pkg/auth"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var fmtDBString = "host=localhost user=vanjiii dbname=%s sslmode=disable"

type Env struct {
	AuthHandler *auth.Handler
}

func NewEnv() (*Env, error) {
	db, err := gorm.Open("postgres", gormString("kookbook"))
	if err != nil {
		log.Fatalf("fail to create a db connection: %v", err)
	}

	return newEnv(db), err
}

func newEnv(db *gorm.DB) *Env {
	ah := auth.NewHandler(db)

	return &Env{
		AuthHandler: ah,
	}
}

func gormString(dbname string) string {
	return fmt.Sprintf(fmtDBString, dbname)
}
