package types

// SemsLoginCreds - Struct to hold User login data
type SemsLoginCreds struct {
	Account        string `json:"account"`
	Password       string `json:"pwd"`
	PowerStationID string `json:"powerstationid"`
}
