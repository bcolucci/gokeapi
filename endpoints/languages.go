package endpoints

import (
	"gokeapi/api"
	"gokeapi/models"
	"strconv"
)

type LanguagesEndpoint struct {
	BaseURI string
	client  *api.Client
}

func NewLanguagesEndpoint(client *api.Client) *LanguagesEndpoint {
	return &LanguagesEndpoint{"/language", client}
}

func (e *LanguagesEndpoint) GetById(id int) *models.Language {
	lang := &models.Language{}
	e.client.Populate(e.BaseURI+"/"+strconv.Itoa(id), lang)
	return lang
}
