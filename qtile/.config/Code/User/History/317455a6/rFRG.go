package creating

import (
	"context"

	"github.com/osmait/blog-go/internal/domain"
	"github.com/osmait/blog-go/internal/platfrom/storage/postgres"
)

type CreateUseService struct {
	userRepository domain.UserRepositoy
}

func NewCreateUseService(userRepository postgres.UserRepository) *CreateUseService {
	return &CreateUseService{
		userRepository: userRepository,
	}

}

func (s *CreateUseService) Save(ctx context.Context, user *domain.User) error {

	return s.userRepository.Save(ctx, user)

}
