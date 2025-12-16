package app

import (
	"context"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/spf13/cobra"
)

var (
	// Command-line flags
	account        string
	password       string
	powerstationID string
	reportType     string
	versionString  string
)

// NewRootCmd creates and returns the root Cobra command.
func NewRootCmd(version string) *cobra.Command {
	versionString = version

	rootCmd := &cobra.Command{
		Use:   "gogoodwe",
		Short: "GoGoodwe - A CLI tool to query your SEMS Solar Inverter API",
		Long: `GoGoodwe is a command-line tool to query the GOODWE SEMS Portal APIs and Solar SEMS API.

Report Types:
  detail      - Detailed inverter monitoring data
  summary     - Daily summary data
  point       - All inverter point data
  plant       - Plant details
  plantchart  - Plant power chart data
  powerflow   - Power flow data,
  kpidata   - KPI data `,
		Version: version,
		RunE:    runRoot,
	}

	// Define flags
	rootCmd.Flags().StringVarP(&account, "account", "a", "", "SEMS Email Account (required)")
	rootCmd.Flags().StringVarP(&password, "password", "p", "", "SEMS Account password (required)")
	rootCmd.Flags().StringVarP(&powerstationID, "powerstationid", "i", "", "SEMS Powerstation ID (required)")
	rootCmd.Flags().StringVarP(&reportType, "reporttype", "r", "detail", "Report Type: detail, summary, point, plant, plantchart, powerflow, kpidata")

	// Mark required flags
	rootCmd.MarkFlagRequired("account")
	rootCmd.MarkFlagRequired("password")
	rootCmd.MarkFlagRequired("powerstationid")

	return rootCmd
}

// runRoot is the main execution function for the root command.
func runRoot(cmd *cobra.Command, args []string) error {
	// Create a context for the API call
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Check for valid email address input
	if !utils.CheckValidEmail(account) {
		return fmt.Errorf("invalid email address format: should be 'user@somedomain.com'")
	}

	// Check for valid powerstation ID
	if !utils.CheckValidPowerstationID(powerstationID) {
		return fmt.Errorf("invalid Powerstation ID format: should be 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'")
	}

	// Convert report type string to integer
	reportTypeInt, err := ParseReportType(reportType)
	if err != nil {
		return err
	}

	// User account struct instance
	apiLoginCreds := auth.NewSemsLoginCredentials(account, password, powerstationID)

	// Get the data from the API, return any errors
	return loginAndFetchData(ctx, apiLoginCreds, reportTypeInt)
}

// // ParseReportType converts a string report type to its integer constant.
// func ParseReportType(reportTypeStr string) (int, error) {
// 	switch reportTypeStr {
// 	case "detail", "0":
// 		return Detail, nil
// 	case "summary", "1":
// 		return Summary, nil
// 	case "point", "2":
// 		return Point, nil
// 	case "plant", "3":
// 		return Plant, nil
// 	case "plantchart", "4":
// 		return PlantChart, nil
// 	case "powerflow", "5":
// 		return PowerFlow, nil
// 	case "kpidata", "6":
// 		return KPIData, nil
// 	default:
// 		return -1, fmt.Errorf("invalid report type '%s'. Valid options are: detail, summary, point, plant, plantchart, powerflow, kpidata", reportTypeStr)
// 	}
// }
