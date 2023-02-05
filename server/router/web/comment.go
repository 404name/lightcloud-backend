package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
}

// InitCommentRouter 初始化 Comment 路由信息
func (s *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	commentRouter := Router.Group("comment").Use(middleware.OperationRecord())
	commentRouterWithoutRecord := Router.Group("comment")
	var commentApi = v1.ApiGroupApp.WebApiGroup.CommentApi
	{
		commentRouter.POST("createComment", commentApi.CreateComment)   // 新建Comment
		commentRouter.DELETE("deleteComment", commentApi.DeleteComment) // 删除Comment
		commentRouter.DELETE("deleteCommentByIds", commentApi.DeleteCommentByIds) // 批量删除Comment
		commentRouter.PUT("updateComment", commentApi.UpdateComment)    // 更新Comment
	}
	{
		commentRouterWithoutRecord.GET("findComment", commentApi.FindComment)        // 根据ID获取Comment
		commentRouterWithoutRecord.GET("getCommentList", commentApi.GetCommentList)  // 获取Comment列表
	}
}
