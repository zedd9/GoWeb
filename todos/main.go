package main

import (
	"log"
	"net/http"
	"os"

	"zeddsWeb/todos/app"
)

func main() {
	port := os.Getenv("PORT")

	m := app.MakeNewHandler(os.Getenv("DATABASE_URL"))
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":"+port, m)
	if err != nil {
		panic(err)
	}
}
