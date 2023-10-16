package find

import (
	"context"

	"github.com/osmait/blog-go/internal/domain"
)

type FindUserService struct {
	userRepository domain.UserRepositoy
}

func NewFind(userRepository domain.UserRepositoy) *FindUserService {
	return &FindUserService{
		userRepository: userRepository,
	}
}

func (f *FindUserService) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return f.userRepository.FindByEmail(ctx, email)

}
