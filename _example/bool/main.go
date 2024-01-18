package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"log"
)

func main() {
	randomResult := helper.RandomBool()
	log.Println("result randomBool:", randomResult)
}
