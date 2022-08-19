package service

import (
	"fmt"
	"net/http"
)

type ExampleService interface {
	GreetHandler(w http.ResponseWriter, r *http.Request)
}

type SimpleExampleService struct{}

var _ ExampleService = (*SimpleExampleService)(nil)

func (svc *SimpleExampleService) GreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}
