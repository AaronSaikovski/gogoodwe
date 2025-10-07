package apihelpers

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

var (
	// Reusable HTTP client for better performance
	httpClient = &http.Client{
		Timeout: 0, // Will be set per request
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// Buffer pool for reusing byte buffers
	bufferPool = sync.Pool{
		New: func() interface{} {
			return new(strings.Builder)
		},
	}
)

// FetchMonitorData fetches data from the Monitor API.
//
// It takes in the authentication information, the URL of the power station,
// the HTTP timeout, and a pointer to a struct to store the output.
// It returns an error if there was a problem with the API call.
func FetchMonitorAPIData(authLoginInfo *auth.LoginInfo, powerStationURL string, HTTPTimeout int, inverterOutput interface{}) error {
	// Validate input parameters
	if authLoginInfo == nil || authLoginInfo.SemsLoginResponse == nil || authLoginInfo.SemsLoginCredentials == nil {
		return fmt.Errorf("invalid authentication information")
	}
	if powerStationURL == "" {
		return fmt.Errorf("powerStationURL cannot be empty")
	}

	// Get the Token header data
	apiResponseJSONData, err := DataTokenJSON(authLoginInfo.SemsLoginResponse)
	if err != nil {
		return err
	}

	// //for 'https://au.semsportal.com/api/v2/Charts/GetPlantPowerChart' specific data
	// apiplantPowerResponseJSONData, err := PowerPlantdataTokenJSON(authLoginInfo.SemsLoginResponse)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("apiplantPowerResponseJSONData", string(apiplantPowerResponseJSONData))

	// Get the Powerstation ID header data
	powerStationIDJSONData, err := PowerStationIdJSON(authLoginInfo.SemsLoginCredentials)
	if err != nil {
		return err
	}

	// Create URL from the Auth API and append the data URL part (use pool for better performance)
	builder := bufferPool.Get().(*strings.Builder)
	defer func() {
		builder.Reset()
		bufferPool.Put(builder)
	}()
	builder.Grow(len(authLoginInfo.SemsLoginResponse.API) + len(powerStationURL))
	builder.WriteString(authLoginInfo.SemsLoginResponse.API)
	builder.WriteString(powerStationURL)
	url := builder.String()

	// Create a new HTTP request with pre-sized buffer
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(powerStationIDJSONData))
	if err != nil {
		return err
	}

	// Add headers
	SetHeaders(req, apiResponseJSONData)
	//SetPowerPlantHeaders(req, apiResponseJSONData, apiplantPowerResponseJSONData)

	// Make the API call with reusable client
	httpClient.Timeout = time.Duration(HTTPTimeout) * time.Second
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal response to struct pointer
	if err := utils.UnmarshalDataToStruct(respBody, inverterOutput); err != nil {
		return err
	}

	return nil
}
