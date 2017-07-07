package api

type MachineVersionDetail struct {
	MachineProxy APIResource `json:"machine"`
	Machine      Machine
}
