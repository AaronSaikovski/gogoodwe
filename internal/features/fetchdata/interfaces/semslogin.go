package interfaces

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
)

type SemsLogin interface {
	SemsLogin(ctx context.Context) (*auth.SemsLoginResponse, error)
}
