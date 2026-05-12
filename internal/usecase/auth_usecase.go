package usecase

import (
	"errors"
	"godp/internal/domain"

	"godp/pkg/hash"
	jwtpkg "godp/pkg/jwt"
)

type AuthUsecase interface {
	Register(req *domain.User) (*domain.User, error)
	Login(email, password string) (string, *domain.User, error)
	Profile(userID uint) (*domain.User, error)
}

type authUsecase struct{ userRepo domain.UserRepository }

func NewAuthUsecase(userRepo domain.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

func (u *authUsecase) Register(req *domain.User) (*domain.User, error) {
	if req.Email == "" || req.Password == "" || req.Username == "" {
		return nil, errors.New("email, username, and password are required")
	}
	hashed, err := hash.Make(req.Password)
	if err != nil {
		return nil, err
	}
	req.Password = hashed
	if req.Role == "" {
		req.Role = "member"
	}
	if err := u.userRepo.Create(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (u *authUsecase) Login(email, password string) (string, *domain.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("email or password is wrong")
	}
	if !hash.Check(password, user.Password) {
		return "", nil, errors.New("email or password is wrong")
	}
	token, err := jwtpkg.Generate(user.ID, user.Role)
	return token, user, err
}

func (u *authUsecase) Profile(userID uint) (*domain.User, error) { return u.userRepo.FindByID(userID) }
