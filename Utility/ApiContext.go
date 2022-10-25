package Utility

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"puppy/ViewModel/common/security"
)

type ApiContext struct {
	echo.Context
}

func (c ApiContext) GetUserId() (userId string, err error) {
	defer func() {
		if r := recover(); r != nil {
			userId = ""
			err = errors.New("user is not login")
		}
	}()
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*security.JwtClaims)
	return claim.UserId, nil
}
