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

func (r *RestCountriesClient) GetTotalPopulationByRegion(region string, countries []model.Country) float64 {
	var totalPopulation float64

	for _, country := range countries {
		if country.Subregion == region {
			totalPopulation += float64(country.Population)
		}
	}

	return totalPopulation
}
