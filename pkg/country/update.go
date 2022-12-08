package country

import (
	"context"

	countrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/country"
	countrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"
)

func UpdateCountry(ctx context.Context, in *countrymgrpb.CountryReq) (*countrymgrpb.Country, error) {
	return countrymgrcli.UpdateCountry(ctx, in)
}
