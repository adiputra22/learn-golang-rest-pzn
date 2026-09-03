package test

import (
	"adiputra22/learn-golang-restapi-pzn/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService()
	if err != nil {
		fmt.Println(simpleService.SimpleRepository)
	} else {
		fmt.Println(simpleService.SimpleRepository)
	}
}
