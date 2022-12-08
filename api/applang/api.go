package applang

import (
	"context"

	applang1 "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	applang1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	applang1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := applang1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
