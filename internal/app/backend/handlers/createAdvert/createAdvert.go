package createAdvert

import (
	"encoding/json"
	"fmt"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/dbContext"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
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
	AdvertLinks []string `json:"AdvertLinks"`
}

type response struct {
	Id   uint `json:"Id"`
	Code uint `json:"Code"`
}

func (handler *Handler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	var requestBody request

	err := decodeRequest(w, r, &requestBody)
	if err != nil {
		return
	}

	log.Println(fmt.Sprintf("Got createAdvertRequest %v", requestBody))

	if !requestBody.isValid(w) {
		return
	}

	links := dbContext.GetLinks(requestBody.AdvertLinks)
	advert := dbContext.Advert{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Price:       requestBody.Price,
		AdvertLinks: links,
	}
	result := handler.Db.Create(&advert)
	response := response{Id: advert.ID, Code: 0}
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error in creating model. %v", err))
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
		processInvalidString(w, "title", TitleMaxLen)
		return false
	}

	if !isStringValid(requestBody.Description, DescriptionMaxLen) {
		processInvalidString(w, "description", DescriptionMaxLen)
		return false
	}

	if len(requestBody.AdvertLinks) > LinksMaxAmount {
		log.Println("Error! Too many photo links")
		http.Error(w, fmt.Sprintf("Request should contain less than %d photo links", LinksMaxAmount), http.StatusBadRequest)
		return false
	}
	return true
}

func processInvalidString(w http.ResponseWriter, name string, maxLen int) {
	log.Println("Error! Title is invalid")
	http.Error(w, fmt.Sprintf("%s is invalid. It should not be empty and should be less than %d symbols length", strings.ToTitle(name), maxLen), http.StatusBadRequest)
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
