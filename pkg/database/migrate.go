package database

import (
	"log"
)

func Migrate() {
	news := &NewsScrapResult{}
	if !DB.Migrator().HasTable(news) {
		err := DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(NewsScrapResult{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
