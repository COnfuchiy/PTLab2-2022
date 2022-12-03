package main

import (
	"LAB2/tests"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestDatabaseHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(tests.DatabaseHandlerTestSuite))
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(tests.AppTestSuite))
}
