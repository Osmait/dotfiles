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

func (s *CreateUseService) Save(ctx context.Context, email string, password string) error {

	user, err := domain.NewUser(email, password)
	if err != nil {
		return err
	}

	return s.userRepository.Save(ctx, user)

}
