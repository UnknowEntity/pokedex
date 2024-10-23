package location

import (
	"encoding/json"

	poke "github.com/UnknowEntity/pokedex/internal"
)

func getLocationList(urlString string) ([]string, error) {
	body, _, err := poke.GetUrl(urlString)

	if err != nil {
		return nil, err
	}

	return byteToList(body)
}

func byteToList(data []byte) ([]string, error) {
	var listLocation poke.List[Location]

	if err := json.Unmarshal(data, &listLocation); err != nil {
		return nil, err
	}

	nextUrl = listLocation.Next
	prevUrl = listLocation.Previous

	locations := make([]string, 0)

	for _, location := range listLocation.Results {
		locations = append(locations, location.Name)
	}

	return locations, nil
}
