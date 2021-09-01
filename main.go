package main

import (
	"log"

	common "github.com/aashritagoel/go-module/common"
	test "github.com/aashritagoel/go-module/test"
)

func main() {
	commonData := common.GetCommonData()
	log.Println(commonData)

	testData := test.GetTestData()
	log.Println(testData)
}
