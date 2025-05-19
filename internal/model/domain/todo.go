package domain

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Title     int       `json:"title"`
	Status    string    `json:"status"`
	User      User      `json:"user"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
