package main

import (
	"pokeapi/api"
	"pokeapi/helpers"
)

func main() {
	client := api.NewAPI()
	resources := client.Berries.FetchResourceList()
	firstBerry := client.Berries.GetItem(resources, 0)
	helpers.PrintPrettyJSON(firstBerry)
}
