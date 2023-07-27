package handler

import (
	"net/http"
	"strconv"
	"web-desa/helper"
	"web-desa/model"
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type wisataHandler struct {
	wisataService model.WisataService
}

func NewWisataHandler(wisataService model.WisataService) model.WisataHandler {
	return &wisataHandler{wisataService: wisataService}
}

func (u *wisataHandler) Mount(group *gin.RouterGroup) {
	group.POST("", u.StoreWisataHandler)        //create
	group.PATCH("/:id", u.EditWisataHandler)    //umodel
	group.GET("/:id", u.DetailWisataHandler)    //getById
	group.DELETE("/:id", u.DeleteWisataHandler) //delete
	group.GET("", u.FetchWisataHandler)         //getAll
}

func (u *wisataHandler) StoreWisataHandler(c *gin.Context) {
	var req request.WisataRequest

	err := c.ShouldBind(&req)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
		return 
	}

	// error tidak terdeteksi
	if err != nil {
		helper.ResponseValidatorErrorJson(c, err)
		return
	}

	link, err := u.wisataService.UploadImage(c)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, err)
		return
	}
	req.Gambar = link

	wisata, err := u.wisataService.StoreWisata(&req)
	if err !=nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", wisata)
}

func (u *wisataHandler) DetailWisataHandler(c *gin.Context) {	
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	wisata, err := u.wisataService.GetByID(uint(idUint))
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "", wisata)
}

func (u *wisataHandler) FetchWisataHandler(c *gin.Context) {	
	wisataList, err := u.wisataService.FetchWisata()
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", wisataList)
}

func (u *wisataHandler) EditWisataHandler(c *gin.Context) {	
	var req request.WisataRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
		return 
	}

	if err != nil {
		helper.ResponseValidatorErrorJson(c, err)
		return
	}

	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	wisata, err := u.wisataService.EditWisata(uint(idUint), &req)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", wisata)
}

func (h *wisataHandler) DeleteWisataHandler(c *gin.Context) {	
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	errDelImage := h.wisataService.DeleteImage(c, uint(idUint))
	if errDelImage != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, errDelImage)
		return
	}

	err := h.wisataService.DestroyWisata(uint(idUint))
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "delete success", "")
}