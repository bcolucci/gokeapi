package main

import (
	"gokeapi/api"
	"gokeapi/helpers"
	"gokeapi/models"
)

func main() {
	cache := helpers.NewCache("./cache", 60*24)
	client := api.NewAPI("http://pokeapi.co/api/v2", cache)
	jp := &models.Language{}
	client.Populate("/language/1/", jp)
	helpers.PrintPrettyJSON(jp)
}
