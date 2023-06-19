////////////////////////////////////// now a custom route that recives json in the body of the request

package main

import (
	"encoding/json"
	"io/ioutil"
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

				// Read the JSON data from the request body
				body, err := ioutil.ReadAll(c.Request().Body)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, "Failed to read request body")
				}

				// Create a new record in the "data" collection
				collection, err := app.Dao().FindCollectionByNameOrId("data")
				if err != nil {
					return err
				}

				record := models.NewRecord(collection)

				form := forms.NewRecordUpsert(app, record)

				// Unmarshal the JSON data into a map[string]interface{}
				var jsonData map[string]interface{}
				err = json.Unmarshal(body, &jsonData)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse JSON data")
				}

				// Set the "data" field of the record with the JSON data
                form.LoadData(map[string]any{
                    "data": jsonData,
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