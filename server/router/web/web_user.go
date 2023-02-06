package web

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WebUserRouter struct {
}

// InitHomePageRouter 初始化 HomePage 路由信息
func (s *WebUserRouter) InitWebUserRouter(Router *gin.RouterGroup, AuthRouter *gin.RouterGroup) {
	webUserouterWithAuth := AuthRouter.Group("user")
	webUserouter := Router.Group("user")
	var WebUserApi = v1.ApiGroupApp.WebApiGroup.WebUserApi
	{
		webUserouter.GET(":id", WebUserApi.GetUserInfo)                  // 根据ID获取用户首页信息
		webUserouter.GET(":id/momentList", WebUserApi.GetMomentList)     // 根据id获取用户主页动态
		webUserouter.GET(":id/activityList", WebUserApi.GetActivityList) // 根据id获取用户参加活动
		webUserouter.GET(":id/commentList", WebUserApi.GetCommentList)   // 工具id获取用户的评论
	}
	{
		webUserouterWithAuth.POST("createActivity", WebUserApi.CreateActivity)   // 发布活动
		webUserouterWithAuth.POST("createArticle", WebUserApi.CreateArticle)     // 发布文章
		webUserouterWithAuth.POST("creatComment", WebUserApi.CreatComment)       // 发布评论
		webUserouterWithAuth.POST("createIntroduce", WebUserApi.CreateIntroduce) // 发布内推

		webUserouterWithAuth.POST("attendActivity", WebUserApi.AttendActivity)       // 参与活动
		webUserouterWithAuth.POST("createInformation", WebUserApi.CreateInformation) // 填写简历/添加社团信息
		webUserouterWithAuth.POST("voteComment", WebUserApi.VoteComment)             // 点赞/点踩评论
		webUserouterWithAuth.POST("collection", WebUserApi.Collection)               // 收藏社团/文章
	}
}
