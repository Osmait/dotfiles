package find

import "github.com/osmait/blog-go/internal/domain"

type Find struct {
	userRepository domain.UserRepositoy
}

func NewFind(userRepository domain.UserRepositoy) *Find {
	return &Find{
		userRepository: userRepository,
	}
}
