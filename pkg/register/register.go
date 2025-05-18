package register

import . "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/types"
import "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/providers"

type Registerable interface {
	Register(args RegisterArgs) error
	Name() string
}

var (
	registerProviders = []Registerable{
		&providers.InternalRegisterProvider{},
	}
)

func GetRegisterProviders() []Registerable {
	return registerProviders
}
