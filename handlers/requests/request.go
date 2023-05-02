package requests

type CoachRequest struct {
	Name       string  `json:"name"`
	Percentage float32 `json:"percentage"`
}

type TeamRequest struct {
	Name string `json:"name"`
	Logo string `json:"logo" validate:"required"`
}

type PlayerRequest struct {
	Name     string `json:"name"`
	Number   int8   `json:"number"`
	Country  string `json:"country"`
	Position string `json:"position"`
	Role     string `json:"role"`
}
