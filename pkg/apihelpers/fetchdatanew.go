package apihelpers

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

// EXPERIMENTAL!!
func FetchMonitorAPIDataNew(wg *sync.WaitGroup, authLoginInfo *auth.LoginInfo, powerStationURL string, HTTPTimeout int, inverterOutput interface{}, ch chan<- string) error {
	defer wg.Done() // signal to WaitGroup that this goroutine is done

	// Get the Token header data
	apiResponseJSONData, err := DataTokenJSON(authLoginInfo.SemsLoginResponse)
	if err != nil {
		ch <- fmt.Sprintf("Token header data error : %s", err)
		return err
	}

	// Get the Powerstation ID header data
	powerStationIDJSONData, err := PowerStationIdJSON(authLoginInfo.SemsLoginCredentials)
	if err != nil {
		ch <- fmt.Sprintf("Powerstation header data error : %s", err)
		return err
	}

	// Create URL from the Auth API and append the data URL part
	url := authLoginInfo.SemsLoginResponse.API + powerStationURL

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(powerStationIDJSONData))
	if err != nil {
		ch <- fmt.Sprintf("HTTP request  error : %s", err)
		return err
	}

	// Add headers
	SetHeaders(req, apiResponseJSONData)

	// Make the API call
	client := &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Response from %s: %s", url, resp.Status)
		return err
	}
	defer resp.Body.Close()

	// Get the response body
	respBody, err := utils.FetchResponseBody(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Response body data error : %s", err)
		return err
	}

	// Unmarshal response to struct pointer
	if err := utils.UnmarshalDataToStruct(respBody, inverterOutput); err != nil {
		ch <- fmt.Sprintf("Unmarshal data error : %s", err)
		return err
	}

	return nil
}
