package country

import (
	"context"

	countrymwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/country"
	countrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
)

func (h *Handler) GetCountries(ctx context.Context) ([]*countrymwpb.Country, uint32, error) {
	return countrymwcli.GetCountries(ctx, &countrymwpb.Conds{}, h.Offset, h.Limit)
}
