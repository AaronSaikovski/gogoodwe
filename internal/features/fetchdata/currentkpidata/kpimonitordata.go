package currentkpidata

// Current KPI Data - Struct to hold data returned from the Inverter Powerstation API
type KPIMonitorData struct {
	Language string `json:"language,omitempty"`
	HasError bool   `json:"hasError,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Code     string `json:"code,omitempty"`
	Data     struct {
		Kpi struct {
			MonthGeneration float64 `json:"month_generation,omitempty"`
			Pac             float64 `json:"pac,omitempty"`
			Power           float64 `json:"power,omitempty"`
			TotalPower      float64 `json:"total_power,omitempty"`
			DayIncome       float64 `json:"day_income,omitempty"`
			TotalIncome     float64 `json:"total_income,omitempty"`
			YieldRate       float64 `json:"yield_rate,omitempty"`
			Currency        string  `json:"currency,omitempty"`
		} `json:"kpi,omitempty"`
	} `json:"data,omitempty"`
}

func NewKPIMonitorData() *KPIMonitorData {
	return &KPIMonitorData{}
}
