package api

type ResourceList struct {
	Count        int           `json:"count"`
	Next         string        `json:"next"`
	Previous     string        `json:"previous"`
	ResultsProxy []APIResource `json:"results"`
	Results      []interface{}
}
