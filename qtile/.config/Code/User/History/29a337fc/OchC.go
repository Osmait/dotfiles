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
	return &Post{
		Id:        uuid.New(),
		Title:     title,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
