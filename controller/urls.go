package controller

import (
	"net/http"

	"github.com/ilhamnyto/url-shortener-go/entity"
	"github.com/ilhamnyto/url-shortener-go/services"
	"github.com/labstack/echo/v4"
)

type UrlController struct {
	service services.InterfaceUrlService
}

func NewUrlController(service services.InterfaceUrlService) *UrlController {
	return &UrlController{service: service}
}

func (u *UrlController) CreateUrl(c echo.Context) error {
	req := entity.CreateUrlRequest{}

	if err := c.Bind(&req); err != nil {
		custErr := entity.GeneralError(err.Error())
		return c.JSON(custErr.StatusCode, custErr)
	}
	
	req.UserID = c.Get("user_id").(int)

	if custErr := u.service.CreateUrl(&req); custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusCreated, Message: "Url created successfully"}

	return c.JSON(succ.StatusCode, succ)
}

func (u *UrlController) Redirect(c echo.Context) error {
	shortUrl := c.Param("short_url")

	longUrl, custErr := u.service.GetUrlByShortUrl(shortUrl)

	if custErr != nil {
		c.JSON(custErr.StatusCode, custErr)
	}

	return c.Redirect(http.StatusMovedPermanently, longUrl)
}

func (u *UrlController) GetUserUrls(c echo.Context) error {
	userId := c.Get("user_id").(int)

	urls, custErr := u.service.GetUrlsByUserId(userId)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusOK, Message: "SUCCESS", Payload: urls}

	return c.JSON(succ.StatusCode, succ)
}

func (u *UrlController) GetUserUrlsByUsername(c echo.Context) error {
	username := c.Param("username")

	urls, custErr := u.service.GetUrlsByUsername(username)

	if custErr != nil {
		return c.JSON(custErr.StatusCode, custErr)
	}

	succ := entity.Response{StatusCode: http.StatusOK, Message: "SUCCESS", Payload: urls}

	return c.JSON(succ.StatusCode, succ)
}
