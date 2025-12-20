package user

// ===== AUTH =====

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// ===== USER =====

type UpdateProfileRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateRoleRequest struct {
	Role string `json:"role"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}
