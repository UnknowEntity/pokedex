package pokemon

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/UnknowEntity/pokedex/config"
	"github.com/UnknowEntity/pokedex/internal"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

func Catch(name string) error {
	if name == config.STRING_FALSE {
		fmt.Println("Wrong format for command catch\nYou should use explore like this\ncatch <pokemon>")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	urlString, err := url.JoinPath(internal.API, POKEMON, name)

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

	var pokemon PokemonDetail

	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	isCatch, odd := catchResult(pokemon.BaseExperience)

	if !isCatch {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)

	myPokemon[name] = pokemon

	if odd < 1 {
		internal.UserExperience += int(odd * float32(pokemon.BaseExperience))
	}

	return nil
}

func catchResult(baseExp int) (bool, float32) {
	odd := float32(internal.UserExperience) / float32(baseExp)

	if odd > 1 {
		return true, odd
	}

	rollResult := random.Intn(100)

	return odd*100 > float32(rollResult), odd
}
