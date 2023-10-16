package creating

import (
	"context"

	"github.com/osmait/blog-go/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type CreateUseService struct {
	userRepository domain.UserRepositoy
}

func NewCreateUseService(userRepository domain.UserRepositoy) *CreateUseService {
	return &CreateUseService{
		userRepository: userRepository,
	}

}

func (s *CreateUseService) Save(ctx context.Context, email, password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := domain.NewUser(email, string(hashPassword))

	if err != nil {
		return err
	}

	return s.userRepository.Save(ctx, user)

}
