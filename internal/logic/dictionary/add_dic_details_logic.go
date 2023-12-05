package dictionary

import (
	"context"
	"server/model/admin_dic_detail"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDicDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDicDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDicDetailsLogic {
	return &AddDicDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDicDetailsLogic) AddDicDetails(req *types.AddDicDetailsReq) (resp *types.AddDicDetailsResp, err error) {
	resp = new(types.AddDicDetailsResp)
	AddDictionaryDetail := &admin_dic_detail.SysDictionaryDetails{
		Label:           req.Lable,
		Value:           req.Value,
		Status:          req.Status,
		Sort:            req.Sort,
		SysDictionaryId: req.SysDictionatyId,
	}
	insert, err := l.svcCtx.DicDetail.Insert(l.ctx, AddDictionaryDetail)
	resp.Msg = "success"

	resp.Id, err = insert.LastInsertId()

	return resp, nil
}
