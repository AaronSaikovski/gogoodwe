package utils

import (
	"fmt"
)

// ProcessRawJSON processes raw JSON bytes by parsing with fastjson and printing colored output.
func ProcessRawJSON(jsonData []byte) error {
	output, err := ParseOutput(jsonData)
	if err != nil {
		return fmt.Errorf("failed to parse output: %w", err)
	}

	PrintOutput(output)

	return nil
}
