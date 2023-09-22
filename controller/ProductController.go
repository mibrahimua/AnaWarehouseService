package controller

import (
	"AnaProductService/request"
	"AnaProductService/response"
	"AnaProductService/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService}
}

// @Summary		Get List of product
// @Description	Get List of product
// @Produce		json
// @Param product_name body request.ProductRequest false "product_name"
// @Success		200	{object} model.Product
// @Router			/product [post]
func (uc *ProductController) GetListProductByName(c *gin.Context) {
	request := request.ProductRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	data, err := uc.productService.GetListProductByName(request.Name)
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
