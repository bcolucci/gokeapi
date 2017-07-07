package berries

import (
	"errors"
	"pokeapi/endpoints"

	cache "github.com/patrickmn/go-cache"
)

type Endpoint struct {
	url string
	mem *cache.Cache
}

func NewEndpoint(url string, mem *cache.Cache) *Endpoint {
	return &Endpoint{url, mem}
}

func (e *Endpoint) FetchResourceList() *endpoints.ResourceList {
	return endpoints.FetchResourceList(e.url, e.mem)
}

func (e *Endpoint) FetchNextResourceList(current *endpoints.ResourceList) *endpoints.ResourceList {
	return endpoints.FetchNextResourceList(current, e.mem)
}

func (e *Endpoint) FetchPreviousResourceList(current *endpoints.ResourceList) *endpoints.ResourceList {
	return endpoints.FetchPreviousResourceList(current, e.mem)
}

func (e *Endpoint) GetItem(current *endpoints.ResourceList, idx int) *Berry {
	if idx > len(current.Results) {
		panic(errors.New("Out of range"))
	}
	return NewBerry(current.Results[idx].URL, e.mem)
}

func (e *Endpoint) GetItems(current *endpoints.ResourceList) []Berry {
	var items []Berry
	nbItems := len(current.Results)
	items = make([]Berry, nbItems, nbItems)
	for idx, result := range current.Results {
		items[idx] = *NewBerry(result.URL, e.mem)
	}
	return items
}
