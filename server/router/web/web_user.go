package web

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WebUserRouter struct {
}

// InitHomePageRouter 初始化 HomePage 路由信息
func (s *WebUserRouter) InitWebUserRouter(Router *gin.RouterGroup) {
	// homePageRouter := Router.Group("homepage").Use(middleware.OperationRecord())
	webUserouterWithoutRecord := Router.Group("user")
	var WebUserApi = v1.ApiGroupApp.WebApiGroup.WebUserApi
	{
		webUserouterWithoutRecord.GET(":id", WebUserApi.GetUserInfo)                  // 根据ID获取用户首页信息
		webUserouterWithoutRecord.GET(":id/momentList", WebUserApi.GetMomentList)     // 根据id获取用户主页动态
		webUserouterWithoutRecord.GET(":id/activityList", WebUserApi.GetActivityList) // 根据id获取用户参加活动
		webUserouterWithoutRecord.GET(":id/commentList", WebUserApi.GetCommentList)   // 工具id获取用户的评论
	}
}
