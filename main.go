package main

import (
	"fmt"
	"github.com/ddvalim/go-qr-code-generator/router"
	"net/http"
)

const port = 8080

func main() {
	fmt.Println(fmt.Sprintf("server running of %d", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router.New()); err != nil {
		panic(err)
	}
}
