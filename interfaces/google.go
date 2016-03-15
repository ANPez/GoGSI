package interfaces

import (
	"github.com/anpez/gogsi/types"
)

type Google interface {
	VerifyToken(string) (*types.User, error)
}
