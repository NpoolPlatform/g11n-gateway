package country

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"
	countrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	country1 "github.com/NpoolPlatform/g11n-gateway/pkg/country"
	countrymwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/country"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithCountry(&in.Country),
		country1.WithFlag(&in.Flag),
		country1.WithCode(&in.Code),
		country1.WithShort(&in.Short),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLang",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountryResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountries",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	countries := []string{}
	for _, c := range in.GetInfos() {
		countries = append(countries, c.GetCountry())
	}

	infos, _, err := countrymwcli.GetCountries(ctx, &countrymwpb.Conds{
		Countries: &basetypes.StringSliceVal{
			Op:    cruder.IN,
			Value: countries,
		},
	}, int32(0), int32(len(countries)))
	if err != nil {
		logger.Sugar().Errorw("CreateCountries", "error", err)
		return &npool.CreateCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	reqs := []*countrymwpb.CountryReq{}
	outs := []*countrymwpb.Country{}

	for _, info := range in.GetInfos() {
		exist := false
		var _info1 *countrymwpb.Country

		for _, info1 := range infos {
			if info.GetCountry() == info1.Country {
				_info1 = info1
				exist = true
				break
			}
		}
		if !exist {
			reqs = append(reqs, info)
			continue
		}

		outs = append(outs, _info1)
	}

	if len(reqs) == 0 {
		return &npool.CreateCountriesResponse{
			Infos: infos,
		}, nil
	}

	reqinfos, err := handler.CreateCountries(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountries",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	reqinfos = append(reqinfos, outs...)

	return &npool.CreateCountriesResponse{
		Infos: reqinfos,
	}, nil
}
