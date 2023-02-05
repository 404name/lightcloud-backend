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

type ActivityRecordApi struct {
}

var activityRecordService = service.ServiceGroupApp.WebServiceGroup.ActivityRecordService

// CreateActivityRecord 创建ActivityRecord
// @Tags ActivityRecord
// @Summary 创建ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.ActivityRecord true "创建ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activityRecord/createActivityRecord [post]
// @Version 1.0
func (activityRecordApi *ActivityRecordApi) CreateActivityRecord(c *gin.Context) {
	var activityRecord web.ActivityRecord
	err := c.ShouldBindJSON(&activityRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ActivityId": {utils.NotEmpty()},
		"UserId":     {utils.NotEmpty()},
	}
	if err := utils.Verify(activityRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := activityRecordService.CreateActivityRecord(activityRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteActivityRecord 删除ActivityRecord
// @Tags ActivityRecord
// @Summary 删除ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.ActivityRecord true "删除ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /activityRecord/deleteActivityRecord [delete]
func (activityRecordApi *ActivityRecordApi) DeleteActivityRecord(c *gin.Context) {
	var activityRecord web.ActivityRecord
	err := c.ShouldBindJSON(&activityRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := activityRecordService.DeleteActivityRecord(activityRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActivityRecordByIds 批量删除ActivityRecord
// @Tags ActivityRecord
// @Summary 批量删除ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /activityRecord/deleteActivityRecordByIds [delete]
func (activityRecordApi *ActivityRecordApi) DeleteActivityRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := activityRecordService.DeleteActivityRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActivityRecord 更新ActivityRecord
// @Tags ActivityRecord
// @Summary 更新ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.ActivityRecord true "更新ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /activityRecord/updateActivityRecord [put]
func (activityRecordApi *ActivityRecordApi) UpdateActivityRecord(c *gin.Context) {
	var activityRecord web.ActivityRecord
	err := c.ShouldBindJSON(&activityRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ActivityId": {utils.NotEmpty()},
		"UserId":     {utils.NotEmpty()},
	}
	if err := utils.Verify(activityRecord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := activityRecordService.UpdateActivityRecord(activityRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActivityRecord 用id查询ActivityRecord
// @Tags ActivityRecord
// @Summary 用id查询ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.ActivityRecord true "用id查询ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /activityRecord/findActivityRecord [get]
func (activityRecordApi *ActivityRecordApi) FindActivityRecord(c *gin.Context) {
	var activityRecord web.ActivityRecord
	err := c.ShouldBindQuery(&activityRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reactivityRecord, err := activityRecordService.GetActivityRecord(activityRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reactivityRecord": reactivityRecord}, c)
	}
}

// GetActivityRecordList 分页获取ActivityRecord列表
// @Tags ActivityRecord
// @Summary 分页获取ActivityRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.ActivityRecordSearch true "分页获取ActivityRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activityRecord/getActivityRecordList [get]
func (activityRecordApi *ActivityRecordApi) GetActivityRecordList(c *gin.Context) {
	var pageInfo webReq.ActivityRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := activityRecordService.GetActivityRecordInfoList(pageInfo); err != nil {
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
