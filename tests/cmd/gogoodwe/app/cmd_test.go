package app_test

import (
	"strings"
	"testing"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/app"
	fetchdata "github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata"
	"github.com/spf13/cobra"
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
		{
			name:      "kpidata string",
			input:     "kpidata",
			expected:  6, // KPIData constant
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
		{
			name:      "kpidata numeric",
			input:     "6",
			expected:  6,
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
			result, err := fetchdata.ParseReportType(tt.input)

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
			result, err := fetchdata.ParseReportType(tt.input)

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

			// Verify root command does NOT have RunE (it's a parent command)
			if cmd.RunE != nil {
				t.Error("NewRootCmd() should not have RunE set (it's a parent command)")
			}

			// Verify subcommands exist
			if !cmd.HasSubCommands() {
				t.Error("NewRootCmd() should have subcommands")
			}
		})
	}
}

// TestRootCmdHasSubcommands verifies the root command has the expected subcommands.
func TestRootCmdHasSubcommands(t *testing.T) {
	cmd := app.NewRootCmd("v1.0.0")

	expectedSubcommands := []string{"getdata", "exporthistory"}

	for _, subcmdName := range expectedSubcommands {
		found := false
		for _, subcmd := range cmd.Commands() {
			if subcmd.Name() == subcmdName {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("NewRootCmd() missing expected subcommand: %s", subcmdName)
		}
	}
}

// TestGetDataSubcommand tests the getdata subcommand.
func TestGetDataSubcommand(t *testing.T) {
	rootCmd := app.NewRootCmd("v1.0.0")

	// Find the getdata subcommand
	var getDataCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "getdata" {
			getDataCmd = cmd
			break
		}
	}

	if getDataCmd == nil {
		t.Fatal("getdata subcommand not found")
	}

	// Verify RunE is set
	if getDataCmd.RunE == nil {
		t.Error("getdata subcommand should have RunE set")
	}

	// Verify short description
	if getDataCmd.Short == "" {
		t.Error("getdata subcommand Short description is empty")
	}

	// Verify long description contains report types
	if !contains(getDataCmd.Long, "detail") {
		t.Error("getdata subcommand Long description missing 'detail' report type")
	}

	if !contains(getDataCmd.Long, "kpidata") {
		t.Error("getdata subcommand Long description missing 'kpidata' report type")
	}
}

// TestGetDataFlags tests that the getdata subcommand has all required flags.
func TestGetDataFlags(t *testing.T) {
	rootCmd := app.NewRootCmd("v1.0.0")

	// Find the getdata subcommand
	var getDataCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "getdata" {
			getDataCmd = cmd
			break
		}
	}

	if getDataCmd == nil {
		t.Fatal("getdata subcommand not found")
	}

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
			flag := getDataCmd.Flags().Lookup(tt.flagName)

			if flag == nil {
				t.Errorf("Flag %s not found in getdata subcommand", tt.flagName)
				return
			}

			// Verify short flag
			if flag.Shorthand != tt.short {
				t.Errorf("Flag %s shorthand = %s, want %s", tt.flagName, flag.Shorthand, tt.short)
			}
		})
	}
}

// TestGetDataFlagDefaults tests default values for getdata flags.
func TestGetDataFlagDefaults(t *testing.T) {
	rootCmd := app.NewRootCmd("v1.0.0")

	// Find the getdata subcommand
	var getDataCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "getdata" {
			getDataCmd = cmd
			break
		}
	}

	if getDataCmd == nil {
		t.Fatal("getdata subcommand not found")
	}

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
			flag := getDataCmd.Flags().Lookup(tt.flagName)

			if flag == nil {
				t.Errorf("Flag %s not found in getdata subcommand", tt.flagName)
				return
			}

			if flag.DefValue != tt.expectedValue {
				t.Errorf("Flag %s default value = %s, want %s", tt.flagName, flag.DefValue, tt.expectedValue)
			}
		})
	}
}

// TestExportHistorySubcommand tests the exporthistory subcommand.
func TestExportHistorySubcommand(t *testing.T) {
	rootCmd := app.NewRootCmd("v1.0.0")

	// Find the exporthistory subcommand
	var exportHistoryCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "exporthistory" {
			exportHistoryCmd = cmd
			break
		}
	}

	if exportHistoryCmd == nil {
		t.Fatal("exporthistory subcommand not found")
	}

	// Verify RunE is set
	if exportHistoryCmd.RunE == nil {
		t.Error("exporthistory subcommand should have RunE set")
	}

	// Verify short description
	if exportHistoryCmd.Short == "" {
		t.Error("exporthistory subcommand Short description is empty")
	}
}

// TestExportHistoryFlags tests that the exporthistory subcommand has all required flags.
func TestExportHistoryFlags(t *testing.T) {
	rootCmd := app.NewRootCmd("v1.0.0")

	// Find the exporthistory subcommand
	var exportHistoryCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "exporthistory" {
			exportHistoryCmd = cmd
			break
		}
	}

	if exportHistoryCmd == nil {
		t.Fatal("exporthistory subcommand not found")
	}

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
			flagName: "timestart",
			short:    "s",
		},
		{
			flagName: "timeend",
			short:    "e",
		},
		{
			flagName: "targets",
			short:    "t",
		},
	}

	for _, tt := range tests {
		t.Run(tt.flagName, func(t *testing.T) {
			flag := exportHistoryCmd.Flags().Lookup(tt.flagName)

			if flag == nil {
				t.Errorf("Flag %s not found in exporthistory subcommand", tt.flagName)
				return
			}

			// Verify short flag
			if flag.Shorthand != tt.short {
				t.Errorf("Flag %s shorthand = %s, want %s", tt.flagName, flag.Shorthand, tt.short)
			}
		})
	}
}

// TestRootCmdDescription verifies the root command description.
func TestRootCmdDescription(t *testing.T) {
	cmd := app.NewRootCmd("v1.0.0")

	if cmd.Short == "" {
		t.Error("Root command Short description is empty")
	}

	if cmd.Long == "" {
		t.Error("Root command Long description is empty")
	}
}
