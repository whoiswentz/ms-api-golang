package domain

import (
	"banking/dto"
	"banking/errs"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
	"time"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username string
	CustomerId sql.NullString
	Accounts sql.NullString
	Role string
}

func (l Login) GenerateToken() (*string, *errs.AppError) {
	var claims jwt.MapClaims
	if l.Accounts.Valid && l.CustomerId.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		log.Println("Failed while signing token: " + err.Error())
		return nil, errs.UnexpectedError("cannot generate token")
	}
	return &signedTokenAsString, nil
}

func (l Login) ToDto() (*dto.LoginResponse, *errs.AppError) {
	token, err := l.GenerateToken()
	return &dto.LoginResponse{Token: *token}, err
}

func (l Login) claimsForUser() jwt.MapClaims {
	accounts := strings.Split(l.Accounts.String, ",")
	return jwt.MapClaims{
		"customer_id": l.CustomerId.String,
		"role":        l.Role,
		"username":    l.Username,
		"accounts":    accounts,
		"exp":         time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
func (l Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"role":     l.Role,
		"username": l.Username,
		"exp":      time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
