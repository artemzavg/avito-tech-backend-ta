package createAdvert

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
	TitleMaxLen       = 200
	DescriptionMaxLen = 1000
	LinksMaxAmount    = 3
)

type request struct {
	Title       string   `json:"Title"`
	Description string   `json:"Description"`
	Price       uint64   `json:"Price"`
	AdvertLinks []string `json:"Photos"`
}

type response struct {
	Id   uint `json:"Id"`
	Code uint `json:"Code"`
}

func (handler Handler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	var request request

	err := decodeRequest(w, r, &request)
	if err != nil {
		return
	}

	log.Println(fmt.Sprintf("Got createAdvertRequest %v", request))

	if !request.isValid(w) {
		return
	}

	photos := dbContext.Photos(request.AdvertLinks)
	advert := dbContext.Advert{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Photos:      photos,
	}
	result := handler.Db.Create(&advert)
	response := response{Id: advert.ID, Code: 0}
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error in creating model. %v", result.Error))
		response.Code = 1
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func decodeRequest(w http.ResponseWriter, r *http.Request, requestBody *request) error {
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func (requestBody *request) isValid(w http.ResponseWriter) bool {
	if !isStringValid(requestBody.Title, TitleMaxLen) {
		log.Println("Error! Title is invalid")
		http.Error(w, fmt.Sprintf("Title is invalid. It should not be empty and should be less than %d symbols length", TitleMaxLen), http.StatusBadRequest)
		return false
	}

	if !isStringValid(requestBody.Description, DescriptionMaxLen) {
		log.Println("Error! Description is invalid")
		http.Error(w, fmt.Sprintf("Description is invalid. It should not be empty and should be less than %d symbols length", DescriptionMaxLen), http.StatusBadRequest)
		return false
	}

	if len(requestBody.AdvertLinks) > LinksMaxAmount {
		log.Println("Error! Too many photo links")
		http.Error(w, fmt.Sprintf("Request should contain less than %d photo links", LinksMaxAmount), http.StatusBadRequest)
		return false
	}
	return true
}

func isStringValid(title string, maxLen int) bool {
	titleLen := len(title)

	if titleLen > maxLen {
		return false
	}

	if titleLen == 0 {
		return false
	}

	return true
}
