package user

/* ===== AUTH ===== */

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

/* ===== USER ===== */

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Role     Role   `json:"role"`
}

type UpdateUserRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ChangeRoleRequest struct {
	Role Role `json:"role" binding:"required"`
}

type ChangeStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
