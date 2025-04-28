package plantdetail

type PlantDetailByPowerstationId struct {
	Language string `json:"language,omitempty"`
	Function any    `json:"function,omitempty"`
	HasError bool   `json:"hasError,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Code     string `json:"code,omitempty"`
	Data     struct {
		Info struct {
			PowerstationID   string  `json:"powerstation_id,omitempty"`
			Time             string  `json:"time,omitempty"`
			DateFormatYm     string  `json:"date_format_ym,omitempty"`
			Stationname      string  `json:"stationname,omitempty"`
			Address          string  `json:"address,omitempty"`
			BatteryCapacity  float64 `json:"battery_capacity,omitempty"`
			CreateTime       string  `json:"create_time,omitempty"`
			Capacity         float64 `json:"capacity,omitempty"`
			PowerstationType string  `json:"powerstation_type,omitempty"`
			Status           int     `json:"status,omitempty"`
			IsStored         bool    `json:"is_stored,omitempty"`
			OnlyBps          bool    `json:"only_bps,omitempty"`
			OnlyBpu          bool    `json:"only_bpu,omitempty"`
			TimeSpan         float64 `json:"time_span,omitempty"`
			OrgCode          string  `json:"org_code,omitempty"`
			OrgName          string  `json:"org_name,omitempty"`
			LocalDate        string  `json:"local_date,omitempty"`
		} `json:"info,omitempty"`
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
		IsEvcharge         bool   `json:"isEvcharge,omitempty"`
		IsTigo             bool   `json:"isTigo,omitempty"`
		IsPowerflow        bool   `json:"isPowerflow,omitempty"`
		IsSec              bool   `json:"isSec,omitempty"`
		IsEnvironmental    bool   `json:"isEnvironmental,omitempty"`
		IsGenset           bool   `json:"isGenset,omitempty"`
		IsMicroInverter    bool   `json:"isMicroInverter,omitempty"`
		HasLayout          bool   `json:"hasLayout,omitempty"`
		LayoutID           string `json:"layout_id,omitempty"`
		PowercontrolStatus int    `json:"powercontrol_status,omitempty"`
		ChartsTypesByPlant []struct {
			Date         string `json:"date,omitempty"`
			TypeName     string `json:"typeName,omitempty"`
			ChartIndices []struct {
				IndexName    string `json:"indexName,omitempty"`
				IndexLabel   string `json:"indexLabel,omitempty"`
				ChartIndexID string `json:"chartIndexId,omitempty"`
				DateRange    []struct {
					Text         string `json:"text,omitempty"`
					Value        string `json:"value,omitempty"`
					Type         string `json:"type,omitempty"`
					Now          string `json:"now,omitempty"`
					DateFormater any    `json:"dateFormater,omitempty"`
				} `json:"dateRange,omitempty"`
			} `json:"chartIndices,omitempty"`
		} `json:"chartsTypesByPlant,omitempty"`
		Soc         []any `json:"soc,omitempty"`
		IndustrySoc []any `json:"industrySoc,omitempty"`
	} `json:"data,omitempty"`
	Components struct {
		Para         string `json:"para,omitempty"`
		LangVer      int    `json:"langVer,omitempty"`
		TimeSpan     int    `json:"timeSpan,omitempty"`
		API          string `json:"api,omitempty"`
		MsgSocketAdr any    `json:"msgSocketAdr,omitempty"`
	} `json:"components,omitempty"`
}

func NewGetPlantDetailByPowerstationId() *PlantDetailByPowerstationId {
	return &PlantDetailByPowerstationId{}

}
