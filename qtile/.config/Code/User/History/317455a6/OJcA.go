package creating

import "github.com/osmait/blog-go/internal/domain"

type CreateUseService struct {
	userRepository domain.UserRepositoy
}

func NewCreateUseService(userRepository domain.UserRepositoy) *CreateUseService {
	return &CreateUseService{
		userRepository: userRepository,
	}

}
