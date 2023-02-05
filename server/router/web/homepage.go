package web

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type HomePageRouter struct {
}

// InitHomePageRouter 初始化 HomePage 路由信息
func (s *HomePageRouter) InitHomePageRouter(Router *gin.RouterGroup) {
	// homePageRouter := Router.Group("homepage").Use(middleware.OperationRecord())
	homePageRouterWithoutRecord := Router.Group("homepage")
	var homePageApi = v1.ApiGroupApp.WebApiGroup.HomePageApi
	{
		homePageRouterWithoutRecord.GET("organizationList", homePageApi.GetOrganizationList) // 根据ID获取HomePage
	}
}
