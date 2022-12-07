package tests

import (
	"LAB2/domain"
	"LAB2/interface/controllers"
	"LAB2/mocks"
	"LAB2/views"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/html"
	"html/template"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AppTestSuite struct {
	suite.Suite
	products            []domain.Product
	purchase            domain.Purchase
	mockProductService  *mocks.MockProductService
	mockPurchaseService *mocks.MockPurchaseService
	mockDiscountService *mocks.MockDiscountService
	ginHandler          *gin.Engine
	totalProductsCount  int
	productPerPage      int
	pagesToTest         []int
	productController   *controllers.ProductController
	purchaseController  *controllers.PurchaseController
	purchasedProduct    domain.Product
	purchasedProductUrl string
}

func (suite *AppTestSuite) SetupSuite() {
	suite.totalProductsCount = 21
	suite.productPerPage = 10
	totalPagesCount := int(suite.totalProductsCount/suite.productPerPage) + 1
	suite.Require().GreaterOrEqual(3, totalPagesCount)
	suite.pagesToTest = []int{
		1, 2, totalPagesCount,
	}

	suite.mockDiscountService = new(mocks.MockDiscountService)

	suite.products = []domain.Product{
		{
			ID:    1,
			Name:  "Pivo",
			Count: 1000,
			Price: 100,
			Discount: domain.Discount{
				ID:        1,
				Percent:   15,
				EndDate:   time.Now().AddDate(0, 0, -14),
				ProductID: 1,
			},
		}, {
			ID:    2,
			Name:  "Pivo",
			Count: 1000,
			Price: 100,
			Discount: domain.Discount{
				ID:        2,
				Percent:   10,
				EndDate:   time.Now().AddDate(0, 0, 14),
				ProductID: 2,
			},
		},
	}

	for i := 2; i < suite.totalProductsCount; i++ {
		addedProduct := domain.Product{
			ID:       suite.products[i-1].ID + 1,
			Name:     suite.products[i-1].Name,
			Price:    suite.products[i-1].Price,
			Discount: domain.Discount{},
		}
		suite.products = append(suite.products, addedProduct)
	}

	for _, product := range suite.products {
		suite.mockDiscountService.
			On("GetPriceWithDiscount", product.Price, product.Discount.Percent).
			Return(suite.calcPriceWithDiscount(product.Price, product.Discount.Percent))
		if product.Discount.ID != 0 && time.Now().After(product.Discount.EndDate) {
			suite.mockDiscountService.
				On("CheckDateOrDeleteDiscount", product.Discount).
				Return(domain.Discount{}, nil)
		} else {
			suite.mockDiscountService.
				On("CheckDateOrDeleteDiscount", product.Discount).
				Return(product.Discount, nil)
		}

	}

	suite.mockProductService = new(mocks.MockProductService)
	for _, page := range suite.pagesToTest {
		startSlice := math.Min(float64(suite.totalProductsCount-1), float64((page-1)*suite.productPerPage))
		endSlice := math.Min(float64(suite.totalProductsCount), startSlice+float64(suite.productPerPage))
		suite.mockProductService.On("GetProducts", page).Return(
			suite.products[int(startSlice):int(endSlice)], nil).Once()
	}
	suite.productController = controllers.NewProductController(suite.mockProductService, suite.mockDiscountService)

	suite.purchasedProduct = suite.products[rand.Intn(len(suite.products))]
	decreasedCountProduct := suite.purchasedProduct
	decreasedCountProduct.Count -= 1

	suite.mockProductService.
		On("DecreaseProductCount", &suite.purchasedProduct).
		Return(nil)

	// for purchase test success and error mock GetProductById
	suite.mockProductService.On("GetProductById",
		suite.purchasedProduct.ID).Return(&suite.purchasedProduct, nil).Once()

	suite.mockProductService.On("GetProductById",
		suite.products[len(suite.products)-1].ID+1).Return(&domain.Product{}, nil)

	suite.purchasedProductUrl = "/buy/" + fmt.Sprintf("%d", suite.purchasedProduct.ID) + "/"
	suite.purchase = domain.Purchase{
		Person:    "Eugeniy",
		Address:   "Moskva",
		Price:     suite.calcPriceWithDiscount(suite.purchasedProduct.Price, suite.purchasedProduct.Discount.Percent),
		ProductID: suite.purchasedProduct.ID,
	}
	suite.mockPurchaseService = new(mocks.MockPurchaseService)
	suite.mockPurchaseService.On("CreatePurchase", &suite.purchase).Return(nil)
	suite.mockPurchaseService.
		On("CheckProductWillHaveDiscount", &suite.purchasedProduct).
		Return(false, nil)
	suite.purchaseController = controllers.NewPurchaseController(suite.mockPurchaseService,
		suite.mockProductService,
		suite.mockDiscountService)

	suite.setupGin()
}

func (suite *AppTestSuite) TestGetProducts() {
	for _, page := range suite.pagesToTest {
		suite.mockProductService.On("GetPaginationInfo", page).Return(page, len(suite.pagesToTest), nil)
		suite.Run(fmt.Sprintf("PaginationTest/Page%d", page), func() {
			suite.testGetProductsPagination(page)
		})
	}
}

func (suite *AppTestSuite) TestCreatePurchase() {
	suite.Run("SuccessCreateTest", suite.testSuccessCreatePurchase)
	suite.Run("ErrorCreateTest", suite.testErrorProductIDCreatePurchase)
	suite.Run("InvalidInputDataTest", suite.testInvalidInputCreatePurchase)
}

func (suite *AppTestSuite) testGetProductsPagination(page int) {
	getProductUrl := "/?page=" + fmt.Sprintf("%d", page)
	response := suite.fetchTestRequest("GET", getProductUrl, nil)
	document := suite.getQueryDocumentFromResponse(response)
	currentPageProductCount := math.Min(
		float64(suite.productPerPage),
		float64(suite.totalProductsCount-((page-1)*suite.productPerPage)))
	suite.Require().Equal(int(currentPageProductCount), document.Find(".product").Size())

	currentPage := document.Find(".current-page").First().Text()
	currentPageAsInt, err := strconv.Atoi(currentPage)
	suite.Require().Nil(err)
	suite.Require().Equal(currentPageAsInt, page)

	if page != 1 {
		suite.Require().NotEqual(0, document.Find(".first-page").Size())
	} else {
		suite.Require().Zero(document.Find(".first-page").Size())
	}

	if page != len(suite.pagesToTest) {
		suite.Require().NotEqual(0, document.Find(".last-page").Size())
	} else {
		suite.Require().Zero(document.Find(".last-page").Size())
	}
	document.Find(".product").Each(func(i int, selection *goquery.Selection) {
		discountText := selection.Find(".discount").First().Text()
		if strings.Contains(discountText, "%") {
			productPrice := selection.Find(".price").First().Text()
			productPriceAsInt, err := strconv.Atoi(productPrice)
			suite.Require().Nil(err)
			product := suite.products[suite.productPerPage*(page-1)+i]
			suite.Require().Equal(product.Price, uint(productPriceAsInt))
		}
	})
}

func (suite *AppTestSuite) testGetPurchaseForm() {
	response := suite.fetchTestRequest("GET", suite.purchasedProductUrl, nil)
	document := suite.getQueryDocumentFromResponse(response)
	productId, exist := document.Find("[name=ProductId]").First().Attr("value")
	suite.Require().True(exist)
	productIdAsInt, err := strconv.Atoi(productId)
	suite.Require().Nil(err)
	suite.Require().Equal(suite.purchasedProduct.ID, productIdAsInt)
}

func (suite *AppTestSuite) testSuccessCreatePurchase() {
	postForm := suite.getCreatePurchasePostForm(
		fmt.Sprintf("%d", suite.purchasedProduct.ID),
		suite.purchase.Person,
		suite.purchase.Address)

	responseData := suite.fetchTestRequest("POST", suite.purchasedProductUrl, strings.NewReader(postForm.Encode()))
	suite.Require().Equal("Спасибо за покупку "+suite.purchase.Person+"!<br><a href='/'>Назад</a>",
		string(responseData))
}

func (suite *AppTestSuite) testErrorProductIDCreatePurchase() {
	postForm := suite.getCreatePurchasePostForm(
		fmt.Sprintf("%d", suite.products[len(suite.products)-1].ID+1),
		suite.purchase.Person,
		suite.purchase.Address)
	responseData := suite.fetchTestRequest("POST", suite.purchasedProductUrl, strings.NewReader(postForm.Encode()))
	document := suite.getQueryDocumentFromResponse(responseData)
	errorText := document.Find(".error").First().Text()
	suite.Require().Equal("Ошибка: такой продукт отсутсвует в продаже!", errorText)
}

func (suite *AppTestSuite) testInvalidInputCreatePurchase() {
	postForm := suite.getCreatePurchasePostForm(
		suite.purchase.Person,
		suite.purchase.Person,
		suite.purchase.Address)
	responseData := suite.fetchTestRequest("POST", suite.purchasedProductUrl, strings.NewReader(postForm.Encode()))
	document := suite.getQueryDocumentFromResponse(responseData)
	errorText := document.Find(".error").First().Text()
	suite.Require().Equal("Ошибка: ошибка обработки полей формы (strconv.ParseUint: parsing \""+suite.purchase.Person+"\": invalid syntax)!", errorText)
}

func (suite *AppTestSuite) testGetProductsWithDiscount() {
	productWithDiscount := suite.products[0]
	productWithDiscount.Discount = domain.NewDefaultDiscount(productWithDiscount.ID)
	//productWithDiscount.
}

func (suite *AppTestSuite) fetchTestRequest(method, url string, body io.Reader) []byte {
	request, err := http.NewRequest(method, url, body)
	suite.Require().Nil(err)
	if method == "POST" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	responseRecorder := httptest.NewRecorder()
	suite.ginHandler.ServeHTTP(responseRecorder, request)
	responseData, err := io.ReadAll(responseRecorder.Body)
	suite.Require().Nil(err)
	return responseData
}

func (suite *AppTestSuite) getCreatePurchasePostForm(productID, person, address string) url.Values {
	postForm := url.Values{}
	postForm.Add("ProductId", productID)
	postForm.Add("Person", person)
	postForm.Add("Address", address)
	return postForm
}

func (suite *AppTestSuite) getQueryDocumentFromResponse(response []byte) *goquery.Document {
	responseAsHtml, err := html.Parse(bytes.NewReader(response))
	suite.Require().Nil(err)
	document := goquery.NewDocumentFromNode(responseAsHtml)
	return document
}

func (suite *AppTestSuite) setupGin() {
	suite.ginHandler = gin.Default()
	suite.ginHandler.SetFuncMap(template.FuncMap{
		"add":       views.Add,
		"sub":       views.Sub,
		"printDate": views.PrintDate,
	})
	suite.ginHandler.LoadHTMLGlob("views/*")
	suite.ginHandler.GET("/", suite.productController.GetProducts)
	suite.ginHandler.GET("/buy/:id/", suite.purchaseController.GetPurchaseForm)
	suite.ginHandler.POST("/buy/:id/", suite.purchaseController.CreatePurchase)
}

func (suite *AppTestSuite) calcPriceWithDiscount(price, percent uint) uint {
	return uint(float64(price) * float64(100-percent) / 100.0)
}

func (suite *AppTestSuite) TearDownSuite() {
	suite.mockProductService.AssertExpectations(suite.T())
}
