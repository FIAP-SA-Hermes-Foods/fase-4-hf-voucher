package dto

type UserInput struct {
	User *User `json:"user,omitempty"`
}

type User struct {
	CPF          string `json:"cpf,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	WantRegister bool   `json:"wantRegister"`
}

type UserOutput struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Token      string `json:"token,omitempty"`
}
