package web

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebUserApi struct {
}

// GetUserInfo
// @Tags      用户相关接口
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/:id [GET]
func (WebUserApi *WebUserApi) GetUserInfo(c *gin.Context) {

	id := c.Param("id")
	if userId, err := strconv.Atoi(id); err != nil {
		response.FailWithMessage("未知的用户id", c)
		return
	} else {
		ReqUser, err := userService.FindUserById(userId)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
			return
		}
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

// GetMomentList
// @Tags      用户相关接口
// @Summary   获取用户相关动态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query request.PageInfo true "分页获取Moments列表"
// @Success   0   {object}  response.Response{data=response.PageResult,msg=string} "分页获取动态列表"
// @Router    /user/:id/momentList [GET]
func (userApi *WebUserApi) GetMomentList(c *gin.Context) {

	var (
		pageInfo request.PageInfo
		userId   int
		err      error
	)

	if err = c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userId, err = strconv.Atoi(c.Param("id")); err != nil {
		response.FailWithMessage("未知的用户id", c)
		return
	}

	if list, total, err := momentsService.GetMomentsInfoList(
		webReq.MomentsSearch{PageInfo: pageInfo, Moments: web.Moments{UserId: &userId}}); err != nil {
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

// GetActivityList
// @Tags      用户相关接口
// @Summary   获取用户参加的活动
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query request.PageInfo true "分页获取用户参与活动列表"
// @Success   0   {object}  response.Response{data=response.PageResult,msg=string} "分页获取用户参与活动列表"
// @Router    /user/:id/activityList [GET]
func (userApi *WebUserApi) GetActivityList(c *gin.Context) {

	var (
		pageInfo request.PageInfo
		userId   int
		err      error
	)

	if err = c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userId, err = strconv.Atoi(c.Param("id")); err != nil {
		response.FailWithMessage("未知的用户id", c)
		return
	}

	if list, total, err := activityRecordService.GetActivityRecordInfoDetailList(webReq.ActivityRecordSearch{PageInfo: pageInfo, ActivityRecord: web.ActivityRecord{UserId: &userId}}); err != nil {
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

// GetCommentList
// @Tags      用户相关接口
// @Summary   获取用户发表的评论
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query request.PageInfo true "分页获取用户评论列表"
// @Success   0   {object}  response.Response{data=response.PageResult,msg=string} "分页获取用户评论列表"
// @Router    /user/:id/commentList [GET]
func (userApi *WebUserApi) GetCommentList(c *gin.Context) {

	var (
		pageInfo request.PageInfo
		userId   int
		err      error
	)

	if err = c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userId, err = strconv.Atoi(c.Param("id")); err != nil {
		response.FailWithMessage("未知的用户id", c)
		return
	}

	if list, total, err := commentService.GetCommentInfoList(webReq.CommentSearch{PageInfo: pageInfo, Comment: web.Comment{UserId: &userId}}); err != nil {
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
