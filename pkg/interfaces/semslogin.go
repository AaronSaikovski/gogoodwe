package interfaces

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

type SemsLogin interface {
	SemsLogin(ctx context.Context) (*auth.SemsLoginResponse, error)
}
