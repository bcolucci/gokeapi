package api

import (
	"pokeapi/berries"
	"time"

	"github.com/patrickmn/go-cache"
)

const apiBaseURL = "http://pokeapi.co/api"
const apiVersion = "v2"

type API struct {
	url     string
	mem     *cache.Cache
	Berries *berries.Endpoint
}

func NewAPI() *API {
	api := &API{}
	api.url = apiBaseURL + "/" + apiVersion
	api.mem = cache.New(60*time.Minute, 5*time.Minute)
	api.Berries = berries.NewEndpoint(api.url+"/berry", api.mem)
	return api
}
