package plantchartdata

import "time"

type PlantPowerChart struct {
	Language string `json:"language,omitempty"`
	Function any    `json:"function,omitempty"`
	HasError bool   `json:"hasError,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Code     string `json:"code,omitempty"`
	Data     struct {
		GenerateData []struct {
			Key     string  `json:"key,omitempty"`
			Value   float64 `json:"value,omitempty"`
			UnitKey string  `json:"unit_Key,omitempty"`
		} `json:"generateData,omitempty"`
		Lines []struct {
			Key        string `json:"key,omitempty"`
			Unit       string `json:"unit,omitempty"`
			FrontColor string `json:"frontColor,omitempty"`
			IsActive   bool   `json:"isActive,omitempty"`
			Axis       int    `json:"axis,omitempty"`
			Sort       int    `json:"sort,omitempty"`
			Type       string `json:"type,omitempty"`
			Xy         []struct {
				X string  `json:"x,omitempty"`
				Y float64 `json:"y,omitempty"`
				Z any     `json:"z,omitempty"`
			} `json:"xy"`
		} `json:"lines,omitempty"`
	} `json:"data,omitempty"`
	Components struct {
		Para         time.Time `json:"para,omitempty"`
		LangVer      int       `json:"langVer,omitempty"`
		TimeSpan     int       `json:"timeSpan,omitempty"`
		API          string    `json:"api,omitempty"`
		MsgSocketAdr any       `json:"msgSocketAdr,omitempty"`
	} `json:"components,omitempty"`
}

func NewPlantPowerChart() *PlantPowerChart {
	return &PlantPowerChart{}
}
