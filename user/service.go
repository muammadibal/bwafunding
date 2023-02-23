package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	CheckAvailabilityEmail(input CheckEmailInput) (bool, error)
	UploadAvatar(ID int, fileLocation string) (User, error)
	UserDetail(ID int) (User, error)
}

type service struct {
	repository Repository
}

func AssignService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
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

func (s *service) Login(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User doesn't exist")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckAvailabilityEmail(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) UploadAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation
	userData, err := s.repository.Update(user)
	if err != nil {
		return userData, err
	}

	return userData, nil
}

func (s *service) UserDetail(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}
