package creating

import (
	"context"

	"github.com/osmait/blog-go/internal/domain"
)

type CreateUseService struct {
	userRepository domain.UserRepositoy
}

func NewCreateUseService(userRepository domain.UserRepositoy) *CreateUseService {
	return &CreateUseService{
		userRepository: userRepository,
	}

}

func (s *CreateUseService) Save(ctx context.Context, user *domain.User) error {

	return s.userRepository.Save(ctx, user)

}
