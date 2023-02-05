package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MomentsRouter struct {
}

// InitMomentsRouter 初始化 Moments 路由信息
func (s *MomentsRouter) InitMomentsRouter(Router *gin.RouterGroup) {
	momentsRouter := Router.Group("moments").Use(middleware.OperationRecord())
	momentsRouterWithoutRecord := Router.Group("moments")
	var momentsApi = v1.ApiGroupApp.WebApiGroup.MomentsApi
	{
		momentsRouter.POST("createMoments", momentsApi.CreateMoments)   // 新建Moments
		momentsRouter.DELETE("deleteMoments", momentsApi.DeleteMoments) // 删除Moments
		momentsRouter.DELETE("deleteMomentsByIds", momentsApi.DeleteMomentsByIds) // 批量删除Moments
		momentsRouter.PUT("updateMoments", momentsApi.UpdateMoments)    // 更新Moments
	}
	{
		momentsRouterWithoutRecord.GET("findMoments", momentsApi.FindMoments)        // 根据ID获取Moments
		momentsRouterWithoutRecord.GET("getMomentsList", momentsApi.GetMomentsList)  // 获取Moments列表
	}
}
