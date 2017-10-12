package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sunho0301/info344-in-class/zipsvr/models"
)

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL:   /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):]
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 {
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(headerAccessControllAllowOrigin, "*")
	// find zip
	zips := ch.Index[cityName]
	// put zips inside the json
	json.NewEncoder(w).Encode(zips)
}
