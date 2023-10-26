package applang

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"
	applangmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
)

func expand(ctx context.Context, infos []*applangmwpb.Lang) ([]*npool.Lang, error) {
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

	outs := []*npool.Lang{}
	for _, info := range infos {
		outs = append(outs, &npool.Lang{
			ID:        info.ID,
			EntID:     info.EntID,
			AppID:     info.AppID,
			AppName:   app.Name,
			LangID:    info.LangID,
			Lang:      info.Lang,
			Logo:      info.Logo,
			Name:      info.Name,
			Short:     info.Short,
			Main:      info.Main,
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		})
	}

	return outs, nil
}
