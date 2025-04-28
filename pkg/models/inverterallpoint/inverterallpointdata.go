package inverteallpoint

type InverterAllPoint struct {
	Language string `json:"language,omitempty"`
	Function any    `json:"function,omitempty"`
	HasError bool   `json:"hasError,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Code     string `json:"code,omitempty"`
	Data     struct {
		Count          int `json:"count,omitempty"`
		InverterPoints []struct {
			Sn   string `json:"sn,omitempty"`
			Dict struct {
				Left []struct {
					IsHT         bool   `json:"isHT,omitempty"`
					IsStoreSkip  bool   `json:"isStoreSkip,omitempty"`
					Key          string `json:"key,omitempty"`
					Value        string `json:"value,omitempty"`
					Unit         string `json:"unit,omitempty"`
					IsFaultMsg   int    `json:"isFaultMsg,omitempty"`
					FaultMsgCode int    `json:"faultMsgCode,omitempty"`
				} `json:"left,omitempty"`
				Right []struct {
					IsHT         bool   `json:"isHT,omitempty"`
					IsStoreSkip  bool   `json:"isStoreSkip,omitempty"`
					Key          string `json:"key,omitempty"`
					Value        string `json:"value,omitempty"`
					Unit         string `json:"unit,omitempty"`
					IsFaultMsg   int    `json:"isFaultMsg,omitempty"`
					FaultMsgCode int    `json:"faultMsgCode,omitempty"`
				} `json:"right,omitempty"`
			} `json:"dict,omitempty"`
			Points []struct {
				TargetIndex   int    `json:"target_index,omitempty"`
				TargetName    string `json:"target_name,omitempty"`
				Display       string `json:"display,omitempty"`
				Unit          string `json:"unit,omitempty"`
				TargetKey     string `json:"target_key,omitempty"`
				TextCn        string `json:"text_cn,omitempty"`
				TargetSnSix   any    `json:"target_sn_six,omitempty"`
				TargetSnSeven any    `json:"target_sn_seven,omitempty"`
				TargetType    any    `json:"target_type,omitempty"`
				StorageName   any    `json:"storage_name,omitempty"`
			} `json:"points,omitempty"`
			IsStored        bool    `json:"is_stored,omitempty"`
			Name            string  `json:"name,omitempty"`
			InPac           float64 `json:"in_pac,omitempty"`
			OutPac          float64 `json:"out_pac,omitempty"`
			Eday            float64 `json:"eday,omitempty"`
			Emonth          float64 `json:"emonth,omitempty"`
			Etotal          float64 `json:"etotal,omitempty"`
			Status          int     `json:"status,omitempty"`
			Soc             string  `json:"soc,omitempty"`
			HTotal          float64 `json:"hTotal,omitempty"`
			LastRefreshTime string  `json:"last_refresh_time,omitempty"`
			Vbattery1       float64 `json:"vbattery1,omitempty"`
			Ibattery1       float64 `json:"ibattery1,omitempty"`
			Master          int     `json:"master,omitempty"`
			IsShowOutput    bool    `json:"is_showOutput,omitempty"`
			LocalDate       string  `json:"local_date,omitempty"`
		} `json:"inverterPoints,omitempty"`
	} `json:"data,omitempty"`
	Components struct {
		Para         string `json:"para,omitempty"`
		LangVer      int    `json:"langVer,omitempty"`
		TimeSpan     int    `json:"timeSpan,omitempty"`
		API          string `json:"api,omitempty"`
		MsgSocketAdr any    `json:"msgSocketAdr,omitempty"`
	} `json:"components,omitempty"`
}

func NewInverterAllPoint() *InverterAllPoint {
	return &InverterAllPoint{}
}
