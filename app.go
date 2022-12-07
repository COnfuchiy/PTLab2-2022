package main

import (
	"LAB2/infrastructure/db"
	"LAB2/interface/controllers"
	"LAB2/interface/repository"
	"LAB2/services"
	"LAB2/views"
	"github.com/gin-gonic/gin"
	"html/template"
)

type App struct {
	dbHandler          *repository.IDatabaseHandler
	ginHandler         *gin.Engine
	productController  *controllers.ProductController
	purchaseController *controllers.PurchaseController
}

func NewApp() *App {
	app := App{}
	err := app.setupDatabase()
	if err != nil {
		return nil
	}
	app.setupControllers()
	app.setupServer()
	app.setupRoutes()
	return &app
}

func (app *App) startApp() error {
	return app.ginHandler.Run()
}

func (app *App) setupControllers() {
	productRepository := repository.NewProductRepository(*app.dbHandler)
	purchaseRepository := repository.NewPurchaseRepository(*app.dbHandler)
	discountRepository := repository.NewDiscountRepository(*app.dbHandler)
	productService := services.NewProductService(productRepository)
	purchaseService := services.NewPurchaseService(purchaseRepository)
	discountService := services.NewDiscountService(discountRepository)
	app.productController = controllers.NewProductController(productService, discountService)
	app.purchaseController = controllers.NewPurchaseController(purchaseService, productService, discountService)
}

func (app *App) setupDatabase() error {
	dbHandler, err := db.NewSqliteDatabaseHandler("shop.db")
	if err != nil {
		return err
	}
	app.dbHandler = &dbHandler
	return nil
}

func (app *App) setupServer() {
	app.ginHandler = gin.Default()
	app.ginHandler.SetFuncMap(template.FuncMap{
		"add":       views.Add,
		"sub":       views.Sub,
		"printDate": views.PrintDate,
	})
	app.ginHandler.LoadHTMLGlob("views/*")
}

func (app *App) setupRoutes() {
	app.ginHandler.GET("/", app.productController.GetProducts)
	app.ginHandler.GET("/buy/:id/", app.purchaseController.GetPurchaseForm)
	app.ginHandler.POST("/buy/:id/", app.purchaseController.CreatePurchase)
}
