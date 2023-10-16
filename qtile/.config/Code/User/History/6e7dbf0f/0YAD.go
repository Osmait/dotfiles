package postgres

import (
	"database/sql"

	"github.com/osmait/blog-go/internal/domain"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Save(post *domain.Post) error {
	_, err := r.db.Exec("INSERT INTO posts (id,title, created_at,user_id) VALUES ($1, $2,$3,$4)", post.ID(), post.Title(), post.CreatedAt())
	return err

}
