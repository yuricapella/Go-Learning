package user

type Repository interface {
	Create(user User) (User, error)
	FindById(id string) (User, error)
}
