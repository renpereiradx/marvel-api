package model

import "time"

type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
	Changed_At time.Time `json:"changed_at"`
}
