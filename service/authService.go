package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repository domain.AuthRepository
}

func NewDefaultAuthService(repository domain.AuthRepository) *DefaultAuthService {
	return &DefaultAuthService{repository: repository}
}

func (s *DefaultAuthService) Login(r dto.LoginRequest) (*dto.LoginResponse,  *errs.AppError) {
	login, err := s.repository.FindBy(r.Username, r.Password)
	if err != nil {
		return nil, err
	}

	token, err := login.ToDto()
	if err != nil {
		return nil, err
	}

	return token, nil
}
