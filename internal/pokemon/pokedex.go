package pokemon

import (
	"fmt"
	"strings"
)

func Pokedex() {
	fmt.Println("Your Pokedex:")
	myPokemonNames := make([]string, 0)

	for key := range myPokemon {
		myPokemonNames = append(myPokemonNames, fmt.Sprintf(" - %s", key))
	}

	fmt.Println(strings.Join(myPokemonNames, "\n"))
}
