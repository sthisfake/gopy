package main

import (
	"v1/controllers"

	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)


func main() {
	app := pocketbase.New()

	// api to give a text and some questions as request and the output will be the answer 
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodPost,
			Path:    "/persianqa",
			Handler: controllers.HandlePersianQaRequest(app),
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}




