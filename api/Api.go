package api

import (
	"encoding/json"
	"gokeapi/helpers"
	"net/http"
)

type API struct {
	BaseURL string
	cache   *helpers.CacheWrapper
}

func NewAPI(baseURL string, cache *helpers.CacheWrapper) *API {
	return &API{baseURL, cache}
}

func (client *API) Populate(uri string, entity interface{}) {
	resourceURL := client.BaseURL + uri
	resource := client.cache.Get(resourceURL)
	if resource != nil {
		helpers.Unserialize(resource, entity)
		return
	}
	response, err := http.Get(resourceURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(entity); err != nil {
		panic(err)
	}
	bytes := helpers.Serialize(entity)
	client.cache.Set(resourceURL, bytes)
}
