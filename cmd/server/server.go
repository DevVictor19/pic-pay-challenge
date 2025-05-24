package main

import (
	"log"

	"github.com/DevVictor19/pic-pay-challenge/internal/infra/server"
)

func main() {
	log.Fatal(server.Start())
}
