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
		//短信验证码
		vcc := new(auth.VerifyCodeController)
		v1.POST("/sms/verify-code", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)

		oauthGroup := v1.Group("/oauth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		oauthGroup.Use(middlewares.LimitIP("1000-H"))
		{
			lgc := new(auth.LoginController)
			//短信登录
			oauthGroup.POST("/sms", middlewares.GuestJWT(), lgc.LoginByPhone)
			oauthGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			oauthGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
			oauthGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)

			// 注册用户
			suc := new(auth.SignupController)
			oauthGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			oauthGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)
			oauthGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			oauthGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			oauthGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
		}

		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)

		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middlewares.AuthJWT(), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
			usersGroup.PUT("/phone", middlewares.AuthJWT(), uc.UpdatePhone)
			usersGroup.PUT("/password", middlewares.AuthJWT(), uc.UpdatePassword)
			usersGroup.PUT("/avatar", middlewares.AuthJWT(), uc.UpdateAvatar)
		}

		cgcGroup := v1.Group("/categories")
		{
			cgc := new(controllers.CategoriesController)
			cgcGroup.GET("", cgc.Index)
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		tpcGroup := v1.Group("/topics")
		{
			tpc := new(controllers.TopicsController)
			tpcGroup.GET("", tpc.Index)
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
			tpcGroup.GET("/:id", tpc.Show)
		}

		linksGroup := v1.Group("/links")
		{
			lsc := new(controllers.LinksController)
			linksGroup.GET("", lsc.Index)
		}
	}
}
