package handler

import (
	"net/http"
	"strconv"
	"web-desa/config/middleware"
	"web-desa/helper"
	"web-desa/model"
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type infoKegiatanHandler struct {
	infoKegiatanService model.InfoKegiatanService
}

func NewInfoKegiatanHandler(infoKegiatanService model.InfoKegiatanService) model.InfoKegiatanHandler {
	return &infoKegiatanHandler{infoKegiatanService: infoKegiatanService}
}

func (h *infoKegiatanHandler) Mount(group *gin.RouterGroup) {
	group.POST("", middleware.ValidateToken(), h.StoreInfoKegiatanHandler)
	group.PATCH("/:id", middleware.ValidateToken(), h.EditInfoKegiatanHandler) 
	group.GET("/:id", h.DetailInfoKegiatanHandler) 
	group.DELETE("/:id", middleware.ValidateToken(), h.DeleteInfoKegiatanHandler) 
	group.GET("", h.FetchInfoKegiatanHandler) 
}

func (h *infoKegiatanHandler) StoreInfoKegiatanHandler(c *gin.Context) {
	var req request.InfoKegiatanRequest

	err := c.ShouldBind(&req)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
		return
	}
	
	link, err := h.infoKegiatanService.UploadImage(c)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, err)
		return
	}
	req.Gambar = link

	infoKegiatan, err := h.infoKegiatanService.StoreInfoKegiatan(&req)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", infoKegiatan)
}

func (h *infoKegiatanHandler) EditInfoKegiatanHandler(c *gin.Context) {
	var req request.InfoKegiatanRequest

	err := c.ShouldBind(&req)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
		return 
	}

	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	link, err := h.infoKegiatanService.UploadImage(c)
	if err != nil {
		infoKegiatan, err := h.infoKegiatanService.GetByID(uint(idUint))
		if err != nil {
			helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
			return 
		}
		req.Gambar = infoKegiatan.Gambar
	} else {
		req.Gambar = link
	}
	
	infoKegiatan, err := h.infoKegiatanService.EditInfoKegiatan(uint(idUint), &req)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", infoKegiatan)
}

func (h *infoKegiatanHandler) FetchInfoKegiatanHandler(c *gin.Context) {
	infoKegiatanList, err := h.infoKegiatanService.FetchInfoKegiatan()
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", infoKegiatanList)
}

func (h *infoKegiatanHandler) DetailInfoKegiatanHandler(c *gin.Context) {
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	infoKegiatan, err := h.infoKegiatanService.GetByID(uint(idUint))
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "", infoKegiatan)
}

func (h *infoKegiatanHandler) DeleteInfoKegiatanHandler(c *gin.Context) {
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	errDelImage := h.infoKegiatanService.DeleteImage(c, uint(idUint))
	if errDelImage != nil {
		helper.ResponseErrorJson(c, http.StatusInternalServerError, errDelImage)
		return
	}

	err := h.infoKegiatanService.DestroyInfoKegiatan(uint(idUint))
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "delete success", "")
}