package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ActivityApi struct {
}

var activityService = service.ServiceGroupApp.WebServiceGroup.ActivityService

// CreateActivity 创建Activity
// @Tags Activity
// @Summary 创建Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Activity true "创建Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activity/createActivity [post]
func (activityApi *ActivityApi) CreateActivity(c *gin.Context) {
	var activity web.Activity
	err := c.ShouldBindJSON(&activity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	activity.CreatedBy = utils.GetUserID(c)
	if err := activityService.CreateActivity(&activity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteActivity 删除Activity
// @Tags Activity
// @Summary 删除Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Activity true "删除Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /activity/deleteActivity [delete]
func (activityApi *ActivityApi) DeleteActivity(c *gin.Context) {
	var activity web.Activity
	err := c.ShouldBindJSON(&activity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	activity.DeletedBy = utils.GetUserID(c)
	if err := activityService.DeleteActivity(activity); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActivityByIds 批量删除Activity
// @Tags Activity
// @Summary 批量删除Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /activity/deleteActivityByIds [delete]
func (activityApi *ActivityApi) DeleteActivityByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := activityService.DeleteActivityByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActivity 更新Activity
// @Tags Activity
// @Summary 更新Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Activity true "更新Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /activity/updateActivity [put]
func (activityApi *ActivityApi) UpdateActivity(c *gin.Context) {
	var activity web.Activity
	err := c.ShouldBindJSON(&activity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	activity.UpdatedBy = utils.GetUserID(c)
	if err := activityService.UpdateActivity(activity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActivity 用id查询Activity
// @Tags Activity
// @Summary 用id查询Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.Activity true "用id查询Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /activity/findActivity [get]
func (activityApi *ActivityApi) FindActivity(c *gin.Context) {
	var activity web.Activity
	err := c.ShouldBindQuery(&activity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reactivity, err := activityService.GetActivity(activity.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reactivity": reactivity}, c)
	}
}

// GetActivityList 分页获取Activity列表
// @Tags Activity
// @Summary 分页获取Activity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.ActivitySearch true "分页获取Activity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activity/getActivityList [get]
func (activityApi *ActivityApi) GetActivityList(c *gin.Context) {
	var pageInfo webReq.ActivitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := activityService.GetActivityInfoList(pageInfo); err != nil {
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
