package location

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/UnknowEntity/pokedex/config"
	poke "github.com/UnknowEntity/pokedex/internal"
)

func Map() error {
	urlString, err := url.JoinPath(poke.API, LOCATION)

	if err != nil {
		return err
	}

	if nextUrl != config.STRING_FALSE {
		urlString = nextUrl
	}

	locations, err := getLocationList(urlString)

	if err != nil {
		return err
	}

	fmt.Println(strings.Join(locations, "\n"))

	return nil
}
