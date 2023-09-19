package message

import (
	"context"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

type queryHandler struct {
	*Handler
	infos []*messagemwpb.Message
	total uint32
}

func (h *queryHandler) formalize(ctx context.Context) ([]*npool.Message, error) {
	appIDs := []string{}
	for _, info := range h.infos {
		appIDs = append(appIDs, info.AppID)
	}

	conds := &appmwpb.Conds{
		IDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}

	apps, _, err := appmwcli.GetApps(ctx, conds, 0, int32(len(appIDs)))
	if err != nil {
		return nil, err
	}

	appMap := map[string]*appmwpb.App{}
	for _, info := range apps {
		appMap[info.ID] = info
	}

	_infos := []*npool.Message{}
	for _, info := range h.infos {
		_info := &npool.Message{
			ID:        info.ID,
			LangID:    info.LangID,
			Lang:      info.Lang,
			MessageID: info.MessageID,
			Message:   info.Message,
			GetIndex:  info.GetIndex,
			Disabled:  info.Disabled,
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		}

		dinfo, ok := appMap[info.AppID]
		if ok {
			_info.AppName = dinfo.Name
		}

		_infos = append(_infos, _info)
	}
	return _infos, nil
}

func (h *Handler) GetMessages(ctx context.Context) ([]*npool.Message, uint32, error) {
	conds := &messagemwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.LangID != nil {
		conds.LangID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.LangID}
	}
	if h.Disabled != nil {
		conds.Disabled = &basetypes.BoolVal{Op: cruder.EQ, Value: *h.Disabled}
	}
	infos, total, err := messagemwcli.GetMessages(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, err
	}

	handler := &queryHandler{
		Handler: h,
		infos:   infos,
		total:   total,
	}

	_infos, err := handler.formalize(ctx)
	if err != nil {
		return nil, 0, err
	}

	return _infos, total, nil
}
