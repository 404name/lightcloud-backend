package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
	webResp "github.com/flipped-aurora/gin-vue-admin/server/model/web/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebHomePageApi struct {
}

// GetOrganizationList
// @Tags      HomePage
// @Summary   分页获取社团列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取社团列表,返回包括列表,总数,页码,每页数量"
// @Router    /homepage/organizationList [get]
func (homePageApi *WebHomePageApi) GetOrganizationList(c *gin.Context) {

	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userService.GetOrganizationList(pageInfo); err != nil {
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

// Info
// @Tags      首页相关接口
// @Summary   首页信息接口
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   0   {object}  response.Response{data=webResp.HomeInfoResponse,msg=string} "首页信息接口,返回最新动态，最热动态，最新活动"
// @Router    /homepage [GET]
func (homePageApi *WebHomePageApi) GetHomeInfo(c *gin.Context) {
	var err error
	var (
		momentList       []web.Moments
		hotMomentList    []web.Moments
		activityList     []web.Activity
		organizationList interface{}
	)

	// 获取最新动态
	if momentList, _, err = momentsService.GetMomentsInfoList(webReq.MomentsSearch{PageInfo: request.PageInfo{Page: 1, PageSize: 5}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 获取最热动态
	if hotMomentList, _, err = momentsService.GetMomentsInfoList(webReq.MomentsSearch{PageInfo: request.PageInfo{Page: 1, PageSize: 5}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 获取最新活动
	if activityList, _, err = activityService.GetActivityInfoList(webReq.ActivitySearch{PageInfo: request.PageInfo{Page: 1, PageSize: 5}}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	if organizationList, _, err = userService.GetOrganizationList(request.PageInfo{Page: 1, PageSize: 5}); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	response.OkWithDetailed(webResp.HomeInfoResponse{
		Version:          "v1.0.0",
		Describe:         "v1.0.0版本上线了参与调研反馈bug赢好礼",
		Href:             "https://www.bilibili.com/",
		HrefDesc:         "了解更多",
		MomentList:       momentList,
		HotMomentList:    hotMomentList,
		ActivityList:     activityList,
		OrganizationList: organizationList,
	}, "获取成功", c)
}

// GetMomentList
// @Tags      首页相关接口
// @Summary   获取首页动态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query webReq.MomentsSearch true "分页获取Moments列表"
// @Success   0   {object}  response.Response{data=response.PageResult,msg=string} "分页获取社团列表,返回包括列表,总数,页码,每页数量"
// @Router    /homepage/momentList [GET]
func (homePageApi *WebHomePageApi) GetMomentList(c *gin.Context) {

	var pageInfo webReq.MomentsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := momentsService.GetMomentsInfoList(pageInfo); err != nil {
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
