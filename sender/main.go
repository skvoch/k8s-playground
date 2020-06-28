package main

import (
	"fmt"
	"github.com/skvoch/k8s-playground/sender/internal/service"
)

func main(){
	if err := service.New().Run(); err != nil {
		fmt.Println(err)
	}
}
