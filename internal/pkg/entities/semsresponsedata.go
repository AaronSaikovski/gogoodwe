/*
# Name: SemsResponseData - SEMS API Response Data struct
# Contains all the JSON Response data returned from the authentication API - "https://www.semsportal.com/api/v2/Common/CrossLogin"
# Will be unmarshalled to a struct via a pointer// Will be unmarshalled to a struct via a pointer
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package entities

// SemsResponseData - SEMS API Response Data struct
type SemsResponseData struct {
	HasError bool   `json:"hasError"`
	Code     int32  `json:"code"`
	Msg      string `json:"msg"`
	Data     struct {
		UID       string `json:"uid"`
		Timestamp int64  `json:"timestamp"`
		Token     string `json:"token"`
		Client    string `json:"client"`
		Version   string `json:"version"`
		Language  string `json:"language"`
	} `json:"data"`
	Components struct {
		Para         any    `json:"para"`
		LangVer      int    `json:"langVer"`
		TimeSpan     int    `json:"timeSpan"`
		API          string `json:"api"`
		MsgSocketAdr string `json:"msgSocketAdr"`
	} `json:"components"`
	API string `json:"api"`
}
