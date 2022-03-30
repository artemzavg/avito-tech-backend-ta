package getAllAdverts

import (
	"encoding/json"
	"fmt"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/dbContext"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Handler struct {
	Db *gorm.DB
}

const (
	PageAdvertAmount = 10
)

type request struct {
	Page int `json:"Page,omitempty"`
	Sort struct {
		Type      string `json:"Type"`
		Direction string `json:"Direction"`
	} `json:"Sort,omitempty"`
}

type responseAdvert struct {
	Title     string `json:"Title"`
	MainPhoto string `json:"MainPhoto"`
	Price     uint64 `json:"Price"`
}

type response struct {
	Adverts []responseAdvert
}

func (handler Handler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	var requestBody request

	err := decodeRequest(w, r, &requestBody)
	if err != nil {
		return
	}

	log.Printf("Request: %v\n", requestBody)

	var advertsAmount int64
	handler.Db.Model(&dbContext.Advert{}).Count(&advertsAmount)

	if advertsAmount <= int64(PageAdvertAmount*(requestBody.Page-1)) {
		log.Println("Too large page is required")
		http.Error(w, "Incorrect page", http.StatusBadRequest)
		return
	}

	var direction string
	var orderBy string

	if requestBody.Sort.Type == "created_at" || requestBody.Sort.Type == "price" {
		orderBy = requestBody.Sort.Type
	} else {
		orderBy = "created_at"
	}

	if requestBody.Sort.Direction == "asc" || requestBody.Sort.Direction == "desc" {
		direction = requestBody.Sort.Direction
	} else {
		if orderBy == "created_at" {
			direction = "desc"
		} else {
			direction = "asc"
		}
	}

	var adverts []dbContext.Advert
	result := handler.Db.Limit(PageAdvertAmount).Offset(PageAdvertAmount * (requestBody.Page - 1)).Order(fmt.Sprintf("%s %s", orderBy, direction)).Find(&adverts)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var response response
	response.Adverts = make([]responseAdvert, len(adverts))
	for i, advert := range adverts {
		var photos []dbContext.Photo
		err := handler.Db.Model(&advert).Association("Photos").Find(&photos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		response.Adverts[i] = responseAdvert{Title: advert.Title, Price: advert.Price, MainPhoto: photos[0].Link}
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func decodeRequest(w http.ResponseWriter, r *http.Request, request *request) error {
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
