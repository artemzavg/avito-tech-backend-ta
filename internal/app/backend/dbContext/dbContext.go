package dbContext

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Advert struct {
	gorm.Model
	Title       string
	Description string
	Price       uint64
	AdvertLinks []PhotoLink
}

type PhotoLink struct {
	gorm.Model
	Link     string
	AdvertId uint
}

func NewContext(connectionString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Advert{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&PhotoLink{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetLinks(links []string) []PhotoLink {
	photoLinks := make([]PhotoLink, len(links), len(links))

	for i, link := range links {
		photoLinks[i] = PhotoLink{Link: link}
	}

	return photoLinks
}
