package app

import (
	"context"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/spf13/cobra"
)

// runGetData is the main execution function for the getdata command.
func runGetData(cmd *cobra.Command, args []string) error {
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
