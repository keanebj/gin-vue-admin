package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	{
		userRouter.POST("admin_register", baseApi.Register)           // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)     // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority) // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)           // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)            // 设置用户信息

		userRouter.POST("admin_registerV", baseApi.RegisterV)               // 管理员注册账号v
		userRouter.POST("changePasswordV", baseApi.ChangePasswordV)         // 用户修改密码v
		userRouter.POST("setUserAuthorityV", baseApi.SetUserAuthorityV)     // 设置用户权限v
		userRouter.DELETE("deleteUserV", baseApi.DeleteUserV)               // 删除用户v
		userRouter.PUT("setUserInfoV", baseApi.SetUserInfoV)                // 设置用户信息v
		userRouter.POST("setUserAuthoritiesV", baseApi.SetUserAuthoritiesV) // 设置用户权限组
		userRouter.POST("resetPasswordV", baseApi.ResetPasswordV)           // 重置用户密码

		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
	}
}
