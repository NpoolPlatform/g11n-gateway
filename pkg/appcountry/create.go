package appcountry

import (
	"context"

	appcountrymwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"
	appcountrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"
)

func (h *Handler) CreateCountry(ctx context.Context) (*npool.Country, error) {
	info, err := appcountrymwcli.CreateCountry(ctx, &appcountrymwpb.CountryReq{
		ID:        h.ID,
		AppID:     h.AppID,
		CountryID: h.CountryID,
	})
	if err != nil {
		return nil, err
	}

	outs, err := expand(ctx, []*appcountrymwpb.Country{info})
	if err != nil {
		return nil, err
	}

	return outs[0], nil
}
