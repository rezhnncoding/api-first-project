package usermodel

type LoginViewModel struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
}
