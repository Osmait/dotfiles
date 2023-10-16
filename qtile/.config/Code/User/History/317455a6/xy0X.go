package creating

import (
	"context"

	"github.com/osmait/blog-go/internal/domain"
	"github.com/osmait/blog-go/internal/platfrom/storage/postgres"
)

type CreateUseService struct {
	userRepository postgres.UserRepository
}

func NewCreateUseService(userRepository postgres.UserRepository) *CreateUseService {
	return &CreateUseService{
		userRepository: userRepository,
	}

}

func (s *CreateUseService) Save(ctx context.Context, email string, password string) error {

	user, err := domain.NewUser(email, password)
	if err != nil {
		return err
	}

	return s.userRepository.Save(ctx, &user)

}
