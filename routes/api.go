// Package routes 注册路由
package routes

import (
	controllers "github.com/fans1992/jiaoma/app/http/controllers/api/v1"
	"github.com/fans1992/jiaoma/app/http/controllers/api/v1/auth"
	"github.com/fans1992/jiaoma/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup

	v1 = r.Group("/api")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		vcc := new(auth.VerifyCodeController)
		//短信验证码
		v1.POST("/sms/verify-code", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)

		oauthGroup := v1.Group("/oauth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		oauthGroup.Use(middlewares.LimitIP("1000-H"))
		{
			lgc := new(auth.LoginController)
			// 短信登录
			oauthGroup.POST("/sms", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 账号密码登录
			oauthGroup.POST("/login", middlewares.GuestJWT(), lgc.LoginByPassword)

			suc := new(auth.SignupController)
			// 注册用户
			oauthGroup.POST("/signup", middlewares.GuestJWT(), suc.SignupUsingPhone)
			oauthGroup.GET("/sms/is_new", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)

			// 重置密码
			pwc := new(auth.PasswordController)
			oauthGroup.POST("/password/reset", middlewares.GuestJWT(), pwc.ResetByPhone)
		}

		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/me", middlewares.AuthJWT(), uc.CurrentUser)

		usersGroup := v1.Group("/users")
		{
			usersGroup.POST("/update/info", middlewares.AuthJWT(), uc.UpdateProfile)
			usersGroup.POST("/update/mobile", middlewares.AuthJWT(), uc.UpdatePhone)
			usersGroup.POST("/update/password", middlewares.AuthJWT(), uc.UpdatePassword)
			usersGroup.POST("/email/verificationCodes", middlewares.AuthJWT(), vcc.SendUsingEmail)
			usersGroup.POST("/update/email", middlewares.AuthJWT(), uc.UpdateEmail)
			usersGroup.POST("/avatar", middlewares.AuthJWT(), uc.UpdateAvatar)
		}

		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("", tpc.Index)
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
			tpcGroup.GET("/:id", middlewares.AuthJWT(), tpc.Show)
		}

		contact := new(controllers.ContactsController)
		contactGroup := v1.Group("/contacts")
		{
			contactGroup.POST("", middlewares.AuthJWT(), contact.Store)
		}
	}
}
