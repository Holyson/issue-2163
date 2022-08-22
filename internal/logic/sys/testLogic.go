package sys

import (
	"context"

	"cors-test/internal/svc"
	"cors-test/internal/types"

	"github.com/Holyson/test-go-zero-cors/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test(req *types.NothingReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
