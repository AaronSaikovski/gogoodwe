package app

import (
	"fmt"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
	"github.com/spf13/cobra"
)

// RunExportHistory is the execution function for the exporthistory command.
func RunExportHistory(cmd *cobra.Command, args []string) error {
	// Get flag values from command
	account, _ := cmd.Flags().GetString("account")
	powerstationID, _ := cmd.Flags().GetString("powerstationid")
	qryStart, _ := cmd.Flags().GetString("timestart")
	qryTimeEnd, _ := cmd.Flags().GetString("timeend")
	targets, _ := cmd.Flags().GetString("targets")

	// Check for valid email address input
	if !utils.CheckValidEmail(account) {
		return fmt.Errorf("invalid email address format: should be 'user@somedomain.com'")
	}

	// Check for valid powerstation ID
	if !utils.CheckValidPowerstationID(powerstationID) {
		return fmt.Errorf("invalid Powerstation ID format: should be 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'")
	}

	// TODO: Implement export history functionality
	// This will need to:
	// 1. Parse the timespan (qryStart, qryTimeEnd)
	// 2. Parse the targets list
	// 3. Call the appropriate API endpoint with context
	// 4. Export data to Excel format

	fmt.Printf("Export History:\n")
	fmt.Printf("  Start Time: %s\n", qryStart)
	fmt.Printf("  End Time: %s\n", qryTimeEnd)
	fmt.Printf("  Targets: %s\n", targets)
	fmt.Printf("  Account: %s\n", account)
	fmt.Printf("  Powerstation ID: %s\n", powerstationID)

	return nil
}
