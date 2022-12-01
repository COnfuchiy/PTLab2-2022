package main

import (
	"LAB2/actions"
	"LAB2/controllers"
	"LAB2/infrastructure/db"
	"LAB2/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Add(firstNum, secondNum int) string {
	return fmt.Sprintf("%d", firstNum+secondNum)
}
func Sub(firstNum, secondNum int) string {
	return fmt.Sprintf("%d", firstNum-secondNum)
}

type App struct {
	dbHandler          *db.DatabaseHandler
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
	app.setupProductController()
	app.setupPurchaseController()
	app.setupServer()
	app.setupRoutes()
	return &app
}

func (app *App) startApp() error {
	return app.ginHandler.Run()
}

func (app *App) setupProductController() {
	productRepository := repository.NewProductDatabaseRepository(app.dbHandler)
	productInteractor := actions.NewProductInteractor(productRepository)
	app.productController = controllers.NewProductController(productInteractor)
}

func (app *App) setupPurchaseController() {
	purchaseRepository := repository.NewPurchaseDatabaseRepository(app.dbHandler)
	purchaseInteractor := actions.NewPurchaseInteractor(purchaseRepository)
	app.purchaseController = controllers.NewPurchaseController(purchaseInteractor, app.productController.ProductInteractor)
}

func (app *App) setupDatabase() error {
	dbHandler, err := db.NewSqliteDatabaseHandler("shop.db")
	if err != nil {
		return err
	}
	app.dbHandler = dbHandler
	return nil
}

func (app *App) setupServer() {
	app.ginHandler = gin.Default()
	app.ginHandler.SetFuncMap(template.FuncMap{
		"add": Add,
		"sub": Sub,
	})
	app.ginHandler.LoadHTMLGlob("views/*")
}

func (app *App) setupRoutes() {
	app.ginHandler.GET("/", app.productController.GetProducts)
	app.ginHandler.GET("/buy/:id/", app.purchaseController.GetPurchaseForm)
	app.ginHandler.POST("/buy/:id/", app.purchaseController.CreatePurchase)
}
