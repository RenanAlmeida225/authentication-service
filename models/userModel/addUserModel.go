package userModel

type AddUserModel struct {
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Role     UserRole   `json:"role"`
	Status   UserStatus `json:"status"`
}
