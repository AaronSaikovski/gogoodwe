package auth

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

// ApiLoginCredentials - Struct to hold User login credentials
type SemsLoginCredentials struct {
	Account        string `json:"account"`
	Password       string `json:"pwd"`
	PowerStationID string `json:"powerstationid"`
	ID             string `json:"id"`
	Date           string `json:"date"` // YYYY-MM-DD
}

func NewSemsLoginCredentials(account, password, powerStationID string) SemsLoginCredentials {
	return SemsLoginCredentials{
		Account:        account,
		Password:       password,
		PowerStationID: powerStationID,
		ID:             powerStationID,
		Date:           utils.GetDate(),
	}

}
