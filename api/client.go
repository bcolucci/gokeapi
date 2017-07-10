package api

import (
	"encoding/json"
	"gokeapi/helpers"
	"net/http"
)

const BaseURL = "http://pokeapi.co/api/v2"

type Client struct {
	BaseURL string
	cache   *helpers.CacheWrapper
}

func NewClient() *Client {
	cache := helpers.NewCache("./cache", 60*24*7)
	return NewCustomClient(BaseURL, cache)
}

func NewCustomClient(baseURL string, cache *helpers.CacheWrapper) *Client {
	return &Client{baseURL, cache}
}

func (client *Client) Populate(uri string, entity interface{}) {
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
