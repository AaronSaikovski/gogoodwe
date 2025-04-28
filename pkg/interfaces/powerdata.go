package interfaces

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

type PowerData interface {
	GetPowerData(*auth.LoginInfo) error
	GetMonitorData(*auth.LoginInfo, interface{}) error
}
