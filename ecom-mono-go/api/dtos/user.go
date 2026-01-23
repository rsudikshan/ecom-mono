package dtos

type UserRegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email	 string `json:"email"`
}

type UserLoginParams struct {
	Email 	 string `json:"email"`
	Password string `json:"password"`
}