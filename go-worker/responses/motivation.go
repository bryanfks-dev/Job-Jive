package responses

type MotivationResponse struct {
	UserId     int    `json:"user_id"`
	Motivation string `json:"motivation"`
}
