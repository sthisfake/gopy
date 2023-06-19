// package middleware

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os/exec"
// 	"v1/models"
// )

// func ProcessInput(input models.RequestPersianQa) (models.OutputPersianQa, error) {
// 	// write input data to json file
// 	// for now we use a json file to comminicate between go and python
// 	inputJson, err := json.MarshalIndent(input, "", "  ")
// 	if err != nil {
// 		return models.OutputPersianQa{}, err
// 	}

// 	err = ioutil.WriteFile("middleware/pythonCode/data.json", inputJson, 0644)
// 	if err != nil {
// 		return models.OutputPersianQa{}, err
// 	}

// 	// Create a command session
// 	cmd := exec.Command("cmd.exe", "/C", `cd C:\Users\Pouya\Desktop\kimia pars\python\venv\Scripts && activate.bat && python ../../../python/perisanQa/main.py`)

// 	var stderr bytes.Buffer
// 	cmd.Stderr = &stderr

// 	// Run the command session
// 	err = cmd.Run()
// 	if err != nil {
// 		fmt.Println("Failed to start the command session:", err)
// 		return models.OutputPersianQa{}, err
// 	}

// 	// read the result from json
// 	outputData, err := ioutil.ReadFile("../../perisanQa/data.json")
// 	if err != nil {
// 		fmt.Println("hereeeeeee session error:", stderr.String())
// 		return models.OutputPersianQa{}, err
// 	}

// 	// map the json into the struct
// 	var output models.OutputPersianQa
// 	err = json.Unmarshal(outputData, &output)
// 	if err != nil {
// 		fmt.Println("or hereeeeeee session error:", stderr.String())
// 		return models.OutputPersianQa{}, err
// 	}

// 	return output, nil
// }

package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"v1/models"

	"github.com/go-cmd/cmd"
)

func ProcessInput(input models.RequestPersianQa) (models.OutputPersianQa, error) {
	// Write input data to a JSON file
	// For now, we use a JSON file to communicate between Go and Python
	inputJson, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return models.OutputPersianQa{}, err
	}

	err = ioutil.WriteFile("middleware/pythonCode/data.json", inputJson, 0644)
	if err != nil {
		return models.OutputPersianQa{}, err
	}

	// Create a new command with the desired command and arguments
	command := cmd.NewCmd("cmd.exe", "/C", "cd", "C:\\Users\\Pouya\\Desktop\\kimia pars\\python\\venv\\Scripts", "&&", "activate.bat", "&&", "python", "../../../python/perisanQa/main.py")

	// Start the command
	_  = command.Start()

	// // Wait for the command to complete with a timeout
	// timeout := 60 * time.Second
	// select {
	// case <-statusChan:
	// 	// The command completed successfully
	// 	// ... Your code to process the output ...
	// case <-time.After(timeout):
	// 	// The command timed out
	// 	// ... Your code to handle the timeout ...
	// }

	// ... Your code to handle the command's output ...

	// Read the result from the JSON file
	outputData, err := ioutil.ReadFile("../../perisanQa/data.json")
	if err != nil {
		fmt.Println("Error reading output file:", err)
		return models.OutputPersianQa{}, err
	}

	// Map the JSON into the struct
	var output models.OutputPersianQa
	err = json.Unmarshal(outputData, &output)
	if err != nil {
		fmt.Println("Error unmarshaling output:", err)
		return models.OutputPersianQa{}, err
	}

	return output, nil
}
