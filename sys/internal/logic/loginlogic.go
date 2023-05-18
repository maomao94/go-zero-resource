package logic

import (
	"context"
	"fmt"
	"github.com/hehanpeng/go-zero-resource/common/ctxdata"
	"github.com/hehanpeng/go-zero-resource/common/errorx"
	"github.com/hehanpeng/go-zero-resource/common/tool"
	"github.com/hehanpeng/go-zero-resource/sys/internal/svc"
	"github.com/hehanpeng/go-zero-resource/sys/pb"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mapping"
	"net/http"
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
	case "sso":
		data, err := l.loginBySso(in.AuthKey, in.Password)
		if err != nil {
			return nil, err
		}
		return &pb.LoginResp{
			AccessToken: data.TokenInfo.TokenValue,
		}, nil
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
		return 0, errorx.NewEnumErrorf(errorx.Code_ErrParam, fmt.Sprintf("mobile:%s", mobile))
	}
	return 1, nil
}

func (l *LoginLogic) loginBySso(mobile, password string) (*ctxdata.SsoLoginResp, error) {
	type Data struct {
		Name string `form:"name"`
		Pwd  string `form:"pwd"`
	}
	var data = Data{
		Name: mobile,
		Pwd:  password,
	}
	resp, err := l.svcCtx.SsoSvc.Do(l.ctx, http.MethodPost, l.svcCtx.Config.SsoUrl.Login, data)
	if err != nil {
		return nil, err
	}
	var val ctxdata.SsoLoginResp
	err = mapping.UnmarshalJsonReader(resp.Body, &val)
	if err != nil {
		return nil, err
	}
	if val.Code != 200 {
		return nil, errorx.NewEnumErrorf(errorx.Code_ErrLogin, fmt.Sprintf("sso:%s", val.Msg))
	}
	return &val, nil
}
