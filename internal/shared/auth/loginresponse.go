package auth

// LoginResponse - SEMS API Response struct
type SemsLoginResponse struct {
	HasError bool   `json:"hasError"`
	Code     int32  `json:"code"`
	Msg      string `json:"msg"`
	Data     struct {
		UID       string `json:"uid"`
		Timestamp int64  `json:"timestamp"`
		Token     string `json:"token"`
		Client    string `json:"client"`
		Version   string `json:"version"`
		Language  string `json:"language"`
	} `json:"data"`
	Components struct {
		Para         any    `json:"para"`
		LangVer      int    `json:"langVer"`
		TimeSpan     int    `json:"timeSpan"`
		API          string `json:"api"`
		MsgSocketAdr string `json:"msgSocketAdr"`
	} `json:"components"`
	API string `json:"api"`
}
