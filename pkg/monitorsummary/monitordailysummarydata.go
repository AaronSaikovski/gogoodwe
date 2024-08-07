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

package monitorsummary

// DailySummaryData - Struct to hold daily summary data
type DailySummaryData struct {
	Language string `json:"language"`
	HasError bool   `json:"hasError"`
	Msg      string `json:"msg"`
	Code     string `json:"code"`
	Data     struct {
		Kpi struct {
			MonthGeneration float64 `json:"month_generation"`
			Power           float64 `json:"power"`
			TotalPower      float64 `json:"total_power"`
			DayIncome       float64 `json:"day_income"`
			TotalIncome     float64 `json:"total_income"`
			Currency        string  `json:"currency"`
		} `json:"kpi"`
		Inverter []struct {
			TotalGeneration string `json:"total_generation"`
			DailyGeneration string `json:"daily_generation"`
		} `json:"inverter"`
	} `json:"data"`
}

func NewDailySummaryData() *DailySummaryData {
	return &DailySummaryData{}
}
