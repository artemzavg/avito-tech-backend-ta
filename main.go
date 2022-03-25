package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Advert struct {
	gorm.Model
	Title       string
	Description string
	Price       uint64
	PhotoLinks  []PhotoLink
}

type PhotoLink struct {
	gorm.Model
	Link     string
	AdvertId uint
}

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION_STRING")))
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&Advert{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&PhotoLink{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.Create(&Advert{Title: "Title", Description: "descr", PhotoLinks: []PhotoLink{{Link: "1"}, {Link: "2"}}, Price: 100000})
}
