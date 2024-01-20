package config

import (
	"github.com/a-fandy/finan/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	const dsn = "root:laravel@tcp(localhost:3306)/finan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	exception.PanicIfError(err)

	//autoMigrate
	// err = db.AutoMigrate(&entity.User{})
	// if err != nil {
	// 	log.Println(err)
	// }

	return db
}
