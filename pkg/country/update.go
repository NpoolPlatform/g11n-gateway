package country

import (
	"context"

	countrymwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/country"
	countrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"
)

func (h *Handler) UpdateCountry(ctx context.Context) (*countrymwpb.Country, error) {
	return countrymwcli.UpdateCountry(ctx, &countrymwpb.CountryReq{
		ID:      h.ID,
		Country: h.Country,
		Flag:    h.Flag,
		Code:    h.Code,
		Short:   h.Short,
	})
}
