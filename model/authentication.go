package model

type AuthSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPassword struct {
	Password string `json:"password"`
}

type AuthSignUp struct {
	UserName           string `json:"userName"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	IsExists           string `json:"isExists"`
	Role               int    `json:"role"`
	AccountCreatedDate string `json:"createdat"`
}

type Error struct {
	Error string
}
