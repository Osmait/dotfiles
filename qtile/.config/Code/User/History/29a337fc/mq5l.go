package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type PostRepository interface {
	Save(ctx context.Context, post *Post) error
}

type Post struct {
	id        uuid.UUID
	title     string
	userId    string
	createdAt string
}

func NewPost(title string, userId string) (*Post, error) {

	idU, err := uuid.NewUUID()
	if err != nil {
		return &Post{}, err
	}
	return &Post{
		id:        idU,
		title:     title,
		createdAt: time.Now().Format(time.RFC3339),
		userId:    userId,
	}, nil
}

func (p *Post) ID() uuid.UUID {
	return p.id
}

func (p *Post) Title() string {
	return p.title

}
func (p *Post) UserID() string {
	return p.userId
}

func (p *Post) CreatedAt() string {
	return p.createdAt
}
