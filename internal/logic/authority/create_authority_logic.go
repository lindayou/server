package authority

import (
	"context"
	"server/model/anthority_model"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAuthorityLogic {
	return &CreateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAuthorityLogic) CreateAuthority(req *types.CreateAuthorityReq) (resp *types.CreateAuthorityResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.CreateAuthorityResp)
	authority := anthority_model.SysAuthorities{
		AuthorityId:   req.AuthorityId,
		AuthorityName: req.AuthorityName,
		ParentId:      req.ParentId,
	}
	_, err = l.svcCtx.Authority.Insert(l.ctx, &authority)
	if err != nil {
		return nil, err
	}
	resp.Message = "创建成功"

	return resp, nil
}
