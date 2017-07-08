package api

import (
	"io/ioutil"
	"net/http"
)

type ResourceRetriever struct {
	cache *Cache
}

func NewResourceRetriever(cache *Cache) *ResourceRetriever {
	return &ResourceRetriever{cache}
}

func (r *ResourceRetriever) Get(url string) []byte {
	resource := r.cache.Get(url)
	if resource != nil {
		return resource
	}
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	r.cache.Set(url, bytes)
	return bytes
}
