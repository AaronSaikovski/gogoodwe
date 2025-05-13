package interfaces

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
)

type SemsLogin interface {
	SemsLogin() (*auth.SemsLoginResponse, error)
}
