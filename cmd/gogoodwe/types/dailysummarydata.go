/*
# Name: DailySummaryData - Struct to hold daily summary data
*/
package types

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
