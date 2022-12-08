package country

import (
	"context"

	countrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/country"
	countrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"
)

func GetCountries(ctx context.Context, in *countrymgrpb.Conds, offset, limit int32) ([]*countrymgrpb.Country, uint32, error) {
	return countrymgrcli.GetCountries(ctx, in, offset, limit)
}
