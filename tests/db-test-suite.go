package tests

import (
	"LAB2/domain"
	"LAB2/infrastructure/db"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type TestDatabase struct {
	DBHandler                  *db.DatabaseHandler
	testSqliteDatabaseFileName string
}

func NewTestDatabase() *TestDatabase {
	testDB := TestDatabase{}
	testDB.testSqliteDatabaseFileName = "test.db"
	_, err := os.Stat(testDB.testSqliteDatabaseFileName)
	if err == nil {
		err := os.Remove(testDB.testSqliteDatabaseFileName)
		if err != nil {
			log.Fatalln(err)
		}
	}
	DB, err := gorm.Open(sqlite.Open(testDB.testSqliteDatabaseFileName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalln(err)
	}
	if res := DB.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		log.Fatalln(res.Error)
	}
	testDB.DBHandler = &db.DatabaseHandler{DB: DB}
	err = DB.AutoMigrate(&domain.Product{}, &domain.Purchase{})
	if err != nil {
		log.Fatalln(err)
	}
	return &testDB
}

func (testDatabase *TestDatabase) CloseAndRemoveDatabase() {
	dbInstance, _ := testDatabase.DBHandler.DB.DB()
	err := dbInstance.Close()
	if err != nil {
		return
	}
	err = os.Remove(testDatabase.testSqliteDatabaseFileName)
	if err != nil {
		log.Println(err)
	}
}

type DatabaseHandlerTestSuite struct {
	suite.Suite
	TestDatabaseHandler *TestDatabase
	products            []domain.Product
	purchase            domain.Purchase
}

func (suite *DatabaseHandlerTestSuite) SetupSuite() {
	suite.products = []domain.Product{
		{
			ID:    1,
			Name:  "Pivo",
			Price: 100,
		}, {
			ID:    2,
			Name:  "Riba",
			Price: 150,
		}, {
			ID:    3,
			Name:  "Myaso",
			Price: 200,
		},
	}
	suite.purchase = domain.Purchase{
		Person:    "Yra",
		Address:   "Moskva",
		Price:     100,
		ProductID: 1,
	}
	suite.TestDatabaseHandler = NewTestDatabase()
	result := suite.TestDatabaseHandler.DBHandler.DB.Create(suite.products)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
}

func (suite *DatabaseHandlerTestSuite) TestCountProducts() {
	productCount, err := suite.TestDatabaseHandler.DBHandler.CountProducts()
	suite.Require().Nil(err)
	suite.Require().Equal(len(suite.products), productCount)
}

func (suite *DatabaseHandlerTestSuite) TestFindProductById() {
	expectedProduct := suite.products[0]
	actualProduct, err := suite.TestDatabaseHandler.DBHandler.FindProductById(expectedProduct.ID)
	suite.Require().Nil(err)
	suite.Require().Equal(expectedProduct.ID, actualProduct.ID)
	nonExistsProduct, err := suite.TestDatabaseHandler.DBHandler.FindProductById(0)
	suite.Require().Equal(uint(0), nonExistsProduct.ID)
	_, err = suite.TestDatabaseHandler.DBHandler.FindProductById(suite.products[len(suite.products)-1].ID * 256)
	suite.Require().NotNil(err)
}

func (suite *DatabaseHandlerTestSuite) TestFindProductsByPagination() {
	page, perPage := 1, 3
	products, err := suite.TestDatabaseHandler.DBHandler.FindProductsByPagination(page, perPage)
	suite.Require().Nil(err)
	suite.Require().Equal(3, len(products))
	page, perPage = 2, 2
	products, err = suite.TestDatabaseHandler.DBHandler.FindProductsByPagination(page, perPage)
	suite.Require().Nil(err)
	suite.Require().Equal(1, len(products))
}

func (suite *DatabaseHandlerTestSuite) TestInsertPurchase() {
	err := suite.TestDatabaseHandler.DBHandler.InsertPurchase(&suite.purchase)
	suite.Require().Nil(err)
	errorPurchase := suite.purchase
	errorPurchase.ProductID = 562165
	errorPurchase.ID = 2
	err = suite.TestDatabaseHandler.DBHandler.InsertPurchase(&errorPurchase)
	suite.Require().NotNil(err)
}

func (suite *DatabaseHandlerTestSuite) TearDownSuite() {
	suite.TestDatabaseHandler.CloseAndRemoveDatabase()
}
