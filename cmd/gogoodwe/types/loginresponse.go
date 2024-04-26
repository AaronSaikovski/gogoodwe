/*
MIT License

# Copyright (c) 2024 Aaron Saikovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

/*
# Name: LoginResponse - SEMS API Response Data struct
# Contains all the JSON Response data returned from the authentication API - "https://www.semsportal.com/api/v2/Common/CrossLogin"
# Will be unmarshalled to a struct via a pointer
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package types

// LoginResponse - SEMS API Response Data struct
type LoginResponse struct {
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
