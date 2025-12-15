package powerflow

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/pkg/apihelpers"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

const (

	// Powerstation API Url
	powerStationURL string = "v2/PowerStation/GetPowerflow"

	// Default timeout value
	HTTPTimeout int = 20
)

// GetMonitorData retrieves monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - ctx: context for the request
// - authLoginInfo: pointer to the LoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: []byte, error
func (powerFlowData *Powerflow) GetMonitorData(ctx context.Context, authLoginInfo *auth.LoginInfo, inverterOutput interface{}) ([]byte, error) {

	return apihelpers.FetchMonitorAPIData(ctx, authLoginInfo, powerStationURL, HTTPTimeout, inverterOutput)
}

// GetPowerData retrieves the power data for a daily summary using the provided authentication information.
//
// Parameters:
// - ctx: context for the request
// - authLoginInfo: a pointer to the authentication information for the user.
//
// Returns:
// - error: an error if there was a problem retrieving the power data.
func (powerFlowData *Powerflow) GetPowerData(ctx context.Context, authLoginInfo *auth.LoginInfo) error {

	// Get monitor data
	//rawJSON, err := powerFlowData.GetMonitorData(ctx, authLoginInfo, powerFlowData)
	_, err := powerFlowData.GetMonitorData(ctx, authLoginInfo, powerFlowData)
	if err != nil {
		return err
	}

	//return apihelpers.ProcessRawJSON(rawJSON)

	return apihelpers.ProcessData(powerFlowData)

}
