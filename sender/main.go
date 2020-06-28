package main

import (
	"fmt"
	"github.com/skvoch/k8s-playground/sender/internal/sender"
)

func main(){
	if err := sender.New().Run(); err != nil {
		fmt.Println(err)
	}
}
