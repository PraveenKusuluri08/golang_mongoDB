package model

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"tokenString"`
}

func (t Token) Valid() error {
	//TODO implement me
	panic("Failed")
}
