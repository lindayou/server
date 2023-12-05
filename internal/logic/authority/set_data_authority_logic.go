package authority

import (
	"context"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDataAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetDataAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDataAuthorityLogic {
	return &SetDataAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetDataAuthorityLogic) SetDataAuthority(req *types.SetDataAuthorityReq) (resp *types.SetDataAuthorityResp, err error) {
	// todo: add your logic here and delete this line

	return
}
