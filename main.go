package main

import (
	"gokeapi/api"
	"gokeapi/helpers"
)

func main() {
	client := api.NewAPI()
	resources := client.Berries.FetchResourceList()
	firstBerry := client.Berries.GetItem(resources, 0)
	helpers.PrintPrettyJSON(firstBerry)
	helpers.PrintPrettyJSON(client.Berries.GetItem(resources, 0))
	secondBerry := client.Berries.GetItemById(1)
	helpers.PrintPrettyJSON(secondBerry)
	helpers.PrintPrettyJSON(client.Berries.GetItemById(1))
	helpers.PrintPrettyJSON(client.Berries.GetItemByName("cheri"))
}
