package powerstation

import (
	"bytes"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

func getMonitorDetailByPowerstationId(LoginCredentials *types.LoginCredentials, LoginApiResponse *types.LoginResponse) (*types.InverterData, error) {

	// Inverter output struct
	inverterOutputData := types.InverterData{}

	// get the Token header data
	apiResponseJsonData, err := dataTokenJSON(LoginApiResponse)
	if err != nil {
		return nil, err
	}

	// get the Powerstation ID header data
	powerStationIdJsonData, err := powerStationIdJSON(LoginCredentials)
	if err != nil {
		return nil, err
	}

	//Get the url from the Auth API and append the data url part
	url := (LoginApiResponse.API + PowerStationURL)

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationIdJsonData))
	if err != nil {
		return nil, err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req, apiResponseJsonData)

	//make the API Call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return nil, err
	}

	//marshall response to SemsRespInfo struct
	inverterDataerr := utils.UnmarshalDataToStruct(respBody, &inverterOutputData)
	if inverterDataerr != nil {
		return nil, inverterDataerr
	}

	return &inverterOutputData, nil
}
