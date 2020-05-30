package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"github.com/zedd9/GoWeb/web16/app"
)

func main() {
	m := app.MakeNewHandler("./test.db")
	defer m.Close()

	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
