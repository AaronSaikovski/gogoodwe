package interfaces

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

type PowerData interface {
	GetPowerData(context.Context, *auth.LoginInfo) error
}
