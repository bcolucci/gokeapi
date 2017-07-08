package api

type API struct {
	BaseURL           string
	resourceRetriever *ResourceRetriever
}

func NewAPI(baseURL string, resourceRetriever *ResourceRetriever) *API {
	return &API{baseURL, resourceRetriever}
}

func (client *API) Get(uri string) []byte {
	return client.resourceRetriever.Get(client.BaseURL + uri)
}
