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

	totalPopulation := client.GetTotalPopulationByRegion("Western Europe", countries)

	fmt.Printf("the total population for western europe is: %f", totalPopulation)
}
