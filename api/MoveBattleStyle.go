package api

type MoveBattleStyle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []Name `json:"names"`
}
