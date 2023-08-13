package handler

import (
	"net/http"
	"strconv"
	"web-desa/helper"
	"web-desa/model"
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type umkmHandler struct {
	umkmService model.UmkmService
}

func NewUmkmHandler(umkmService model.UmkmService) model.UmkmHandler {
	return &umkmHandler{umkmService: umkmService}
}

func (u *umkmHandler) Mount(group *gin.RouterGroup) {
	group.POST("", u.StoreUmkmHandler)        //create
	group.PATCH("/:id", u.EditUmkmHandler)    //umodel
	group.GET("/:id", u.DetailUmkmHandler)    //getById
	group.DELETE("/:id", u.DeleteUmkmHandler) //delete
	group.GET("", u.FetchUmkmHandler)         //getAll
}

func (u *umkmHandler) StoreUmkmHandler(c *gin.Context) {
	var req request.UmkmRequest

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

	link, err := u.umkmService.UploadImage(c)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, err)
		return
	}
	req.Gambar = link

	umkm, err := u.umkmService.StoreUmkm(&req)
	if err !=nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", umkm)
}

func (u *umkmHandler) DetailUmkmHandler(c *gin.Context) {	
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	umkm, err := u.umkmService.GetByID(uint(idUint))
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "", umkm)
}

func (u *umkmHandler) FetchUmkmHandler(c *gin.Context) {	
	umkmList, err := u.umkmService.FetchUmkm()
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", umkmList)
}

func (u *umkmHandler) EditUmkmHandler(c *gin.Context) {	
	var req request.UmkmRequest

	err := c.ShouldBind(&req)
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

	link, err := u.umkmService.UploadImage(c)
	if err != nil {
		umkm, err := u.umkmService.GetByID(uint(idUint))
		if err != nil {
			helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
			return 
		}
		req.Gambar = umkm.Gambar
	} else {
		errDelImage := u.umkmService.DeleteImage(c, uint(idUint))
		if errDelImage != nil {
			helper.ResponseErrorJson(c, http.StatusInternalServerError, errDelImage)
			return
		}
		req.Gambar = link
	}

	umkm, err := u.umkmService.EditUmkm(uint(idUint), &req)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", umkm)
}

func (h *umkmHandler) DeleteUmkmHandler(c *gin.Context) {	
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	errDelImage := h.umkmService.DeleteImage(c, uint(idUint))
	if errDelImage != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, errDelImage)
		return
	}

	err := h.umkmService.DestroyUmkm(uint(idUint))
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "delete success", "")
}