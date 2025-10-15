package authModel

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseAuthModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
