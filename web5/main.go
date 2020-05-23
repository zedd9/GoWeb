package main

import (
	"net/http"

	"github.com/zedd9/goweb/web5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
