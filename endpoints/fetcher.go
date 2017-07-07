package endpoints

import (
	"encoding/json"
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

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

//TODO use cache mechanism
func FetchResourceItem(url string, mem *cache.Cache, v interface{}) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
  decoder := json.NewDecoder(response.Body)
  if err = decoder.Decode(v); err != nil {
		panic(err)
	}
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
