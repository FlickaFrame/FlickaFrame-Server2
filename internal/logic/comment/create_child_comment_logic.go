package comment

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChildCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChildCommentLogic {
	return &CreateChildCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateChildCommentLogic) CreateChildComment(req *types.CreateChildCommentReq) (resp *types.CreateChildCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
