package main

import (
	"gokeapi/api"
	"gokeapi/helpers"
)

func main() {
	cache := api.NewCache("./cache", 60)
	resourceRetriever := api.NewResourceRetriever(cache)
	client := api.NewAPI("http://pokeapi.co/api/v2", resourceRetriever)
	helpers.PrintPrettyJSON(client.Get("/type/3/"))
}
