package main

import (
	"net/http"

	"github.com/zedd9/GoWeb/web3/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
