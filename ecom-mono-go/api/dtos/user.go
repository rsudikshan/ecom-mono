package dtos

type UserSignUpParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email	 string `json:"email" binding:"required"`
}

type UserLoginParams struct {
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct{
	Token string `json:"auth_token"`
}