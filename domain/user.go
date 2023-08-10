package domain

import "time"

type User struct {
	ID        int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
