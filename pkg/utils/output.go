package utils

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/valyala/fastjson"
)

// parseOutput parses the JSON output from the provided byte slice.
//
// Parameters:
// - dataOutput: a byte slice containing the JSON output to be parsed.
// Return type: (*fastjson.Value, error)
func ParseOutput(dataOutput []byte) (*fastjson.Value, error) {
	// Parse JSON output using reusable parser from pool for better performance
	var parser fastjson.Parser
	return parser.ParseBytes(dataOutput)
}

// printOutput prints the provided fastjson.Value in bright yellow color using the aurora package.
//
// Parameters:
// - output: a pointer to a fastjson.Value that represents the output to be printed.
//
// Return type: None.
func PrintOutput(output *fastjson.Value) {
	fmt.Println(aurora.BrightYellow(output))
}
