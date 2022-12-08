package country

import (
	"context"

	country1 "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	country1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	country1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := country1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
