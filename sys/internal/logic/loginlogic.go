package logic

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/common/errorx"
	"github.com/hehanpeng/go-zero-resource/common/tool"
	"github.com/hehanpeng/go-zero-resource/sys/internal/svc"
	"github.com/hehanpeng/go-zero-resource/sys/pb"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var err error
	var userId int64
	switch in.AuthType {
	case "system":
		userId, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		return nil, errors.New("AuthType error.")
	}
	if err != nil {
		return nil, err
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	generateTokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		AccessToken:  generateTokenResp.AccessToken,
		AccessExpire: generateTokenResp.AccessExpire,
		RefreshAfter: generateTokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
	//user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	//if err != nil && err != model.ErrNotFound {
	//	return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	//}
	//if user == nil {
	//	return 0, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
	//}
	if !(tool.Md5ByString(password) == "e10adc3949ba59abbe56e057f20f883e") {
		return 0, errorx.NewEnumErrorf(errorx.ErrCode_Y_1002001, mobile)
	}
	return 1, nil
}
