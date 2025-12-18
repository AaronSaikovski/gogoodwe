package monitorsummary

// DailySummaryData - Struct to hold daily summary data
type DailySummaryData struct {
	Language string `json:"language,omitempty"`
	HasError bool   `json:"hasError,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Code     string `json:"code,omitempty"`
	Data     struct {
		Kpi struct {
			MonthGeneration float64 `json:"month_generation,omitempty"`
			Power           float64 `json:"power,omitempty"`
			TotalPower      float64 `json:"total_power,omitempty"`
			DayIncome       float64 `json:"day_income,omitempty"`
			TotalIncome     float64 `json:"total_income,omitempty"`
			Currency        string  `json:"currency,omitempty"`
		} `json:"kpi,omitempty"`
		Inverter []struct {
			TotalGeneration string `json:"total_generation,omitempty"`
			DailyGeneration string `json:"daily_generation,omitempty"`
		} `json:"inverter,omitempty"`
	} `json:"data,omitempty"`
}

func NewDailySummaryData() *DailySummaryData {
	return &DailySummaryData{}
}
