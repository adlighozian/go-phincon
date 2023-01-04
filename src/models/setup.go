package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-belajar?parseTime=true"), &gorm.Config{
		NowFunc: func() time.Time {
			currentTime := time.Now()
			_, offset := currentTime.Zone()
			mysqlTime := currentTime.Add(time.Second * time.Duration(offset))
			return mysqlTime
		},
	})

	if err != nil {
		panic(err)
	}

	// database.AutoMigrate(&Product{})
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Company{})
	database.AutoMigrate(&LogsPresent{})
	database.AutoMigrate(&LogsLogin{})
	database.AutoMigrate(&Home{})
	DB = database
}
