package app_test

import (
	"strings"
	"testing"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/app"
)

// contains checks if a string contains a substring.
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// TestParseReportType tests the parseReportType function with various inputs.
func TestParseReportType(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  int
		wantError bool
	}{
		// Test string-based report types
		{
			name:      "detail string",
			input:     "detail",
			expected:  0, // Detail constant
			wantError: false,
		},
		{
			name:      "summary string",
			input:     "summary",
			expected:  1, // Summary constant
			wantError: false,
		},
		{
			name:      "point string",
			input:     "point",
			expected:  2, // Point constant
			wantError: false,
		},
		{
			name:      "plant string",
			input:     "plant",
			expected:  3, // Plant constant
			wantError: false,
		},
		{
			name:      "plantchart string",
			input:     "plantchart",
			expected:  4, // PlantChart constant
			wantError: false,
		},
		{
			name:      "powerflow string",
			input:     "powerflow",
			expected:  5, // PowerFlow constant
			wantError: false,
		},
		// Test numeric-based report types
		{
			name:      "detail numeric",
			input:     "0",
			expected:  0,
			wantError: false,
		},
		{
			name:      "summary numeric",
			input:     "1",
			expected:  1,
			wantError: false,
		},
		{
			name:      "point numeric",
			input:     "2",
			expected:  2,
			wantError: false,
		},
		{
			name:      "plant numeric",
			input:     "3",
			expected:  3,
			wantError: false,
		},
		{
			name:      "plantchart numeric",
			input:     "4",
			expected:  4,
			wantError: false,
		},
		{
			name:      "powerflow numeric",
			input:     "5",
			expected:  5,
			wantError: false,
		},
		// Test invalid inputs
		{
			name:      "invalid string",
			input:     "invalid",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "invalid number",
			input:     "99",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "empty string",
			input:     "",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "negative number",
			input:     "-1",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "wrong case",
			input:     "DETAIL",
			expected:  -1,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := app.ParseReportType(tt.input)

			if (err != nil) != tt.wantError {
				t.Errorf("ParseReportType() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError && result != tt.expected {
				t.Errorf("ParseReportType() got %d, want %d", result, tt.expected)
			}

			if tt.wantError && err == nil {
				t.Errorf("ParseReportType() expected error but got none")
			}
		})
	}
}

// TestNewRootCmd tests the NewRootCmd function.
func TestNewRootCmd(t *testing.T) {
	tests := []struct {
		name    string
		version string
	}{
		{
			name:    "with version",
			version: "v1.0.0",
		},
		{
			name:    "empty version",
			version: "",
		},
		{
			name:    "dev version",
			version: "dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := app.NewRootCmd(tt.version)

			// Verify command name
			if cmd.Use != "gogoodwe" {
				t.Errorf("NewRootCmd() Use = %s, want gogoodwe", cmd.Use)
			}

			// Verify short description
			if cmd.Short == "" {
				t.Error("NewRootCmd() Short description is empty")
			}

			// Verify long description
			if cmd.Long == "" {
				t.Error("NewRootCmd() Long description is empty")
			}

			// Verify version
			if cmd.Version != tt.version {
				t.Errorf("NewRootCmd() Version = %s, want %s", cmd.Version, tt.version)
			}

			// Verify RunE is set
			if cmd.RunE == nil {
				t.Error("NewRootCmd() RunE is not set")
			}
		})
	}
}

// TestRootCmdFlags tests that the root command has all required flags.
func TestRootCmdFlags(t *testing.T) {
	cmd := app.NewRootCmd("v1.0.0")

	tests := []struct {
		flagName string
		short    string
	}{
		{
			flagName: "account",
			short:    "a",
		},
		{
			flagName: "password",
			short:    "p",
		},
		{
			flagName: "powerstationid",
			short:    "i",
		},
		{
			flagName: "reporttype",
			short:    "r",
		},
	}

	for _, tt := range tests {
		t.Run(tt.flagName, func(t *testing.T) {
			flag := cmd.Flags().Lookup(tt.flagName)

			if flag == nil {
				t.Errorf("Flag %s not found", tt.flagName)
				return
			}

			// Verify short flag
			if flag.Shorthand != tt.short {
				t.Errorf("Flag %s shorthand = %s, want %s", tt.flagName, flag.Shorthand, tt.short)
			}
		})
	}
}

// TestRootCmdFlagDefaults tests default values for flags.
func TestRootCmdFlagDefaults(t *testing.T) {
	cmd := app.NewRootCmd("v1.0.0")

	tests := []struct {
		flagName      string
		expectedValue string
	}{
		{
			flagName:      "reporttype",
			expectedValue: "detail",
		},
	}

	for _, tt := range tests {
		t.Run(tt.flagName, func(t *testing.T) {
			flag := cmd.Flags().Lookup(tt.flagName)

			if flag == nil {
				t.Errorf("Flag %s not found", tt.flagName)
				return
			}

			if flag.DefValue != tt.expectedValue {
				t.Errorf("Flag %s default value = %s, want %s", tt.flagName, flag.DefValue, tt.expectedValue)
			}
		})
	}
}

// TestParseReportTypeEdgeCases tests edge cases for parseReportType.
func TestParseReportTypeEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  int
		wantError bool
	}{
		{
			name:      "whitespace before",
			input:     " detail",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "whitespace after",
			input:     "detail ",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "mixed case",
			input:     "Detail",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "partial match",
			input:     "det",
			expected:  -1,
			wantError: true,
		},
		{
			name:      "double space",
			input:     "  0",
			expected:  -1,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := app.ParseReportType(tt.input)

			if (err != nil) != tt.wantError {
				t.Errorf("ParseReportType() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError && result != tt.expected {
				t.Errorf("ParseReportType() got %d, want %d", result, tt.expected)
			}
		})
	}
}

// TestRootCmdHasRunE verifies that the root command has a RunE function.
func TestRootCmdHasRunE(t *testing.T) {
	cmd := app.NewRootCmd("v1.0.0")

	if cmd.RunE == nil {
		t.Error("NewRootCmd() RunE function is not set")
	}
}

// TestRootCmdDescription verifies the command descriptions are non-empty.
func TestRootCmdDescription(t *testing.T) {
	cmd := app.NewRootCmd("v1.0.0")

	if cmd.Short == "" {
		t.Error("Root command Short description is empty")
	}

	if cmd.Long == "" {
		t.Error("Root command Long description is empty")
	}

	if !contains(cmd.Long, "detail") {
		t.Error("Root command Long description missing report type examples")
	}
}
