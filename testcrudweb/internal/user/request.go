package user

// CreateUserRequest representa os dados necessários para criar um usuário
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}
