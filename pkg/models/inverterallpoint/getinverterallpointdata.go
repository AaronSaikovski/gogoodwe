package inverteallpoint

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/pkg/apihelpers"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

const (

	// Powerstation API Url
	powerStationURL string = "/v3/PowerStation/GetInverterAllPoint"

	// Default timeout value
	HTTPTimeout int = 20
)

// GetMonitorData retrieves monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - ctx: context for cancellation
// - authLoginInfo: pointer to the LoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: ([]byte, error)
func (inverterData *InverterAllPoint) GetMonitorData(ctx context.Context, authLoginInfo *auth.LoginInfo, inverterOutput interface{}) ([]byte, error) {

	return apihelpers.FetchMonitorAPIData(ctx, authLoginInfo, powerStationURL, HTTPTimeout, inverterOutput)
}

// GetPowerData retrieves the power data for a daily summary using the provided authentication information.
//
// Parameters:
// - ctx: context for cancellation
// - authLoginInfo: a pointer to the authentication information for the user.
//
// Returns:
// - error: an error if there was a problem retrieving the power data.
func (inverterData *InverterAllPoint) GetPowerData(ctx context.Context, authLoginInfo *auth.LoginInfo) error {

	// Get monitor data (returns raw JSON to avoid double marshaling)
	rawJSON, err := inverterData.GetMonitorData(ctx, authLoginInfo, inverterData)
	if err != nil {
		return err
	}

	// Process raw JSON directly without remarshaling
	return apihelpers.ProcessRawJSON(rawJSON)

}
