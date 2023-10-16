package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id        uuid.UUID
	Title     string
	CreatedAt string
}

func NewPost(title string) *Post {

	idU, err := uuid.NewUUID()
	if err != nil {
		return User{}, err
	}
	return &Post{
		Id:        idU,
		Title:     title,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
