package pokemon

import "fmt"

func Inspect(name string) {
	pokemon, ok := myPokemon[name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Println(pokemon.Detail())
}
