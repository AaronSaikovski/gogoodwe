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

package inverterdetail

type InverterAllPoint struct {
	Language string `json:"language"`
	Function any    `json:"function"`
	HasError bool   `json:"hasError"`
	Msg      string `json:"msg"`
	Code     string `json:"code"`
	Data     struct {
		Count          int `json:"count"`
		InverterPoints []struct {
			Sn   string `json:"sn"`
			Dict struct {
				Left []struct {
					IsHT         bool   `json:"isHT"`
					IsStoreSkip  bool   `json:"isStoreSkip"`
					Key          string `json:"key"`
					Value        string `json:"value"`
					Unit         string `json:"unit"`
					IsFaultMsg   int    `json:"isFaultMsg"`
					FaultMsgCode int    `json:"faultMsgCode"`
				} `json:"left"`
				Right []struct {
					IsHT         bool   `json:"isHT"`
					IsStoreSkip  bool   `json:"isStoreSkip"`
					Key          string `json:"key"`
					Value        string `json:"value"`
					Unit         string `json:"unit"`
					IsFaultMsg   int    `json:"isFaultMsg"`
					FaultMsgCode int    `json:"faultMsgCode"`
				} `json:"right"`
			} `json:"dict"`
			Points []struct {
				TargetIndex   int    `json:"target_index"`
				TargetName    string `json:"target_name"`
				Display       string `json:"display"`
				Unit          string `json:"unit"`
				TargetKey     string `json:"target_key"`
				TextCn        string `json:"text_cn"`
				TargetSnSix   any    `json:"target_sn_six"`
				TargetSnSeven any    `json:"target_sn_seven"`
				TargetType    any    `json:"target_type"`
				StorageName   any    `json:"storage_name"`
			} `json:"points"`
			IsStored        bool    `json:"is_stored"`
			Name            string  `json:"name"`
			InPac           float64 `json:"in_pac"`
			OutPac          float64 `json:"out_pac"`
			Eday            float64 `json:"eday"`
			Emonth          float64 `json:"emonth"`
			Etotal          float64 `json:"etotal"`
			Status          int     `json:"status"`
			Soc             string  `json:"soc"`
			HTotal          float64 `json:"hTotal"`
			LastRefreshTime string  `json:"last_refresh_time"`
			Vbattery1       float64 `json:"vbattery1"`
			Ibattery1       float64 `json:"ibattery1"`
			Master          int     `json:"master"`
			IsShowOutput    bool    `json:"is_showOutput"`
			LocalDate       string  `json:"local_date"`
		} `json:"inverterPoints"`
	} `json:"data"`
	Components struct {
		Para         string `json:"para"`
		LangVer      int    `json:"langVer"`
		TimeSpan     int    `json:"timeSpan"`
		API          string `json:"api"`
		MsgSocketAdr any    `json:"msgSocketAdr"`
	} `json:"components"`
}

func NewInverterAllPoint() *InverterAllPoint {
	return &InverterAllPoint{}
}
