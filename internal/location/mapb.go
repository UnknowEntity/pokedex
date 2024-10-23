package location

import (
	"fmt"
	"strings"

	"github.com/UnknowEntity/pokedex/config"
)

func Mapb() error {
	var urlString string

	if prevUrl == config.STRING_FALSE {
		fmt.Println("first page\ncannot go back")
		return nil
	}

	urlString = prevUrl

	locations, err := getLocationList(urlString)

	if err != nil {
		return err
	}

	fmt.Println(strings.Join(locations, "\n"))

	return nil
}
