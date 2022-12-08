package lang

import (
	"context"

	langmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"
)

func CreateLang(ctx context.Context, in *langmgrpb.LangReq) (*langmgrpb.Lang, error) {
	return langmgrcli.CreateLang(ctx, in)
}

func CreateLangs(ctx context.Context, in []*langmgrpb.LangReq) ([]*langmgrpb.Lang, error) {
	return langmgrcli.CreateLangs(ctx, in)
}
