package controllers

import (
	"LAB2/actions"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductInteractor actions.ProductInteractor
	DefaultPageSize   int
}

func NewProductController(productInteractor actions.ProductInteractor) *ProductController {
	return &ProductController{productInteractor, 10}
}

func (controller ProductController) GetProducts(c *gin.Context) {
	currentPage := 1
	if c.Query("page") != "" {
		pageNum, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err == nil && pageNum > 1 {
			currentPage = int(pageNum)
		}
	}
	totalCount, err := controller.ProductInteractor.CountProducts()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
	}
	totalPagesCount := int(math.Ceil(float64(totalCount) / float64(controller.DefaultPageSize)))
	if currentPage > totalCount {
		currentPage = totalPagesCount
	}
	products, err := controller.ProductInteractor.FindProductsByPagination(currentPage, controller.DefaultPageSize)
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"products":        products,
		"currentPage":     currentPage,
		"totalPagesCount": totalPagesCount,
	})
}
