package user

type UserService struct {
	repository Repository
}

func NewUserService(repository Repository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) CreateUser(createUserRequest CreateUserRequest) (User, error) {
	newUser := NewUser(createUserRequest.Name)
	return service.repository.Create(newUser)
}

func (service *UserService) FindUserById(id string) (User, error) {
	return service.repository.FindById(id)
}
