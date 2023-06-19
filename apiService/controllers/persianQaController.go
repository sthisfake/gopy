package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"v1/crud"
	"v1/middleware"
	"v1/models"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func HandlePersianQaRequest(app *pocketbase.PocketBase) echo.HandlerFunc {

	return func(c echo.Context) error {

		// Read the data from the request body
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to read request body")
		}

		// map it to the struct
		var inputData models.RequestPersianQa
		err = json.Unmarshal(body, &inputData)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse JSON data")
		}


		result, err := middleware.ProcessInput(inputData)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process input data")
		}

		// add the data to database
		record, err := crud.CreateDataRecord(app, result)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create data record")
		}

		// response to client
		return c.JSON(http.StatusOK, record)
	}

}