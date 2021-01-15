// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"fgstgu/go/src/c/kratos/hello-world/internal/dao"
	"fgstgu/go/src/c/kratos/hello-world/internal/service"
	"fgstgu/go/src/c/kratos/hello-world/internal/server/grpc"
	"fgstgu/go/src/c/kratos/hello-world/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
