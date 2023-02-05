package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/web"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OrganizationInformationApi struct {
}

var organizationInformationService = service.ServiceGroupApp.WebServiceGroup.OrganizationInformationService


// CreateOrganizationInformation 创建OrganizationInformation
// @Tags OrganizationInformation
// @Summary 创建OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.OrganizationInformation true "创建OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organizationInformation/createOrganizationInformation [post]
func (organizationInformationApi *OrganizationInformationApi) CreateOrganizationInformation(c *gin.Context) {
	var organizationInformation web.OrganizationInformation
	err := c.ShouldBindJSON(&organizationInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := organizationInformationService.CreateOrganizationInformation(organizationInformation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrganizationInformation 删除OrganizationInformation
// @Tags OrganizationInformation
// @Summary 删除OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.OrganizationInformation true "删除OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organizationInformation/deleteOrganizationInformation [delete]
func (organizationInformationApi *OrganizationInformationApi) DeleteOrganizationInformation(c *gin.Context) {
	var organizationInformation web.OrganizationInformation
	err := c.ShouldBindJSON(&organizationInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := organizationInformationService.DeleteOrganizationInformation(organizationInformation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrganizationInformationByIds 批量删除OrganizationInformation
// @Tags OrganizationInformation
// @Summary 批量删除OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /organizationInformation/deleteOrganizationInformationByIds [delete]
func (organizationInformationApi *OrganizationInformationApi) DeleteOrganizationInformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := organizationInformationService.DeleteOrganizationInformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrganizationInformation 更新OrganizationInformation
// @Tags OrganizationInformation
// @Summary 更新OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.OrganizationInformation true "更新OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /organizationInformation/updateOrganizationInformation [put]
func (organizationInformationApi *OrganizationInformationApi) UpdateOrganizationInformation(c *gin.Context) {
	var organizationInformation web.OrganizationInformation
	err := c.ShouldBindJSON(&organizationInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := organizationInformationService.UpdateOrganizationInformation(organizationInformation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrganizationInformation 用id查询OrganizationInformation
// @Tags OrganizationInformation
// @Summary 用id查询OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.OrganizationInformation true "用id查询OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /organizationInformation/findOrganizationInformation [get]
func (organizationInformationApi *OrganizationInformationApi) FindOrganizationInformation(c *gin.Context) {
	var organizationInformation web.OrganizationInformation
	err := c.ShouldBindQuery(&organizationInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorganizationInformation, err := organizationInformationService.GetOrganizationInformation(organizationInformation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorganizationInformation": reorganizationInformation}, c)
	}
}

// GetOrganizationInformationList 分页获取OrganizationInformation列表
// @Tags OrganizationInformation
// @Summary 分页获取OrganizationInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.OrganizationInformationSearch true "分页获取OrganizationInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organizationInformation/getOrganizationInformationList [get]
func (organizationInformationApi *OrganizationInformationApi) GetOrganizationInformationList(c *gin.Context) {
	var pageInfo webReq.OrganizationInformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := organizationInformationService.GetOrganizationInformationInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
