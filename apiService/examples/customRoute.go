//custom route teseted ok , with only a title field , auto generated when sending the request

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/dummy",
			Handler: func(c echo.Context) error {
				// Create a new record in the "data" collection
				collection, err := app.Dao().FindCollectionByNameOrId("data")
				if err != nil {
					return err
				}

				record := models.NewRecord(collection)

				form := forms.NewRecordUpsert(app, record)

				form.LoadData(map[string]interface{}{
					"title": "Dummy Title",
				})

				if err := form.Submit(); err != nil {
					return err
				}

				// Return the created record as a response
				return c.JSON(http.StatusOK, record)
			},
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}