

package di

import (
	"demo/internal/dao"
	"demo/internal/service"
	"demo/internal/server/grpc"

	"github.com/google/wire"
)


func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider,  grpc.New, NewApp))
}
