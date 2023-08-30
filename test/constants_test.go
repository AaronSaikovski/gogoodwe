package testing

import (
	"testing"

	"github.com/AaronSaikovski/gogoodwe/internal/pkg/constants"
)

// TestAuthLoginUrl Test Auth Login Url
func TestAuthLoginUrl(t *testing.T) {

	AuthLoginUrlExpected := "https://www.semsportal.com/api/v2/Common/CrossLogin"
	if constants.AuthLoginUrl != AuthLoginUrlExpected {
		t.Errorf("AuthLoginUrl Const expected '%s' but got '%s'", AuthLoginUrlExpected, constants.AuthLoginUrl)
	}
}

// TestPowerStationURL Test Powerstation API Url
func TestPowerStationURL(t *testing.T) {

	PowerStationURLExpected := "v2/PowerStation/GetMonitorDetailByPowerstationId"
	if constants.PowerStationURL != PowerStationURLExpected {
		t.Errorf("PowerStationURL const expected '%s' but got '%s'", PowerStationURLExpected, constants.PowerStationURL)
	}
}

// HTTPTimeout Test HTTPTimeout value
// func TestHTTPTimeout(t *testing.T) {

// 	HTTPTimeoutPowerStationURLExpected = 20

// 	PowerStationURLExpected := "v2/PowerStation/GetMonitorDetailByPowerstationId"
// 	if constants.HTTPTimeout != PowerStationURLExpected {
// 		t.Errorf("PowerStationURL const expected '%d' but got '%d'", PowerStationURLExpected, constants.PowerStationURL)
// 	}
// }

// Default timeout value

//API login success response message
//SemsLoginSuccessResponse string = "Successful"

// TestSemsLoginSuccessResponse Test SemsLoginSuccessResponse test
func TestSemsLoginSuccessResponse(t *testing.T) {

	SemsLoginSuccessResponseExpected := "Successful"

	if constants.SemsLoginSuccessResponse != SemsLoginSuccessResponseExpected {
		t.Errorf("SemsLoginSuccessResponse const expected '%s' but got '%s'", SemsLoginSuccessResponseExpected, constants.SemsLoginSuccessResponse)
	}
}
