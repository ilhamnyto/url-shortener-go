package controller

import (
	"net/http"

	"github.com/ilhamnyto/url-shortener-go/entity"
	"github.com/ilhamnyto/url-shortener-go/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.InterfaceUserService
}

func NewUserController(service services.InterfaceUserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) CreateUser(c echo.Context) error {
	req := entity.CreateUserRequest{}

	if err := c.Bind(&req); err != nil {
		custErr := entity.GeneralError(err.Error())
		return c.JSON(custErr.StatusCode, custErr)
	}
	
	if custErr := u.service.CreateUser(&req); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusCreated, Message: "User Created Successfully!"}
	
	return c.JSON(succ.StatusCode, succ)
}

func (u *UserController) Login(c echo.Context) error {
	req := entity.UserLoginRequest{}

	if err := c.Bind(&req); err != nil {
		custErr := entity.GeneralError(err.Error())
		return c.JSON(custErr.StatusCode, custErr)
	}
	
	token, custErr := u.service.Login(&req)
	
	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusOK, Message: "Login Success!", Payload: token }

	return c.JSON(succ.StatusCode, succ)
}

func (u *UserController) UpdatePassword(c echo.Context) error {
	req := entity.UpdatePasswordRequest{}
	userId := c.Get("user_id").(int)

	if err := c.Bind(&req); err != nil {
		custErr := entity.GeneralError(err.Error())
		return c.JSON(custErr.StatusCode, custErr)
	}
	
	if custErr := u.service.UpdateUserPassword(userId, &req); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusOK, Message: "Password updated successfully!" }

	return c.JSON(succ.StatusCode, succ)
}

func (u *UserController) UserProfile(c echo.Context) error {

	userId := c.Get("user_id").(int)

	user, custErr := u.service.UserProfile(userId)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusOK, Message: "Success", Payload: user }

	return c.JSON(succ.StatusCode, succ)
}

func (u *UserController) UserProfileByUsername(c echo.Context) error {
	username := c.Param("username")

	user, custErr := u.service.UserProfileByUsername(username)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusOK, Message: "Success", Payload: user}

	return c.JSON(succ.StatusCode, succ)
}

