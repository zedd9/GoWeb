package main

import (
	"log"
	"net/http"

	"github.com/zedd9/GoWeb/web16/app"
)

func main() {
	m := app.MakeNewHandler("./test.db")
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
