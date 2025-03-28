package models

import "time"

type Incident struct {
	ID          int       `json:"id"`
	Reporter    string    `json:"reporter"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
