package mw

import (
	"context"
	"dousheng/biz/dal/mysql"
	"dousheng/pkg/constants"
	utils2 "dousheng/pkg/utils"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

// InitJwt 初始化 jwt 中间件
func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		// 用于设置所属领域名称，默认为 hertz jwt
		Realm: "dousheng jwt",
		// 用于设置签名密钥（必要配置）
		Key: []byte(constants.SecretKey),
		// 用于设置 token 过期时间，默认为一小时
		Timeout: 24 * 30 * time.Hour,
		// 用于设置最大 token 刷新时间，允许客户端在 TokenTime + MaxRefresh 内
		// 刷新 token 的有效时间，追加一个 Timeout 的时长
		MaxRefresh: time.Hour,
		// 用于设置 token 的获取源，可以选择 header、query、cookie、param、form，默认为 header:Authorization
		TokenLookup: "query: token",
		// 用于设置从 header 中获取 token 时的前缀，默认为 Bearer
		TokenHeadName: "Bearer",
		// 用于设置登录的响应函数
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			userId, _ := c.Get("user_id")
			c.JSON(http.StatusOK, utils.H{
				"status_code":    code,
				"status_message": "success",
				"user_id":        userId,
				"token":          token,
				//"expire":         expire.Format(time.RFC3339),
			})
		},
		// 用于设置登录时认证用户信息的函数（必要配置）
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := mysql.CheckUser(loginStruct.Username, utils2.MD5(loginStruct.Password))
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, errors.New("user doesn't exist or wrong password")
			}
			return users[0], nil
		},
		// 用于设置检索身份的键，默认为 identity
		IdentityKey: IdentityKey,
		// 用于设置获取身份信息的函数，默认与 IdentityKey 配合使用
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &mysql.UserRaw{
				Model: gorm.Model{
					ID: uint(claims[IdentityKey].(float64)),
				},
			}
		},
		// 用于设置登陆成功后为向 token 中添加自定义负载信息的函数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*mysql.UserRaw); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		// 用于设置 jwt 校验流程发生错误时响应所包含的错误信息
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		// 用于设置 jwt 验证流程失败的响应函数
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code":    code,
				"status_message": message,
				"token":          "",
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
