package web

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WebOrganizationRouter struct {
}

// InitHomePageRouter 初始化 HomePage 路由信息
func (s *WebOrganizationRouter) InitWebOrganizationRouter(Router *gin.RouterGroup, AuthRouter *gin.RouterGroup) {
	webOrganizationRouterWithAuth := AuthRouter.Group("organization")
	webOrganizationRouter := Router.Group("organization")
	var WebApi = v1.ApiGroupApp.WebApiGroup.WebOrganizationApi

	{
		webOrganizationRouter.GET(":id/homePage", WebApi.GetHomePage)       //  获取社团主页信息
		webOrganizationRouter.GET("activity/:id", WebApi.GetActivityDetail) // 获取活动信息
	}
	{
		webOrganizationRouterWithAuth.POST("activity/:id/emailAll", WebApi.SentActivityEmail) // 发布活动邮件
		// webUserouterWithAuth.POST("createArticle", WebUserApi.CreateArticle)     // 发布文章
		// webUserouterWithAuth.POST("creatComment", WebUserApi.CreatComment)       // 发布评论
		// webUserouterWithAuth.POST("createIntroduce", WebUserApi.CreateIntroduce) // 发布内推

		// webUserouterWithAuth.POST("attendActivity", WebUserApi.AttendActivity)       // 参与活动
		// webUserouterWithAuth.POST("createInformation", WebUserApi.CreateInformation) // 填写简历/添加社团信息
		// webUserouterWithAuth.POST("voteComment", WebUserApi.VoteComment)             // 点赞/点踩评论
		// webUserouterWithAuth.POST("collection", WebUserApi.Collection)               // 收藏社团/文章
	}
}
