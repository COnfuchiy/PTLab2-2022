package controllers

import (
	"LAB2/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService  domain.IProductService
	discountService domain.IDiscountService
}

func NewProductController(productService domain.IProductService,
	discountService domain.IDiscountService) *ProductController {
	return &ProductController{productService, discountService}
}

func (controller ProductController) GetProducts(c *gin.Context) {
	currentPage := 1
	if c.Query("page") != "" {
		pageNum, err := strconv.Atoi(c.Query("page"))
		if err == nil && pageNum > 1 {
			currentPage = pageNum
		}
	}
	currentPage, totalPagesCount, err := controller.productService.GetPaginationInfo(currentPage)
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}
	products, err := controller.productService.GetProducts(currentPage)
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}
	for i := range products {
		if products[i].Discount.ID != 0 {
			products[i].Discount, err = controller.discountService.CheckDateOrDeleteDiscount(products[i].Discount)
			if err != nil {
				c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
				return
			}
			products[i].NewPrice = controller.discountService.GetPriceWithDiscount(products[i].Price, products[i].Discount.Percent)
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"products":        products,
		"currentPage":     currentPage,
		"totalPagesCount": totalPagesCount,
	})
}
