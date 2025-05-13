package powerflow

type Powerflow struct {
	Language string `json:"language,omitempty"`
	Function any    `json:"function,omitempty"`
	HasError bool   `json:"hasError,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Code     string `json:"code,omitempty"`
	Data     struct {
		HasGenset                bool `json:"hasGenset,omitempty"`
		HasMoreInverter          bool `json:"hasMoreInverter,omitempty"`
		HasPowerflow             bool `json:"hasPowerflow,omitempty"`
		Powerflow                any  `json:"powerflow,omitempty"`
		HasGridLoad              bool `json:"hasGridLoad,omitempty"`
		IsStored                 bool `json:"isStored,omitempty"`
		IsParallelInventers      bool `json:"isParallelInventers,omitempty"`
		IsMixedParallelInventers bool `json:"isMixedParallelInventers,omitempty"`
		IsEvCharge               bool `json:"isEvCharge,omitempty"`
		EvCharge                 any  `json:"evCharge,omitempty"`
	} `json:"data"`
	Components struct {
		Para         string `json:"para,omitempty"`
		LangVer      int    `json:"langVer,omitempty"`
		TimeSpan     int    `json:"timeSpan,omitempty"`
		API          string `json:"api,omitempty"`
		MsgSocketAdr any    `json:"msgSocketAdr,omitempty"`
	} `json:"components"`
}

func NewPowerflow() *Powerflow {
	return &Powerflow{}
}
