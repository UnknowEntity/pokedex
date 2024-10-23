package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/UnknowEntity/pokedex/config"
	"github.com/UnknowEntity/pokedex/internal"
	"github.com/UnknowEntity/pokedex/internal/pokemon"
)

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        pokemon.Pokemon `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func Explore(paramString string) error {
	area := strings.TrimSpace(paramString)

	if area == config.STRING_FALSE {
		fmt.Println("Wrong format for command explore\nYou should use explore like this\nexplore <area>")
		return nil
	}

	fmt.Printf("Exploring %s...\n", area)

	urlString, err := url.JoinPath(internal.API, LOCATION, area)

	if err != nil {
		return err
	}

	body, statusCode, err := internal.GetUrl(urlString)

	if err != nil {
		return err
	}

	if statusCode == http.StatusNotFound {
		fmt.Println("Not found.")
		return nil
	}

	var locationArea LocationArea

	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}

	pokemons := make([]string, 0)

	for _, pokemon := range locationArea.PokemonEncounters {
		pokemons = append(pokemons, fmt.Sprintf(" - %s", pokemon.Pokemon.Name))
	}

	fmt.Println("Found Pokemon:")
	fmt.Println(strings.Join(pokemons, "\n"))

	return nil
}
