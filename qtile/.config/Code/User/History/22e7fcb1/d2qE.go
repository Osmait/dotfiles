package creating

import "github.com/osmait/blog-go/internal/domain"

type CreatePostService struct {
	postRepository domain.PostRepository
}

func NewCreatePostService(postRepository domain.PostRepository) *CreatePostService {
	return &CreatePostService{
		postRepository: postRepository,
	}
}
