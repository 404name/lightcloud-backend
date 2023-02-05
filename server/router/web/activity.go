package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct {
}

// InitActivityRouter 初始化 Activity 路由信息
func (s *ActivityRouter) InitActivityRouter(Router *gin.RouterGroup) {
	activityRouter := Router.Group("activity").Use(middleware.OperationRecord())
	activityRouterWithoutRecord := Router.Group("activity")
	var activityApi = v1.ApiGroupApp.WebApiGroup.ActivityApi
	{
		activityRouter.POST("createActivity", activityApi.CreateActivity)             // 新建Activity
		activityRouter.DELETE("deleteActivity", activityApi.DeleteActivity)           // 删除Activity
		activityRouter.DELETE("deleteActivityByIds", activityApi.DeleteActivityByIds) // 批量删除Activity
		activityRouter.PUT("updateActivity", activityApi.UpdateActivity)              // 更新Activity
	}
	{
		activityRouterWithoutRecord.GET("findActivity", activityApi.FindActivity)       // 根据ID获取Activity
		activityRouterWithoutRecord.GET("getActivityList", activityApi.GetActivityList) // 获取Activity列表
	}
}
