package controller

import (
	"net/http"
	"rpl-simple-backend/dto"
	"rpl-simple-backend/service"
	"rpl-simple-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhoneNumberController interface {
	CreatePhoneNumber(ctx *gin.Context)
	GetAllPhoneNumber(ctx *gin.Context)
	GetPhoneNumberById(ctx *gin.Context)
}

type phoneNumberController struct {
	phoneNumberService service.PhoneNumberService
}

func NewPhoneNumberController(phoneNumberService service.PhoneNumberService) PhoneNumberController {
	return &phoneNumberController{
		phoneNumberService: phoneNumberService,
	}
}

func (c *phoneNumberController) CreatePhoneNumber(ctx *gin.Context) {
	var req dto.PhoneNumberRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed to process request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	phoneNumber, err := c.phoneNumberService.CreatePhoneNumber(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed to process request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("success create phone number", phoneNumber)
	ctx.JSON(http.StatusOK, res)
}

func (c *phoneNumberController) GetAllPhoneNumber(ctx *gin.Context) {
	phoneNumbers, err := c.phoneNumberService.GetAllPhoneNumber(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("failed to process request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("success get all phone number", phoneNumbers)
	ctx.JSON(http.StatusOK, res)
}

func (c *phoneNumberController) GetPhoneNumberById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("failed to process request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	phoneNumber, err := c.phoneNumberService.GetPhoneNumberById(ctx, uint64(id))
	if err != nil {
		res := utils.BuildResponseFailed("failed to process request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("success get phone number by id", phoneNumber)
	ctx.JSON(http.StatusOK, res)
}
