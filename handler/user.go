package handler

import (
	"errors"
	"net/http"
	"strconv"
	"web-desa/config/middleware"
	"web-desa/helper"
	"web-desa/model"
	"web-desa/request"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userService model.UserService
}

type UserHandler interface {
	Mount(group *gin.RouterGroup)
}

func NewUserHandler(userService model.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) Mount(group *gin.RouterGroup)  {
	group.POST("/register", h.RegisterUserHandler)
	group.POST("/login", h.UserLogin)
    group.GET("", h.GetAllUserHandler)
    group.PATCH("/:id", h.UpdateUserHandler)		
    group.DELETE("/:id", h.DeleteUserHandler)
}

func HashPassword(password string) string{
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err)
		return 
	}
	
	if len(userRequest.Password) <= 6 {
		helper.ResponseValidationErrorJson(c, "Too short password", err)
		return 
	}

	registerUser, err := h.userService.Register(&userRequest)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", registerUser)
}

func (h *userHandler) UserLogin(c *gin.Context) {
	var req request.UserRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err)
		return 
	}

	user, err := h.userService.GetByUsername(req.Username)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, errors.New("invalid username or password"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, errors.New("invalid username or password"))
		return
	}

	tokenJwt, err := middleware.GenerateToken(user.ID)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, errors.New("error generating token"))
		return
	}

	helper.ResponseSuccessJson(c, "success", gin.H{
		"data": user,
		"token": tokenJwt,
	})
}

func (h *userHandler) GetAllUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", users)
}

func (h *userHandler) UpdateUserHandler(c *gin.Context)  {
	var UserRequest request.UserRequest

	err := c.ShouldBindJSON(&UserRequest)

	idString := c.Param("id")
	id, _ := strconv.ParseUint(idString, 10, 32)

	update, err := h.userService.EditUser(uint(id), &UserRequest)
	if err != nil {
		helper.ResponseDetailErrorJson(c, "Error Update User",err)
		return 
	}

	helper.ResponseSuccessJson(c, "success", update)
}

func (h *userHandler) DeleteUserHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.ParseUint(idString, 10, 32)

	err := h.userService.DestroyUser(uint(id))
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "delete success", "")
}