package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Location struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state,omitempty"`
}

func getCoordinates(cityName, apiKey string) (*Location, error) {
	baseURL := "http://api.openweathermap.org/geo/1.0/direct"
	reqURL := fmt.Sprintf("%s?q=%s&limit=1&appid=%s", baseURL, url.QueryEscape(cityName), apiKey)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("code de statut d'erreur : %d", resp.StatusCode)
	}

	var locations []Location
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return nil, fmt.Errorf("erreur de décodage JSON : %v", err)
	}

	if len(locations) == 0 {
		return nil, fmt.Errorf("aucune correspondance trouvée pour la ville : %s", cityName)
	}

	return &locations[0], nil
}

func main() {
	var cityName string
	fmt.Print("Entrez le nom de la ville : ")
	fmt.Scanln(&cityName)

	apiKey := "c732a4f732342956ec521490b59a7dce"

	location, err := getCoordinates(cityName, apiKey)
	if err != nil {
		log.Fatalf("Erreur : %v", err)
	}

	fmt.Printf("Coordonnées de %s, %s :\n", location.Name, location.Country)
	fmt.Printf("Latitude : %f\n", location.Lat)
	fmt.Printf("Longitude : %f\n", location.Lon)
}
