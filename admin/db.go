package admin

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (db *gorm.DB) {
	var err error

	db, err = gorm.Open(postgres.Open(os.Getenv("DB_PARAMS")))
	if err != nil {
		panic(err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)

	return
}

func initWebsiteData(db *gorm.DB) {
	var cnt int64
	if err := db.Table("page_builder_pages").Count(&cnt).Error; err != nil {
		panic(err)
	}

	if cnt == 0 {
		if err := db.Exec(initWebsiteSQL).Error; err != nil {
			panic(err)
		}
	}
}

func initMediaLibraryData(db *gorm.DB) {
	var cnt int64
	if err := db.Table("media_libraries").Count(&cnt).Error; err != nil {
		panic(err)
	}

	if cnt == 0 {
		if err := db.Exec(initMediaLibrarySQL).Error; err != nil {
			panic(err)
		}
	}
}
