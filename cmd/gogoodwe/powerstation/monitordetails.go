package powerstation

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/interfaces"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// Generic function to retrieve data from the API via an ISemsDataConstraint Interface of defined structs
func getMonitorData[T interfaces.ISemsDataConstraint](LoginCredentials *types.LoginCredentials, LoginApiResponse *types.LoginResponse, InverterOutput *T) error {

	// get the Token header data
	apiResponseJsonData, err := dataTokenJSON(LoginApiResponse)
	if err != nil {
		return err
	}

	// get the Powerstation ID header data
	powerStationIdJsonData, err := powerStationIdJSON(LoginCredentials)
	if err != nil {
		return err
	}

	//Get the url from the Auth API and append the data url part
	url := (LoginApiResponse.API + PowerStationURL)

	// Create a new http request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationIdJsonData))
	if err != nil {
		return err
	}

	//Add headers pass in the pointer to set the headers on the request object
	setHeaders(req, apiResponseJsonData)

	//make the API Call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//cleanup
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return err
	}

	//marshall response to struct pointer
	inverterDataerr := utils.UnmarshalDataToStruct(respBody, &InverterOutput)
	if inverterDataerr != nil {
		return inverterDataerr
	}

	return nil

}

// Get Monitor Detailed data
func getMonitorDetailByPowerstationId(LoginCredentials *types.LoginCredentials, LoginApiResponse *types.LoginResponse) {
	var powerstationData types.InverterData

	err := getMonitorData(LoginCredentials, LoginApiResponse, &powerstationData)
	if err != nil {
		utils.HandleError(err)
	}

	dataOutput, err := getDataJSON(powerstationData)
	if err != nil {
		utils.HandleError(errors.New("error: converting powerstation data"))
	}

	output, err := parseOutput(dataOutput)
	if err != nil {
		utils.HandleError(err)
	}
	printOutput(output)

}

// Get Monitor sumary data
func getMonitorSummaryByPowerstationId(LoginCredentials *types.LoginCredentials, LoginApiResponse *types.LoginResponse) {

	var powerstationData types.DailySummaryData
	err := getMonitorData(LoginCredentials, LoginApiResponse, &powerstationData)
	if err != nil {
		utils.HandleError(err)
	}

	dataOutput, err := getDataJSON(powerstationData)
	if err != nil {
		utils.HandleError(errors.New("error: converting powerstation summary data"))
	}

	output, err := parseOutput(dataOutput)
	if err != nil {
		utils.HandleError(err)
	}
	printOutput(output)

}
