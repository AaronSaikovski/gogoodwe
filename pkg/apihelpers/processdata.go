package apihelpers

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

// ProcessData processes the given inverter data by marshaling it to JSON,
// parsing the output, and printing the output using the PrintOutput function
// from the utils package. It returns an error if any of the processing steps
// fail.
//
// Parameters:
// - inverterData: The data to be processed. It should be of type interface{}.
//
// Returns:
// - error: An error if any of the processing steps fail, otherwise nil.
func ProcessData(inverterData interface{}) error {

	// Get data JSON
	dataOutput, err := utils.MarshalStructToJSON(inverterData)
	if err != nil {
		return err
	}

	// Parse output
	output, err := utils.ParseOutput(dataOutput)
	if err != nil {
		return err
	}

	// Print output
	utils.PrintOutput(output)

	return nil
}
