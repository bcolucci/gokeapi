package models

type Characteristic struct {
	ID             int           `json:"id"`
	GeneModulo     int           `json:"gene_modulo"`
	PossibleValues []int         `json:"possible_values"`
	Descriptions   []Description `json:"descriptions"`
}
