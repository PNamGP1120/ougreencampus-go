package user

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required,min=2"`
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
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required,min=2"`
	Role     string `json:"role" binding:"required"` // admin/organizer/student
	Avatar   string `json:"avatar"`
}

type UpdateProfileRequest struct {
	Name   string `json:"name" binding:"required,min=2"`
	Avatar string `json:"avatar"`
}

type ChangeRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

type ChangeStatusRequest struct {
	Status string `json:"status" binding:"required"` // active/blocked
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
