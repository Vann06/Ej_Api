package models

type Incident struct {
	ID          int    `json:"id"`
	Reporter    string `json:"reporter"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}
