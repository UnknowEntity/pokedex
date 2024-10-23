package internal

import (
	"time"

	"github.com/UnknowEntity/pokedex/internal/cache"
)

const API = "https://pokeapi.co/api/v2"

var URLCache = cache.NewCache(time.Second * 5)

var UserExperience = 50
