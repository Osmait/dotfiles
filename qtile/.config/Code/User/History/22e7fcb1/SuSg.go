package creating

import (
	"context"

	"github.com/osmait/blog-go/internal/domain"
)

type CreatePostService struct {
	postRepository domain.PostRepository
}

func NewCreatePostService(postRepository domain.PostRepository) *CreatePostService {
	return &CreatePostService{
		postRepository: postRepository,
	}
}

func (c *CreatePostService) Save(ctx context.Context, title string, userId string) error {

	post, err := domain.NewPost(title, userId)

	if err != nil {
		return err

	}
	return c.postRepository.Save(ctx, post)

}
