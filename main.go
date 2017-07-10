package main

import (
	"gokeapi/api"
	"gokeapi/helpers"
)

func main() {
	client := api.NewClient()
	jp := client.Languages.GetById(1)
	helpers.PrintPrettyJSON(jp)
}
