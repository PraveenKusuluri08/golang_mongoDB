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
