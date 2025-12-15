package currentkpidata

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/pkg/apihelpers"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

const (

	// Powerstation API Url
	powerStationURL string = "v3/PowerStation/GetMonitorDetailByPowerstationId"

	// Default timeout value
	HTTPTimeout int = 20
)

// GetKPIMonitorData retrieves KPI monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - ctx: context for cancellation
// - authLoginInfo: pointer to the LoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: ([]byte, error) - returns raw JSON bytes and error
func (kpiData *KPIMonitorData) GetMonitorData(ctx context.Context, authLoginInfo *auth.LoginInfo, inverterOutput interface{}) ([]byte, error) {

	return apihelpers.FetchMonitorAPIData(ctx, authLoginInfo, powerStationURL, HTTPTimeout, inverterOutput)
}

// GetPowerData retrieves the power data for a detailed inverter using the provided authentication information.
//
// Parameters:
// - ctx: context for cancellation
// - authLoginInfo: a pointer to the auth.LoginInfo struct containing the login credentials and API response
//
// Returns:
// - error: an error if there was a problem retrieving the power data
func (kpiData *KPIMonitorData) GetPowerData(ctx context.Context, authLoginInfo *auth.LoginInfo) error {

	// Get monitor data and unmarshal into struct
	_, err := kpiData.GetMonitorData(ctx, authLoginInfo, kpiData)
	if err != nil {
		return err
	}

	// Process only the KPI data fields (Data.Kpi)
	return apihelpers.ProcessData(kpiData.Data.Kpi)
}
