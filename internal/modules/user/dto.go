package user

// Requests

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
	Avatar   string `json:"avatar"` // NEW (optional)
}

type UpdateRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

type ToggleActiveRequest struct {
	IsActive bool `json:"is_active"`
}

type UpdateAvatarRequest struct {
	Avatar string `json:"avatar" binding:"required"` // NEW
}

// Responses (avoid leaking password)

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar,omitempty"`
	IsActive bool   `json:"is_active"`
}

func ToUserResponse(u User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Email:    u.Email,
		Role:     u.Role,
		Avatar:   u.Avatar,
		IsActive: u.IsActive,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Avatar   string `json:"avatar"`
}
