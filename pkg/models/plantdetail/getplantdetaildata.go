package plantdetail

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/apihelpers"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

const (

	// Powerstation API Url
	powerStationURL string = "/v3/PowerStation/GetPlantDetailByPowerstationId"

	// Default timeout value
	HTTPTimeout int = 20
)

// GetMonitorData retrieves monitor data using login credentials and response, storing it in inverterOutput.
//
// Parameters:
// - authLoginInfo: pointer to the LoginInfo struct containing the login credentials and API response
// - inverterOutput: pointer to the data output
// Return type: error
func (plantDetailData *PlantDetailByPowerstationId) GetMonitorData(authLoginInfo *auth.LoginInfo, inverterOutput interface{}) error {

	return apihelpers.FetchMonitorAPIData(authLoginInfo, powerStationURL, HTTPTimeout, inverterOutput)
}

// GetPowerData retrieves the power data for a daily summary using the provided authentication information.
//
// Parameters:
// - authLoginInfo: a pointer to the authentication information for the user.
//
// Returns:
// - error: an error if there was a problem retrieving the power data.
func (plantDetailData *PlantDetailByPowerstationId) GetPowerData(authLoginInfo *auth.LoginInfo) error {

	// Get monitor data
	if err := plantDetailData.GetMonitorData(authLoginInfo, plantDetailData); err != nil {
		return err
	}

	return apihelpers.ProcessData(plantDetailData)

}
