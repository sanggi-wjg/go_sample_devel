package database

import (
	"log"
)

func MigrateTables() {
	news := &NewsScrapResult{}
	if !db.Migrator().HasTable(news) {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(NewsScrapResult{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
