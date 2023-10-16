package user

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
