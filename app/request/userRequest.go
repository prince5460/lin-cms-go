package request

import (
	"github.com/gogf/gf/util/gvalid"
	"lin-cms-go/core"

	"github.com/gogf/gf/net/ghttp"
)

type UserLoginRequest struct {
	Username string `v:"required#账号不能为空" p:"username"`
	Password string `v:"required#密码不能为空" p:"password"`
}
type UserRegRequest struct {
	Username        string `v:"required#账号不能为空" p:"username"`
	Password        string `v:"required#密码不能为空" p:"password"`
	GroupIds        []int  `v:"required#GroupIds不能为空"p:"group_ids"`
	ConfirmPassword string `v:"required#ConfirmPassword不能为空"p:"confirm_password"`
}

func check(err error, r *ghttp.Request) {
	if v, ok := err.(*gvalid.Error); ok {
		core.FailResp(r, 1, v.FirstString())
	}
	core.FailResp(r, 1, err.Error())
}
func Login(r *ghttp.Request) *UserLoginRequest {
	var data UserLoginRequest
	if err := r.Parse(&data); err != nil {
		check(err, r)
	}
	return &data
}
func Register(r *ghttp.Request) *UserRegRequest {
	var data UserRegRequest
	if err := r.Parse(&data); err != nil {
		check(err, r)
	}
	return &data
}
