package message

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func Expand(ctx context.Context, infos []*messagemwpb.Message) ([]*npool.Message, error) {
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

	outs := []*npool.Message{}
	for _, info := range infos {
		outs = append(outs, &npool.Message{
			ID:        info.ID,
			EntID:     info.EntID,
			AppName:   app.Name,
			LangID:    info.LangID,
			Lang:      info.Lang,
			MessageID: info.MessageID,
			Message:   info.Message,
			GetIndex:  info.GetIndex,
			Disabled:  info.Disabled,
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		})
	}

	return outs, nil
}
