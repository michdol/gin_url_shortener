package main

import (
	"fmt"
	"url_shortener/Config"
	"url_shortener/Models"
	"url_shortener/Routes"
	"github.com/jinzhu/gorm"
)


var err error

func main() {
	// Connect to DB
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	Config.DB.AutoMigrate(&Models.Url{})

	r := Routes.SetupRouter()

	r.Run()
}