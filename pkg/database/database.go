package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

type EnvDatabase struct {
	User, Password, Host, Port, DatabaseName string
}

func Setup(e *EnvDatabase) error {
	conn, err := gorm.Open(
		mysql.Open(getDatabaseDSN(e)),
		&gorm.Config{
			Logger: getGormLogger(),
		},
	)
	if err != nil {
		return err
	}

	db = conn
	MigrateTables()
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func getDatabaseDSN(e *EnvDatabase) string {
	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		e.User, e.Password, e.Host, e.Port, e.DatabaseName,
	)
	fmt.Println(str)
	return str
}

func getGormLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Able color
		},
	)
}
