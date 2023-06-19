////////////////////////////////////////// now a rother that runs the python script in the background and saves the output in the collection

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

type requestBody struct{
	Text string `json:"text"`
}

type Input struct {
	Text string `json:"input"`
}

type Output struct {
	Input  string  `json:"input"`
	Result string `json:"result"`
}

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodPost,
			Path:    "/dummy",
			Handler: handleRequest(app),
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func handleRequest(app *pocketbase.PocketBase)

func handleDummy(app *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read the JSON data from the request body
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to read request body")
		}

		// Unmarshal the JSON data into the Input struct
		var inputData requestBody
		err = json.Unmarshal(body, &inputData)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse JSON data")
		}

		// Map the requestStruct into Input Struct

		input := Input(inputData)

		// Process the input data and get the result
		result, err := processInput(input)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process input data")
		}

		// Create a new record in the "data" collection
		record, err := createDataRecord(app, result)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create data record")
		}

		// Return the created record as a response
		return c.JSON(http.StatusOK, record)
	}
}

func processInput(input Input) (Output, error) {
	// Write input data to JSON file
	inputJSON, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return Output{}, err
	}

	err = ioutil.WriteFile("data.json", inputJSON, 0644)
	if err != nil {
		return Output{}, err
	}

	// Execute Python script
	cmd := exec.Command("python", "script.py")
	err = cmd.Run()
	if err != nil {
		return Output{}, err
	}

	// Read output data from JSON file
	outputData, err := ioutil.ReadFile("data.json")
	if err != nil {
		return Output{}, err
	}

	// Unmarshal the JSON data into the Output struct
	var output Output
	err = json.Unmarshal(outputData, &output)
	if err != nil {
		return Output{}, err
	}

	return output, nil
}

func createDataRecord(app *pocketbase.PocketBase, output Output) (*models.Record, error) {

	// Find the "data" collection
	collection, err := app.Dao().FindCollectionByNameOrId("data")
	if err != nil {
		return nil, err
	}

	// Create a new record in the collection
	record := models.NewRecord(collection)

	form := forms.NewRecordUpsert(app, record)

	// Set the "title" field of the record
	form.LoadData(map[string]interface{}{
		"data": output,
	})

	if err := form.Submit(); err != nil {
		return nil, err
	}

	return record, nil
}
