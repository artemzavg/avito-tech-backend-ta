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
	Photos      []Photo
}

type Photo struct {
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

	err = db.AutoMigrate(&Photo{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Photos(links []string) []Photo {
	photoLinks := make([]Photo, len(links), len(links))

	for i, link := range links {
		photoLinks[i] = Photo{Link: link}
	}

	return photoLinks
}

func Links(photos []Photo) []string {
	links := make([]string, len(photos), len(photos))

	for i, link := range photos {
		links[i] = link.Link
	}

	return links
}
