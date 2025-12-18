package fetchdata

import (
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
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
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	// Parse output
	output, err := utils.ParseOutput(dataOutput)
	if err != nil {
		return fmt.Errorf("failed to parse output: %w", err)
	}

	// Print output
	utils.PrintOutput(output)

	return nil
}

// ProcessRawJSON processes raw JSON bytes directly without remarshaling.
// This is more efficient when you already have JSON bytes.
//
// Parameters:
// - jsonData: The raw JSON bytes to be processed.
//
// Returns:
// - error: An error if any of the processing steps fail, otherwise nil.
func ProcessRawJSON(jsonData []byte) error {
	// Parse output directly from raw JSON
	output, err := utils.ParseOutput(jsonData)
	if err != nil {
		return fmt.Errorf("failed to parse output: %w", err)
	}

	// Print output
	utils.PrintOutput(output)

	return nil
}
