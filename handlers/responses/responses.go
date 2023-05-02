package responses

type TeamResponse struct {
	Id      int16            `json:"id"`
	Name    string           `json:"name"`
	Logo    string           `json:"logo"`
	Coach   *CoachResponse   `json:"coach"`
	Players []PlayerResponse `json:"players"`
}

type PlayerResponse struct {
	Id       int16  `json:"id"`
	Name     string `json:"name"`
	Number   int8   `json:"number"`
	Country  string `json:"country"`
	Position string `json:"position"`
	Role     string `json:"role"`
}

type CoachResponse struct {
	Id         int16   `json:"id"`
	Name       string  `json:"name"`
	Percentage float32 `json:"percentage"`
}
