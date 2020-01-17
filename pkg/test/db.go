package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var fmtDBString = "host=localhost user=vanjiii dbname='%s' sslmode=disable search_path='%s'"

func NewDB(t *testing.T) (*gorm.DB, func()) {
	schema := "test_" + uuid.New().String()
	pconn := gormString("kookbook", schema)

	db, err := gorm.Open("postgres", pconn)
	cleanup := func() {
		_, err := db.Raw(fmt.Sprintf("DROP SCHEMA %q CASCADE", schema)).Rows()
		db.Close()
		if err != nil {
			t.Fatalf("fail to drop test database: %q! Clean it up manually!!!", err)
		}
	}

	if err != nil {
		t.Fatalf("fail to create a db connection: %v", err)
	}

	if _, err = db.Raw(fmt.Sprintf("CREATE SCHEMA %q", schema)).Rows(); err != nil {
		t.Fatalf("Create schema failed: %s", err)
	}

	if err := dbMigrate(pconn); err != nil {
		cleanup()
		t.Fatalf("Running goose failed: %q", err)
	}

	return db, cleanup
}

func dbMigrate(pconn string) error {
	var dirname, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get root dir: %w", err)
	}

	mpath := filepath.Join(dirname, "..", "migrations")

	goose := exec.Command("goose", "-dir", mpath, "postgres", pconn, "up")
	return goose.Run()
}

func gormString(dbname, schema string) string {
	return fmt.Sprintf(fmtDBString, dbname, schema)
}
