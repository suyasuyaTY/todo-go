package main

import (
	"fmt"
	"log"

	"my-todilist/controller"
	"my-todilist/model"
)

const port = 8000

func main(){
	err := model.Connect()
	if err !=nil {
		log.Fatal(err.Error())
	}
	router := controller.GetRouter()
	router.Run(fmt.Sprintf(":%d", port))
}