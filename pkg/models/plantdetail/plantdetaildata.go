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

package plantdetail

type PlantDetailByPowerstationId struct {
	Language string `json:"language"`
	Function any    `json:"function"`
	HasError bool   `json:"hasError"`
	Msg      string `json:"msg"`
	Code     string `json:"code"`
	Data     struct {
		Info struct {
			PowerstationID   string  `json:"powerstation_id"`
			Time             string  `json:"time"`
			DateFormatYm     string  `json:"date_format_ym"`
			Stationname      string  `json:"stationname"`
			Address          string  `json:"address"`
			BatteryCapacity  float64 `json:"battery_capacity"`
			CreateTime       string  `json:"create_time"`
			Capacity         float64 `json:"capacity"`
			PowerstationType string  `json:"powerstation_type"`
			Status           int     `json:"status"`
			IsStored         bool    `json:"is_stored"`
			OnlyBps          bool    `json:"only_bps"`
			OnlyBpu          bool    `json:"only_bpu"`
			TimeSpan         float64 `json:"time_span"`
			OrgCode          string  `json:"org_code"`
			OrgName          string  `json:"org_name"`
			LocalDate        string  `json:"local_date"`
		} `json:"info"`
		Kpi struct {
			MonthGeneration float64 `json:"month_generation"`
			Pac             float64 `json:"pac"`
			Power           float64 `json:"power"`
			TotalPower      float64 `json:"total_power"`
			DayIncome       float64 `json:"day_income"`
			TotalIncome     float64 `json:"total_income"`
			YieldRate       float64 `json:"yield_rate"`
			Currency        string  `json:"currency"`
		} `json:"kpi"`
		IsEvcharge         bool   `json:"isEvcharge"`
		IsTigo             bool   `json:"isTigo"`
		IsPowerflow        bool   `json:"isPowerflow"`
		IsSec              bool   `json:"isSec"`
		IsEnvironmental    bool   `json:"isEnvironmental"`
		IsGenset           bool   `json:"isGenset"`
		IsMicroInverter    bool   `json:"isMicroInverter"`
		HasLayout          bool   `json:"hasLayout"`
		LayoutID           string `json:"layout_id"`
		PowercontrolStatus int    `json:"powercontrol_status"`
		ChartsTypesByPlant []struct {
			Date         string `json:"date"`
			TypeName     string `json:"typeName"`
			ChartIndices []struct {
				IndexName    string `json:"indexName"`
				IndexLabel   string `json:"indexLabel"`
				ChartIndexID string `json:"chartIndexId"`
				DateRange    []struct {
					Text         string `json:"text"`
					Value        string `json:"value"`
					Type         string `json:"type"`
					Now          string `json:"now"`
					DateFormater any    `json:"dateFormater"`
				} `json:"dateRange"`
			} `json:"chartIndices"`
		} `json:"chartsTypesByPlant"`
		Soc         []any `json:"soc"`
		IndustrySoc []any `json:"industrySoc"`
	} `json:"data"`
	Components struct {
		Para         string `json:"para"`
		LangVer      int    `json:"langVer"`
		TimeSpan     int    `json:"timeSpan"`
		API          string `json:"api"`
		MsgSocketAdr any    `json:"msgSocketAdr"`
	} `json:"components"`
}

func NewGetPlantDetailByPowerstationId() *PlantDetailByPowerstationId {
	return &PlantDetailByPowerstationId{}

}
