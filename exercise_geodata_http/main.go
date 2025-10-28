package main

import (
	"fmt"
	"test/exercise_geodata_http/client"
)

func main() {
	client := client.NewRestCountriesClient()

	countries, err := client.GetCountries()
	if err != nil {
		panic(err)
	}

	countriesFiltered, err := client.FilterByRegion("Western Europe", countries)
	if err != nil {
		panic(err)
	}

	regionTotalPopulation := client.GetTotalPopulation(countriesFiltered)

	fmt.Println(regionTotalPopulation)
}
