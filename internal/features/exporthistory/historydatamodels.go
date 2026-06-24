package exporthistory

// Inverter represents an inverter in the history data request
type Inverter struct {
	SN           string `json:"sn"`
	Name         string `json:"name"`
	ChangeNum    int    `json:"change_num"`
	ChangeType   int    `json:"change_type"`
	RelationSN   any    `json:"relation_sn"`
	RelationName any    `json:"relation_name"`
	Status       int    `json:"status"`
}

// PwsHistory represents a powerstation history entry
type PwsHistory struct {
	ID        string     `json:"id"`
	PwName    string     `json:"pw_name"`
	Status    int        `json:"status"`
	PwAddress string     `json:"pw_address"`
	Inverters []Inverter `json:"inverters"`
}

// Target represents a data target metric to retrieve
type Target struct {
	TargetKey   string `json:"target_key"`
	TargetIndex int    `json:"target_index"`
}

// HistoryDataRequest is the request payload for the ExportExcelStationHistoryData API
type HistoryDataRequest struct {
	DataType     int          `json:"data_type"`
	TimesType    int          `json:"times_type"`
	QryTimeStart string       `json:"qry_time_start"`
	QryTimeEnd   string       `json:"qry_time_end"`
	Times        int          `json:"times"`
	QryStatus    int          `json:"qry_status"`
	PwsHistorys  []PwsHistory `json:"pws_historys"`
	Targets      []Target     `json:"targets"`
}

// HistoryDataResponse is the response from the ExportExcelStationHistoryData API
type HistoryDataResponse struct {
	HasError        bool   `json:"hasError"`
	Code            any    `json:"code"`
	Msg             string `json:"msg"`
	Data            any    `json:"data"`
	TranslationCode string `json:"translationCode"`
	Components      struct {
		Para         any    `json:"para"`
		LangVer      int    `json:"langVer"`
		TimeSpan     int    `json:"timeSpan"`
		MsgSocketAdr string `json:"msgSocketAdr"`
	} `json:"components"`
}
