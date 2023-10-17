package appcountry

import (
	"context"

	appcountrymwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/appcountry"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"
	appcountrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
)

func (h *Handler) DeleteCountry(ctx context.Context) (*npool.Country, error) {
	exist, err := appcountrymwcli.ExistAppCountryConds(ctx, &appcountrymwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}

	info, err := appcountrymwcli.DeleteCountry(ctx, &appcountrymwpb.CountryReq{ID: h.ID})
	if err != nil {
		return nil, err
	}

	outs, err := expand(ctx, []*appcountrymwpb.Country{info})
	if err != nil {
		return nil, err
	}

	return outs[0], nil
}
