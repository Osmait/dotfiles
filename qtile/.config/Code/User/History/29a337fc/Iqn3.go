package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	id        uuid.UUID
	title     string
	createdAt string
}

func NewPost(title string) (*Post, error) {

	idU, err := uuid.NewUUID()
	if err != nil {
		return &Post{}, err
	}
	return &Post{
		id:        idU,
		title:     title,
		createdAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (p *Post) ID() uuid.UUID {
	return p.id
}

func (p *Post) Title() string {
	return p.title

}

func (p *Post) CreatedAt() string {
	return p.createdAt
}
