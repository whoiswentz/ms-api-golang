package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	Verify(string, string, map[string]string) (bool, *errs.AppError)
}

type DefaultAuthService struct {
	repository domain.AuthRepository
	rolePermissions domain.RolePermissions
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

func (s *DefaultAuthService) Verify(token string, routeName string, vars map[string]string) (bool, *errs.AppError) {
	jwtToken, err := jwtTokenFromString(token)
	if err != nil {
		return false, errs.UnexpectedError(err.Error())
	}

	if jwtToken.Valid {
		mapClaims := jwtToken.Claims.(jwt.MapClaims)
		claims, err := domain.BuildClaimsFromJwtMapClaims(mapClaims)
		if err != nil {
			return false, errs.UnexpectedError(err.Error())
		}

		if claims.IsUserRole() && !claims.IsRequestVerifiedWithTokenClaims(vars) {
			return false, nil
		}
		return s.rolePermissions.IsAuthorizedFor(claims.Role, routeName), nil
	}
	return false, errs.UnexpectedError("invalid token")
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}
