package dto

//Binding like required, min , max, validation input from clients should happens here i DTO
//*its my opinion, it can be wrong

//UserRegisterDTO is a structure that clients need to fill register
//It doesnt represents users table in database
type UserRegisterDTO struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

//UserLoginDTO is a structure that clients need to fill when login
type UserLoginDTO struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

//UserUpdateDTO is a structure when user needs to update their own profile
type UserUpdateDTO struct {
	ID       uint64 `form:"id" json:"id"`
	Name     string `form:"name" json:"name,omitempty"`
	Email    string `form:"email" json:"email,omitempty"`
	Password string `form:"password" json:"password,omitempty"`
}
