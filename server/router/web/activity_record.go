package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ActivityRecordRouter struct {
}

// InitActivityRecordRouter 初始化 ActivityRecord 路由信息
func (s *ActivityRecordRouter) InitActivityRecordRouter(Router *gin.RouterGroup) {
	activityRecordRouter := Router.Group("activityRecord").Use(middleware.OperationRecord())
	activityRecordRouterWithoutRecord := Router.Group("activityRecord")
	var activityRecordApi = v1.ApiGroupApp.WebApiGroup.ActivityRecordApi
	{
		activityRecordRouter.POST("createActivityRecord", activityRecordApi.CreateActivityRecord)   // 新建ActivityRecord
		activityRecordRouter.DELETE("deleteActivityRecord", activityRecordApi.DeleteActivityRecord) // 删除ActivityRecord
		activityRecordRouter.DELETE("deleteActivityRecordByIds", activityRecordApi.DeleteActivityRecordByIds) // 批量删除ActivityRecord
		activityRecordRouter.PUT("updateActivityRecord", activityRecordApi.UpdateActivityRecord)    // 更新ActivityRecord
	}
	{
		activityRecordRouterWithoutRecord.GET("findActivityRecord", activityRecordApi.FindActivityRecord)        // 根据ID获取ActivityRecord
		activityRecordRouterWithoutRecord.GET("getActivityRecordList", activityRecordApi.GetActivityRecordList)  // 获取ActivityRecord列表
	}
}
