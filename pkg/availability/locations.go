package availability

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Locations represents all the defined vaccine locations.
var Locations map[int]Location

//go:embed locations.json
// locationData is the byte representation of the predefined location list, embedded from locations.json.
//
// To update, run the following code from the "Select a location" page:
//  console.log(JSON.stringify(JSON.parse(window.__INITIAL_STATE__).locations))
var locationData []byte

func init() {
	Locations = map[int]Location{}

	var locationList []Location
	err := json.Unmarshal(locationData, &locationList)
	if err != nil {
		panic(fmt.Errorf("error loading location data: %w", err))
	}

	for _, location := range locationList {
		Locations[location.LocId] = location
	}
}
