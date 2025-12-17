package app

import (
	"github.com/spf13/cobra"
)

var (
	// Command-line flags
	account        string
	password       string
	powerstationID string
	reportType     string
	versionString  string

	// Export history flags
	qryStart   string
	qryTimeEnd string
	targets    string
)

// NewRootCmd creates and returns the root Cobra command.
func NewRootCmd(version string) *cobra.Command {
	versionString = version

	rootCmd := &cobra.Command{
		Use:     "gogoodwe",
		Short:   "GoGoodwe - A CLI tool to query your SEMS Solar Inverter API",
		Long:    `GoGoodwe is a command-line tool to query the GOODWE SEMS Portal APIs and Solar SEMS API.`,
		Version: version,
	}

	// Add subcommands
	rootCmd.AddCommand(newGetDataCmd())
	rootCmd.AddCommand(newExportHistoryCmd())

	return rootCmd
}

// newGetDataCmd creates and returns the getdata subcommand.
func newGetDataCmd() *cobra.Command {
	getDataCmd := &cobra.Command{
		Use:   "getdata",
		Short: "Get data from SEMS API",
		Long: `Query the GOODWE SEMS Portal APIs and retrieve various types of data.

Report Types:
  detail      - Detailed inverter monitoring data
  summary     - Daily summary data
  point       - All inverter point data
  plant       - Plant details
  plantchart  - Plant power chart data
  powerflow   - Power flow data
  kpidata     - KPI data`,
		RunE: runGetData,
	}

	// Define flags for getdata command
	getDataCmd.Flags().StringVarP(&account, "account", "a", "", "SEMS Email Account (required)")
	getDataCmd.Flags().StringVarP(&password, "password", "p", "", "SEMS Account password (required)")
	getDataCmd.Flags().StringVarP(&powerstationID, "powerstationid", "i", "", "SEMS Powerstation ID (required)")
	getDataCmd.Flags().StringVarP(&reportType, "reporttype", "r", "detail", "Report Type: detail, summary, point, plant, plantchart, powerflow, kpidata")

	// Mark required flags
	getDataCmd.MarkFlagRequired("account")
	getDataCmd.MarkFlagRequired("password")
	getDataCmd.MarkFlagRequired("powerstationid")

	return getDataCmd
}

// newExportHistoryCmd creates and returns the exporthistory subcommand.
func newExportHistoryCmd() *cobra.Command {
	exportHistoryCmd := &cobra.Command{
		Use:   "exporthistory",
		Short: "Export Excel Station History Data",
		Long:  `Export historical data from the SEMS API to Excel format with specified timespan and targets.`,
		RunE:  runExportHistory,
	}

	// Disable flag sorting to preserve definition order
	exportHistoryCmd.Flags().SortFlags = false

	// Define flags for exporthistory command
	exportHistoryCmd.Flags().StringVarP(&account, "account", "a", "", "SEMS Email Account (required)")
	exportHistoryCmd.Flags().StringVarP(&password, "password", "p", "", "SEMS Account password (required)")
	exportHistoryCmd.Flags().StringVarP(&powerstationID, "powerstationid", "i", "", "SEMS Powerstation ID (required)")
	exportHistoryCmd.Flags().StringVarP(&qryStart, "timestart", "s", "", "Query start date/time (format: YYYY-MM-DD HH:MM)")
	exportHistoryCmd.Flags().StringVarP(&qryTimeEnd, "timeend", "e", "", "Query end date/time (format: YYYY-MM-DD HH:MM)")
	exportHistoryCmd.Flags().StringVarP(&targets, "targets", "t", "", "Comma-separated list of targets (e.g., Vpv1,Vpv2,Ipv1)")

	// Mark required flags
	exportHistoryCmd.MarkFlagRequired("account")
	exportHistoryCmd.MarkFlagRequired("password")
	exportHistoryCmd.MarkFlagRequired("powerstationid")
	exportHistoryCmd.MarkFlagRequired("timestart")
	exportHistoryCmd.MarkFlagRequired("timeend")
	exportHistoryCmd.MarkFlagRequired("targets")

	return exportHistoryCmd
}