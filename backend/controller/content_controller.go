package controller

import (
	"go-restapi/model"
	"go-restapi/usecase"
	"strconv"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IContentController interface {
	GetAllContents(c echo.Context) error
	GetContentById(c echo.Context) error
	CreateContent(c echo.Context) error
	UpdateContent(c echo.Context) error
	DeleteContent(c echo.Context) error
}

type contentController struct {
	cu usecase.IContentUsecase
}

func NewContentController(cu usecase.IContentUsecase) IContentController {
	return &contentController{cu}
}

func (cc *contentController) GetAllContents(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	contentsRes, err := cc.cu.GetAllContents(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, contentsRes)
}

func (cc *contentController) GetContentById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("contentId")
	contentId, _ := strconv.Atoi(id)
	contentRes, err := cc.cu.GetContentById(uint(userId.(float64)), uint(contentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, contentRes)
}

func (cc *contentController) CreateContent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	content := model.Content{}
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	content.UserId = uint(userId.(float64))
	contentRes, err := cc.cu.CreateContent(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, contentRes)
}

func (cc *contentController) UpdateContent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("contentId")
	contentId, _ := strconv.Atoi(id)

	content := model.Content{}
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	contentRes, err := cc.cu.UpdateContent(content, uint(userId.(float64)), uint(contentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, contentRes)
}

func (cc *contentController) DeleteContent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("contentId")
	contentId, _ := strconv.Atoi(id)

	if err := cc.cu.DeleteContent(uint(userId.(float64)), uint(contentId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
