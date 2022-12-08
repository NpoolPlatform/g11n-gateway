package api

import (
	"context"

	g11n "github.com/NpoolPlatform/message/npool/g11n/gw/v1"

	"github.com/NpoolPlatform/g11n-gateway/api/appcountry"
	"github.com/NpoolPlatform/g11n-gateway/api/applang"
	"github.com/NpoolPlatform/g11n-gateway/api/country"
	"github.com/NpoolPlatform/g11n-gateway/api/lang"
	"github.com/NpoolPlatform/g11n-gateway/api/message"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	g11n.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	g11n.RegisterGatewayServer(server, &Server{})
	lang.Register(server)
	applang.Register(server)
	appcountry.Register(server)
	country.Register(server)
	message.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := g11n.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := lang.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := applang.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appcountry.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := country.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := message.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
