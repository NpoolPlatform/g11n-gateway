package appcountry

import (
	"context"

	appcountrymwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"
	appcountrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func (h *Handler) GetCountries(ctx context.Context) ([]*npool.Country, uint32, error) {
	infos, total, err := appcountrymwcli.GetCountries(ctx, &appcountrymwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}},
		h.Offset,
		h.Limit,
	)
	if err != nil {
		return nil, 0, err
	}

	outs, err := expand(ctx, infos)
	if err != nil {
		return nil, 0, err
	}

	return outs, total, nil
}
