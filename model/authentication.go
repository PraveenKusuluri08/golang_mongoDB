package model

type AuthSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPassword struct {
	Password string `json:"password"`
}

type Error struct {
	Error string
}
type ChangePassword struct {
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}
