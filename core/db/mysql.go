package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	godotenv.Load()
	MYSQL_DSN := "root:passwordhere@tcp(localhost:13306)/fulfillment?parseTime=true"
	//dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(MYSQL_DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&model.Product{})
	return db
}
