package app

type UserRegister struct {
	Username string `form:"username" valid:"required"`
	Email    string `form:"email" valid:"required,email"`
	Password string `form:"password" valid:"required,minstringlength(6)"`
}

type UserLogin struct {
	Email    string `form:"email" valid:"required,email"`
	Password string `form:"password" valid:"required"`
}

type UserUpdate struct {
	Username string `form:"username" valid:"required"`
	Email    string `form:"email" valid:"required,email"`
	Password string `form:"password" valid:"required,minstringlength(6)"`
}
