package domain

import (
	"time"

	"github.com/google/uuid"
)

type PostRepository interface {
	Save(post *Post) error
}

type Post struct {
	id        uuid.UUID `json:"id" db:"id`
	title     string    `json:"title" db:"title"`
	userId    uuid.UUID `json:"userId" db:"user_id"`
	createdAt string    `json:"createdAt" db:"created_at"`
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
func (p *Post) UserID() uuid.UUID {
	return p.userId
}

func (p *Post) CreatedAt() string {
	return p.createdAt
}
