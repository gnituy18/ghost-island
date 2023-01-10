package user

import "time"

type User struct {
	Id string
	Name string
	Email string
	AvatarURL string

	CreatedAt time.Time
	UpdatedAt time.Time
}
