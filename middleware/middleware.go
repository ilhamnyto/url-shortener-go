package middleware

import (
	"strings"

	"github.com/ilhamnyto/url-shortener-go/entity"
	"github.com/ilhamnyto/url-shortener-go/pkg/token"
	"github.com/labstack/echo/v4"
)

func ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			custErr := entity.UnauthorizedError("Token not found")
			return c.JSON(custErr.StatusCode, custErr)
		}
		
		tokenString := strings.Split(auth, "Bearer ")
		if len(tokenString) != 2  {
			custErr := entity.UnauthorizedError("Invalid token")
			return c.JSON(custErr.StatusCode, custErr)
		}
		
		token, err := token.ValidateToken(tokenString[1])
		
		if err != nil {
			custErr := entity.UnauthorizedError(err.Error())
			return c.JSON(custErr.StatusCode, custErr)
		}

		c.Set("user_id", token.UserID)
		
		return next(c)
	}
}