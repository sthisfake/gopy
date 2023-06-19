package crud

import (
	nativeModel "v1/models"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func CreateDataRecord(app *pocketbase.PocketBase, output nativeModel.OutputPersianQa) (*models.Record, error) {

	// find the table for it
	collection, err := app.Dao().FindCollectionByNameOrId("data")
	if err != nil {
		return nil, err
	}

	// create a row in the table
	record := models.NewRecord(collection)

	form := forms.NewRecordUpsert(app, record)

	// pocketbase syntax to add things to field
	form.LoadData(map[string]interface{}{
		"data": output,
	})

	// submit the pocketbase form
	if err := form.Submit(); err != nil {
		return nil, err
	}

	return record, nil
}