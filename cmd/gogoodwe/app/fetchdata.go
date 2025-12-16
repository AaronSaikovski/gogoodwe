package app

// Main package - This is the main program entry point
import (
	"context"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/interfaces"
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

	//HERE

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
