package user

type CreateNewUserViewModel struct {
	LastName      string `validate:"required,min=3,max=30"`
	FirstName     string `validate:"required,min=3,max=30"`
	Email         string `validate:"required,email"`
	BirthDate     string `validate:"required,datetime=2006/01/02"`
	BirthLocation string `validate:"required"`
	Age           int    `validate:"required"`
	PhoneNumber   string `validate:"required,startswith=09,lt=12"`
	UserName      string
	Password      string
	CreatorUserId string
	AvatarName    string
}

type EditUserViewModel struct {
	TargetUserId string
	LastName     string `validate:"required"`
	FirstName    string `validate:"required"`
	Email        string `validate:"required"`
	UserName     string `validate:"required"`
	Password     string `validate:"required"`
}
type EditUserRoleViewModel struct {
	TargetUserId string
	Roles        []string `validate:"required"`
}
type EditUserPasswordViewModel struct {
	TargetUserId string
	Password     string `validate:"required"`
}
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}
