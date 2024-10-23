package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/UnknowEntity/pokedex/config"
	"github.com/UnknowEntity/pokedex/internal/location"
	"github.com/UnknowEntity/pokedex/internal/pokemon"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		input, err := scanner.ReadString('\n')

		if err != nil {
			log.Fatalln(err)
		}

		input = strings.Trim(input, "\n\r ")

		if isEnd := handleInput(input); isEnd {
			break
		}
	}

	fmt.Println("Thanks for using pokedex!\nGoodbye")
}

const helpScreen = `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`

func handleInput(input string) bool {
	params := strings.Split(input, config.SPACE)

	if len(params) == 0 {
		return false
	}

	command, parameters := params[0], params[1:]

	switch command {
	case "exit":
		return true

	case "help":
		fmt.Println(helpScreen)

	case "map":
		if err := location.Map(); err != nil {
			log.Fatalln(err)
		}

	case "mapb":

		if err := location.Mapb(); err != nil {
			log.Fatalln(err)
		}

	case "explore":

		if err := location.Explore(strings.Join(parameters, config.STRING_FALSE)); err != nil {
			log.Fatalln(err)
		}

	case "catch":

		if err := pokemon.Catch(strings.Join(parameters, config.STRING_FALSE)); err != nil {
			log.Fatalln(err)
		}

	case "inspect":

		pokemon.Inspect(strings.Join(parameters, config.STRING_FALSE))

	case "pokedex":

		pokemon.Pokedex()

	}

	return false
}
