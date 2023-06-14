package country

import (
	"context"

	countrymwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/country"
	countrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
)

func (h *Handler) CreateCountry(ctx context.Context) (*countrymwpb.Country, error) {
	return countrymwcli.CreateCountry(ctx, &countrymwpb.CountryReq{
		ID:      h.ID,
		Country: h.Country,
		Flag:    h.Flag,
		Code:    h.Code,
		Short:   h.Short,
	})
}

func (h *Handler) CreateCountries(ctx context.Context) ([]*countrymwpb.Country, error) {
	return countrymwcli.CreateCountries(ctx, h.Reqs)
}
