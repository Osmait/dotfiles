package find

import (
	"context"

	"github.com/osmait/blog-go/internal/domain"
)

type Find struct {
	userRepository domain.UserRepositoy
}

func NewFind(userRepository domain.UserRepositoy) *Find {
	return &Find{
		userRepository: userRepository,
	}
}

func (f *Find) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return f.userRepository.FindByEmail(ctx, email)

}
