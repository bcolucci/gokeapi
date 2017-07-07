package endpoints

import (
	"encoding/json"
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

type ResourceList struct {
	Count    uint             `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Results  []ResourceResult `json:"results"`
}

func FetchResourceList(url string, mem *cache.Cache) *ResourceList {
	if list, found := mem.Get(url); found {
		return list.(*ResourceList)
	}
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	var list ResourceList
	if err = decoder.Decode(&list); err != nil {
		panic(err)
	}
	mem.Set(url, &list, cache.DefaultExpiration)
	return &list
}

func FetchNextResourceList(current *ResourceList, mem *cache.Cache) *ResourceList {
	if current.Next == "" {
		return nil
	}
	return FetchResourceList(current.Next, mem)
}

func FetchPreviousResourceList(current *ResourceList, mem *cache.Cache) *ResourceList {
	if current.Previous == "" {
		return nil
	}
	return FetchResourceList(current.Previous, mem)
}
