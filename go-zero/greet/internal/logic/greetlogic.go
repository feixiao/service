package logic

import (
	"context"

	"github.com/feixiao/learning/go-zero/greet/internal/svc"
	"github.com/feixiao/learning/go-zero/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{Message: "just test"}, nil
}
