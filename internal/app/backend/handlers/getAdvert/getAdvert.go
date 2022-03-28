package getAdvert

import (
	"encoding/json"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend/dbContext"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type Handler struct {
	Db *gorm.DB
}

type request struct {
	Id     uint     `json:"Id"`
	Fields []string `json:"Fields,omitempty"`
}

type response struct {
	Title         string   `json:"Title"`
	Price         uint64   `json:"Price"`
	MainPhotoLink string   `json:"MainPhotoLink"`
	Description   string   `json:"Description,omitempty"`
	AllPhotos     []string `json:"AllPhotos,omitempty"`
}

func (handler Handler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	var requestBody request

	err := decodeRequest(w, r, &requestBody)
	if err != nil {
		return
	}

	var advert dbContext.Advert

	result := handler.Db.First(&advert, requestBody.Id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	var photos []dbContext.Photo
	err = handler.Db.Model(&advert).Association("Photos").Find(&photos)
	if err != nil {
		return
	}

	response := response{Title: advert.Title, Price: advert.Price, MainPhotoLink: photos[0].Link}

	if contains(requestBody.Fields, "description") {
		response.Description = advert.Description
	}

	if contains(requestBody.Fields, "photos") {
		response.AllPhotos = dbContext.Links(photos)
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

func contains(fields []string, value string) bool {
	for _, s := range fields {
		if strings.ToLower(s) == value {
			return true
		}
	}

	return false
}
