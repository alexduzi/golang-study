package client

import (
	"encoding/json"
	"io"
	"net/http"
	"test/exercise_geodata_http/model"
)

const URL_REST_COUNTRIES = "https://restcountries.com/v3.1/region/europe"

type RestCountriesClient struct {
}

func NewRestCountriesClient() *RestCountriesClient {
	return &RestCountriesClient{}
}

func (r *RestCountriesClient) GetCountries() ([]model.Country, error) {
	req, err := http.NewRequest("GET", URL_REST_COUNTRIES, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var countries []model.Country
	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, err
	}

	return countries, err
}

func (r *RestCountriesClient) FilterByRegion(region string, countries []model.Country) ([]model.Country, error) {
	countriesByRegion := []model.Country{}

	for _, country := range countries {
		if country.Subregion == region {
			countriesByRegion = append(countriesByRegion, country)
		}
	}

	return countriesByRegion, nil
}

func (r *RestCountriesClient) GetTotalPopulation(countries []model.Country) map[string]float64 {
	totalPopulation := make(map[string]float64)
	for _, country := range countries {
		totalPopulation[country.Subregion] += float64(country.Population)
	}
	return totalPopulation
}
