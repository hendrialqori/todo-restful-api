package domain

import "time"

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	UserName  string    `json:"username"`
	Role      Role      `json:"role"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
}
