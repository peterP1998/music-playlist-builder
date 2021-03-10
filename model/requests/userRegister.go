package requests

type UserRegister struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Email string `json:"email"`
}