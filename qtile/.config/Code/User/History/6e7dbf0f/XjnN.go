package postgres

import (
	"context"
	"database/sql"

	"github.com/osmait/blog-go/internal/domain"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Save(ctx context.Context, post *domain.Post) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO posts (id,title, create_at,users_id) VALUES ($1, $2,$3,$4)", post.ID(), post.Title(), post.CreatedAt(), post.UserID())
	return err

}
