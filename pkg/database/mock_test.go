package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
	//newsScrapResult NewsScrapResult
}

func createTestSuite() *Suite {
	suite := &Suite{}
	var (
		db  *sql.DB
		err error
	)
	db, suite.mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("fail to create sql mock: %v", err)
	}
	if db == nil {
		log.Fatal("suite db is null")
	}
	if suite.mock == nil {
		log.Fatal("suite mock is null")
	}

	//dialect := mysql.New(mysql.Config{
	//	DSN:                       "sqlmock_db_0",
	//	DriverName:                "mysql",
	//	Conn:                      db,
	//	ServerVersion:             "8.0.27",
	//	SkipInitializeWithVersion: true,
	//})
	//suite.db, err = gorm.Open(dialect, &gorm.Config{})
	suite.db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("fail to open gorm db: %v", err)
	}
	if suite.db == nil {
		log.Fatal("gorm db is null")
	}
	//defer db.Close()
	return suite
}

func TestCreateTestSuite(t *testing.T) {
	createTestSuite()
}
