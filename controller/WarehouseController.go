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

// @Summary		Get List of product
// @Description	Get List of product
// @Produce		json
// @Param product_name body request.ProductRequest false "product_name"
// @Success		200	{object} model.Product
// @Router			/product [post]
func (uc *ProductController) StocksTransferRequest(c *gin.Context) {
	request := request.WarehouseTransferRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	data, err := uc.productService.StocksTransferRequest(request)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Data:    data,
		Message: "success",
	})
}
