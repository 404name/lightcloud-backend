package web

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrganizationInformationRouter struct {
}

// InitOrganizationInformationRouter 初始化 OrganizationInformation 路由信息
func (s *OrganizationInformationRouter) InitOrganizationInformationRouter(Router *gin.RouterGroup) {
	organizationInformationRouter := Router.Group("organizationInformation").Use(middleware.OperationRecord())
	organizationInformationRouterWithoutRecord := Router.Group("organizationInformation")
	var organizationInformationApi = v1.ApiGroupApp.WebApiGroup.OrganizationInformationApi
	{
		organizationInformationRouter.POST("createOrganizationInformation", organizationInformationApi.CreateOrganizationInformation)             // 新建OrganizationInformation
		organizationInformationRouter.DELETE("deleteOrganizationInformation", organizationInformationApi.DeleteOrganizationInformation)           // 删除OrganizationInformation
		organizationInformationRouter.DELETE("deleteOrganizationInformationByIds", organizationInformationApi.DeleteOrganizationInformationByIds) // 批量删除OrganizationInformation
		organizationInformationRouter.PUT("updateOrganizationInformation", organizationInformationApi.UpdateOrganizationInformation)              // 更新OrganizationInformation
	}
	{
		organizationInformationRouterWithoutRecord.GET("findOrganizationInformation", organizationInformationApi.FindOrganizationInformation)       // 根据ID获取OrganizationInformation
		organizationInformationRouterWithoutRecord.GET("getOrganizationInformationList", organizationInformationApi.GetOrganizationInformationList) // 获取OrganizationInformation列表
	}
}
