package controllers

import (
	"LAB2/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseController struct {
	purchaseService  domain.IPurchaseService
	productService   domain.IProductService
	purchaseHtmlForm string
}

type CreatePurchaseInput struct {
	ProductId uint   `json:"product_id" binding:"required"`
	Person    string `json:"person" binding:"required"`
	Address   string `json:"address" binding:"required"`
}

func NewPurchaseController(purchaseService domain.IPurchaseService,
	productService domain.IProductService) *PurchaseController {
	return &PurchaseController{
		purchaseService,
		productService,
		"purchase_form.html",
	}
}

func (controller PurchaseController) GetPurchaseForm(c *gin.Context) {
	if c.Param("id") != "" {
		c.HTML(http.StatusOK, controller.purchaseHtmlForm, gin.H{
			"productId": c.Param("id"),
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func (controller PurchaseController) viewErrorForm(c *gin.Context, input CreatePurchaseInput, message string) {
	c.HTML(http.StatusBadRequest, controller.purchaseHtmlForm, gin.H{
		"productId": input.ProductId,
		"person":    input.Person,
		"errors":    []string{"Ошибка при работе с формой: " + message + "!"},
	})
}

func (controller PurchaseController) CreatePurchase(c *gin.Context) {
	var input CreatePurchaseInput

	if err := c.Bind(&input); err != nil {
		controller.viewErrorForm(c, input, "ошибка обработки полей формы ("+err.Error()+")")
		return
	}
	product, err := controller.productService.GetProductById(input.ProductId)
	if err != nil {
		controller.viewErrorForm(c, input, err.Error())
		return
	}

	if product.ID == 0 {
		controller.viewErrorForm(c, input, "такой продукт отсутсвует в продаже")
		return
	}

	purchase := domain.Purchase{
		Person:    input.Person,
		Address:   input.Address,
		Price:     product.Price,
		ProductID: input.ProductId,
	}

	err = controller.purchaseService.CreatePurchase(&purchase)
	if err != nil {
		controller.viewErrorForm(c, input, err.Error())
		return
	}

	c.Data(http.StatusOK,
		"text/html; charset=utf-8",
		[]byte("Спасибо за покупку "+input.Person+"!<br><a href='/'>Назад</a>"),
	)
}
