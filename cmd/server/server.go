package main

import (
	"log"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/configs"
)

func main() {
	log.Fatal(configs.StartServer())
}
