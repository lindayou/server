package user

import (
	"context"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePassLogic {
	return &ChangePassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePassLogic) ChangePass(req *types.ChangePasswordReq) (resp *types.ChangePasswordResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.ChangePasswordResp)
	err = l.svcCtx.UserModel.ChangePass(l.ctx, req.PriPassword, req.ChangePassword, req.UserId)
	if err != nil {
		return nil, err
	}
	resp.Message = "修改成功"
	return resp, nil
}
