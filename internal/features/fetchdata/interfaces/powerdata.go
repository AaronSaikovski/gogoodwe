package interfaces

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
)

type PowerData interface {
	GetPowerData(context.Context, *auth.LoginInfo) error
}
