package monitordetail

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/apihelpers"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
)

const (
	powerStationURL string = "v3/PowerStation/GetMonitorDetailByPowerstationId"
)

// GetMonitorData retrieves monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - ctx: context for cancellation
// - authLoginInfo: pointer to the LoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: ([]byte, error) - returns raw JSON bytes and error
func (summaryData *MonitorData) GetMonitorData(ctx context.Context, authLoginInfo *auth.LoginInfo, inverterOutput interface{}) ([]byte, error) {

	return apihelpers.FetchMonitorAPIData(ctx, authLoginInfo, powerStationURL, inverterOutput)
}

// GetPowerData retrieves the power data for a detailed inverter using the provided authentication information.
//
// Parameters:
// - ctx: context for cancellation
// - authLoginInfo: a pointer to the auth.LoginInfo struct containing the login credentials and API response
//
// Returns:
// - error: an error if there was a problem retrieving the power data
func (detailData *MonitorData) GetPowerData(ctx context.Context, authLoginInfo *auth.LoginInfo) error {

	rawJSON, err := detailData.GetMonitorData(ctx, authLoginInfo, detailData)
	if err != nil {
		return err
	}

	return utils.ProcessRawJSON(rawJSON)
}
