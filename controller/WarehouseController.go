package controller

import (
	"AnaWarehouseService/request"
	"AnaWarehouseService/response"
	"AnaWarehouseService/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *service.WarehouseService
}

func NewProductController(productService *service.WarehouseService) *ProductController {
	return &ProductController{productService}
}

// @Summary		Transfer Stock of product between warehouse
// @Description	Transfer Stock of product between warehouse
// @Produce		json
// @Param body body request.WarehouseTransferRequest false "body"
// @Success		200	{object} model.Product
// @Router			/warehouse/transfer [post]
func (uc *ProductController) StocksTransferRequest(c *gin.Context) {
	request := request.WarehouseTransferRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	err := uc.productService.StocksTransferRequest(request)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: "Stock are not enough",
		})
		return
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Message: "success",
	})
}

// @Summary		Update status warehouse
// @Description	Update status warehouse
// @Produce		json
// @Param body body request.WarehouseStatus false "body"
// @Success		200	{object} model.Product
// @Router			/warehouse/status [post]
func (uc *ProductController) UpdateStatusWarehouse(c *gin.Context) {
	request := request.WarehouseStatus{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	err := uc.productService.UpdateStatusWarehouse(request)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Message: "success",
	})
}
