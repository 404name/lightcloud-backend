package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IntroduceRouter struct {
}

// InitIntroduceRouter 初始化 Introduce 路由信息
func (s *IntroduceRouter) InitIntroduceRouter(Router *gin.RouterGroup) {
	introduceRouter := Router.Group("introduce").Use(middleware.OperationRecord())
	introduceRouterWithoutRecord := Router.Group("introduce")
	var introduceApi = v1.ApiGroupApp.WebApiGroup.IntroduceApi
	{
		introduceRouter.POST("createIntroduce", introduceApi.CreateIntroduce)   // 新建Introduce
		introduceRouter.DELETE("deleteIntroduce", introduceApi.DeleteIntroduce) // 删除Introduce
		introduceRouter.DELETE("deleteIntroduceByIds", introduceApi.DeleteIntroduceByIds) // 批量删除Introduce
		introduceRouter.PUT("updateIntroduce", introduceApi.UpdateIntroduce)    // 更新Introduce
	}
	{
		introduceRouterWithoutRecord.GET("findIntroduce", introduceApi.FindIntroduce)        // 根据ID获取Introduce
		introduceRouterWithoutRecord.GET("getIntroduceList", introduceApi.GetIntroduceList)  // 获取Introduce列表
	}
}
