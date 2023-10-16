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

func NewPost(title string) (*Post, error) {

	idU, err := uuid.NewUUID()
	if err != nil {
		return &Post{}, err
	}
	return &Post{
		Id:        idU,
		Title:     title,
		CreatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (p *Post) ID() uuid.UUID {
	return p.Id

}
