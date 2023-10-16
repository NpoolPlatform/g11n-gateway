package appcountry

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"
	appcountrymwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/appcountry"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
)

func expand(ctx context.Context, infos []*appcountrymwpb.Country) ([]*npool.Country, error) {
	if len(infos) == 0 {
		return nil, nil
	}

	app, err := appmwcli.GetApp(ctx, infos[0].AppID)
	if err != nil {
		return nil, err
	}
	if app == nil {
		return nil, fmt.Errorf("invalid app")
	}

	outs := []*npool.Country{}
	for _, info := range infos {
		outs = append(outs, &npool.Country{
			ID:        info.ID,
			EntID:     info.EntID,
			AppID:     info.AppID,
			AppName:   app.Name,
			CountryID: info.CountryID,
			Country:   info.Country,
			Flag:      info.Flag,
			Code:      info.Code,
			Short:     info.Short,
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		})
	}

	return outs, nil
}
