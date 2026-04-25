package app

import (
	"context"
	"fmt"
	"time"

	"github.com/AaronSaikovski/gogoodwe/internal/features/exporthistory"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
	"github.com/spf13/cobra"
)

// RunExportHistory is the execution function for the exporthistory command.
func RunExportHistory(cmd *cobra.Command, args []string) error {
	// Create a context with timeout for the API call
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Get flag values from command
	account, _ := cmd.Flags().GetString("account")
	password, _ := cmd.Flags().GetString("password")
	powerstationID, _ := cmd.Flags().GetString("powerstationid")
	qryTimeStart, _ := cmd.Flags().GetString("timestart")
	qryTimeEnd, _ := cmd.Flags().GetString("timeend")
	targetsStr, _ := cmd.Flags().GetString("targets")

	// Check for valid email address input
	if !utils.CheckValidEmail(account) {
		return fmt.Errorf("invalid email address format: should be 'user@somedomain.com'")
	}

	// Check for valid powerstation ID
	if !utils.CheckValidPowerstationID(powerstationID) {
		return fmt.Errorf("invalid Powerstation ID format: should be 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'")
	}

	// Validate date range is within 7 days
	within7Days, err := utils.IsWithin7Days(qryTimeStart, qryTimeEnd)
	if err != nil {
		return fmt.Errorf("invalid date format: %w (expected YYYY-MM-DD HH:MM)", err)
	}
	if !within7Days {
		return fmt.Errorf("date range exceeds maximum of 7 days")
	}

	// Parse targets
	targets, err := exporthistory.ParseTargets(targetsStr)
	if err != nil {
		return err
	}

	// Login
	apiLoginCreds := auth.NewSemsLoginCredentials(account, password, powerstationID)
	var loginService interface {
		SemsLogin(ctx context.Context) (*auth.SemsLoginResponse, error)
	} = &apiLoginCreds

	loginApiResponse, err := loginService.SemsLogin(ctx)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	loginInfo := &auth.LoginInfo{
		SemsLoginCredentials: &apiLoginCreds,
		SemsLoginResponse:    loginApiResponse,
		StartDate:            qryTimeStart,
		EndDate:              qryTimeEnd,
	}

	// Fetch and display the export history data
	return exporthistory.FetchExportHistory(ctx, loginInfo, qryTimeStart, qryTimeEnd, targets)
}
