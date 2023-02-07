package web

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
	webResp "github.com/flipped-aurora/gin-vue-admin/server/model/web/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebOrganizationApi struct {
}

// GetHomePage
// @Tags      社团&活动模块
// @Summary   获取社团主页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   0   {object}  response.Response{data=response.OrganizationHomePageResp,msg=string} 获取社团主页
// @Router    /organization/:id/homePage [GET]
func (WebOrganizationApi *WebOrganizationApi) GetHomePage(c *gin.Context) {

	var (
		id   int
		err  error
		resp webResp.OrganizationHomePageResp
		page request.PageInfo = request.PageInfo{Page: 1, PageSize: 100}
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败(未找到社团)", c)
		return
	}

	// 获取社团信息
	if resp.BaseInfo, err = userService.FindUserById(id); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 获取社团模块

	if resp.ExtraInfo, _, err = organizationInformationService.GetOrganizationInformationInfoList(webReq.OrganizationInformationSearch{PageInfo: page, OrganizationInformation: web.OrganizationInformation{Oid: &id}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 获取社团评价
	if resp.Comment, _, err = commentService.GetCommentInfoList(
		webReq.CommentSearch{PageInfo: page, Comment: web.Comment{ToId: &id, ToType: utils.IntToPoint(global.CommentTypeUser)}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 获取社团活动
	if resp.Activity, _, err = activityService.GetActivityInfoList(webReq.ActivitySearch{PageInfo: page, Activity: web.Activity{UserId: &id}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 获取社团动态
	if resp.Moments, _, err = momentsService.GetMomentsInfoList(webReq.MomentsSearch{PageInfo: page, Moments: web.Moments{FromId: &id, FromType: utils.IntToPoint(global.MomentsTypeOrganization)}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(resp, "获取成功", c)
}

// GetActivityDetail
// @Tags      社团&活动模块
// @Summary   活动详细页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   0   {object}  response.Response{data=response.ActivityResp,msg=string} "分页获取社团列表,返回包括列表,总数,页码,每页数量"
// @Router    /organization/activity/:id [GET]
func (WebOrganizationApi *WebOrganizationApi) GetActivityDetail(c *gin.Context) {
	var (
		id   int
		err  error
		resp webResp.ActivityResp
		page request.PageInfo = request.PageInfo{Page: 1, PageSize: 100}
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败(未找到社团)", c)
		return
	}
	// 获取活动信息
	if resp.BaseInfo, err = activityService.GetActivity(uint(id)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 获取社团信息
	if resp.Organization, err = userService.FindUserById(*resp.BaseInfo.UserId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 获取社团模块

	if resp.ExtraInfo, _, err = activityRecordService.GetActivityRecordInfoList(webReq.ActivityRecordSearch{PageInfo: page, ActivityRecord: web.ActivityRecord{ActivityId: &id}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 获取社团评价
	if resp.Comment, _, err = commentService.GetCommentInfoList(
		webReq.CommentSearch{PageInfo: page, Comment: web.Comment{ToType: utils.IntToPoint(global.CommentTypeActivity), ToId: &id}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 获取社团动态
	if resp.Moments, _, err = momentsService.GetMomentsInfoList(webReq.MomentsSearch{PageInfo: page, Moments: web.Moments{FromType: utils.IntToPoint(global.MomentsTypeActivity), FromId: &id}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(resp, "获取成功", c)
}

// GetActivityDetail
// @Tags      社团&活动模块
// @Summary   发送邮箱通知参与人员(开发中)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   0   {object}  response.Response{data=response.ActivityResp,msg=string} "分页获取社团列表,返回包括列表,总数,页码,每页数量"
// @Router    /organization/activity/:id/emailAll [post]
func (WebOrganizationApi *WebOrganizationApi) SentActivityEmail(c *gin.Context) {
	response.OkWithDetailed(nil, "开发中", c)
}
