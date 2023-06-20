package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// 套件测试, 可设置setUp, tearDown, SetUpSuite, tearDownSuite
// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
