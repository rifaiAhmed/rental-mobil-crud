package api

import (
	"net/http"
	"rental-mobil/constants"
	"rental-mobil/helpers"
	"rental-mobil/internal/interfaces"
	models "rental-mobil/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarAPI struct {
	CarService interfaces.ICarService
}

func (api *CarAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Car
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	if err := api.CarService.CreateCar(c, &req); err != nil {
		log.Error("failed to create car: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, req)
}

func (api *CarAPI) GetAll(c *gin.Context) {
	var (
		log             = helpers.Logger
		objComponent, _ = helpers.ComptServerSidePre(c)
		tipe            = c.Query("type")
	)
	if objComponent.Limit == 0 {
		objComponent.Limit = helpers.GetLimitData()
	}
	obj, err := api.CarService.GetAll(c, objComponent, tipe)
	if err != nil {
		log.Error("failed to get data : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	if len(obj) == 0 {
		helpers.SendResponseHTTP(c, http.StatusOK, "data empty", nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, obj)
}

func (api *CarAPI) Update(c *gin.Context) {
	var (
		log = helpers.Logger
		req *models.Car
	)
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		log.Error("failed to get id : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	req.ID = uint(inputID.ID)
	err = api.CarService.Update(c, req)
	if err != nil {
		log.Error("failed to upadte car: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, req)
}

func (api *CarAPI) Delete(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		log.Error("failed to get id : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	err = api.CarService.Delete(c, inputID.ID)
	if err != nil {
		log.Error("failed to upadte car: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}

func (api *CarAPI) Find(c *gin.Context) {
	var (
		log = helpers.Logger
		obj models.Car
	)
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil || inputID.ID == 0 {
		log.Error("failed to get id : ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	obj, err = api.CarService.FindByID(c, inputID.ID)
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Error("failed to upadte car: ", err)
		helpers.SendResponseHTTP(c, http.StatusOK, "Data not found", nil)
		return
	}
	if err != nil {
		log.Error("failed to upadte car: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, obj)
}
