package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func AssignService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	userData := User{}
	userData.Name = input.Name
	userData.Email = input.Email
	userData.Occupation = input.Occupation

	passw, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return userData, err
	}

	userData.PasswordHash = string(passw)
	userData.Role = "user"

	user, err := s.repository.Save(userData)

	if err != nil {
		return user, err
	}

	return user, err
}
