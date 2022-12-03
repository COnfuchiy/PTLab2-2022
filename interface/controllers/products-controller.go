package controllers

import (
	"LAB2/domain"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService  domain.IProductService
	defaultPageSize int
}

func NewProductController(productService domain.IProductService) *ProductController {
	return &ProductController{productService, 10}
}

func (controller ProductController) GetProducts(c *gin.Context) {
	currentPage := 1
	if c.Query("page") != "" {
		pageNum, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err == nil && pageNum > 1 {
			currentPage = int(pageNum)
		}
	}
	totalCount, err := controller.productService.GetProductsCount()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}
	totalPagesCount := int(math.Ceil(float64(totalCount) / float64(controller.defaultPageSize)))
	if currentPage > totalCount {
		currentPage = totalPagesCount
	}
	products, err := controller.productService.GetProducts(currentPage, controller.defaultPageSize)
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products":        products,
		"currentPage":     currentPage,
		"totalPagesCount": totalPagesCount,
	})
}
