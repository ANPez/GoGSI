package interfaces

import (
	"github.com/ANPez/gogsi/types"
)

type Google interface {
	VerifyToken(string) (*types.User, error)
}
