package main

import (
	"fmt"
	"github.com/lucchesisp/go-dev-book/src/config"
	"github.com/lucchesisp/go-dev-book/src/router"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println("Starting server on port:", config.Port)

	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
