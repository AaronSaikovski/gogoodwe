/*
MIT License

# Copyright (c) 2024 Aaron Saikovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package apihelpers

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

// ProcesData processes the given inverter data by marshaling it to JSON,
// parsing the output, and printing the output using the PrintOutput function
// from the utils package. It returns an error if any of the processing steps
// fail.
//
// Parameters:
// - inverterData: The data to be processed. It should be of type interface{}.
//
// Returns:
// - error: An error if any of the processing steps fail, otherwise nil.
func ProcesData(inverterData interface{}) error {

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
