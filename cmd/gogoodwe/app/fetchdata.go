package app

// Main package - This is the main program entry point
import (
	"context"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/interfaces"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"

	"github.com/spf13/cobra"
)

// LoginAndfetchData handles the login and data retrieval process
func loginAndFetchData(ctx context.Context, apiLoginCreds auth.SemsLoginCredentials, ReportType int) error {

	// Assign the login interface
	var loginService interfaces.SemsLogin = &apiLoginCreds

	// Do the login
	loginApiResponse, err := loginService.SemsLogin(ctx)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	//Populate the loginInfo struct
	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: &apiLoginCreds,
		SemsLoginResponse:    loginApiResponse,
	}

	// fetch the data via the interface lookup
	var dataService interfaces.PowerData = lookupMonitorData(ReportType)
	if err := dataService.GetPowerData(ctx, loginInfo); err != nil {
		return fmt.Errorf("data retrieval failed: %w", err)
	}

	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context error: %w", err)
	}

	return nil

}

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
