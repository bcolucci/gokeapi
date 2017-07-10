package main

import (
	"gokeapi/api"
	"gokeapi/endpoints"
	"gokeapi/helpers"
)

func main() {
	client := api.NewClient()
	languages := endpoints.NewLanguagesEndpoint(client)
	jp := languages.GetById(1)
	helpers.PrintPrettyJSON(jp)
}
