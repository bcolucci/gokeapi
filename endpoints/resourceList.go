package endpoints

type ResourceList struct {
	Count    uint             `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Results  []ResourceResult `json:"results"`
}
