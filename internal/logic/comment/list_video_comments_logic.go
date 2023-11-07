package comment

import (
	"context"
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/comment"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVideoCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoCommentsLogic {
	return &ListVideoCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVideoCommentsLogic) ListVideoComments(req *types.ListVideoCommentsReq) (resp *types.ListVideoCommentsResp, err error) {
	resp = &types.ListVideoCommentsResp{}
	doerId := jwt.GetUidFromCtx(l.ctx)
	opts := &comment_model.CommentOption{Paginator: &orm.ListOptions{}}
	err = copier.Copy(&opts, req)
	if err != nil {
		return nil, err
	}
	comments, err := l.svcCtx.CommentModel.FindParentCommentByVideoId(l.ctx, util.MustString2Int64(req.VideoId), opts)
	resp.Comments, err = NewConvert(l.ctx, l.svcCtx).BuildParentCommentList(l.ctx, doerId, comments)
	return
}
