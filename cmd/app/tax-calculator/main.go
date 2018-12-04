package main

import (
	"fmt"
	"github.com/alisyahbana/tax-calculator/pkg/common/app"
	"github.com/alisyahbana/tax-calculator/pkg/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	router := httprouter.New()
	SetRoute(router)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	fmt.Println(fmt.Sprintf("Starting Tax-Calculator API HTTP Server on %d", app.GetConfig().Port))
	http.ListenAndServe(fmt.Sprintf(":%d", app.GetConfig().Port), n)
}

func SetRoute(router *httprouter.Router) {
	router.GET("/", handler.InfoHandler)
	router.POST("/create", handler.CreateHandler)
}
