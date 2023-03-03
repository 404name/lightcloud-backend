package web

import (
	"encoding/json"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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

// CreatComment
// @Tags      用户相关接口
// @Summary   发布评论
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body web.Comment true "创建Comment(在extra里面携带封面)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/creatComment [post]
func (userApi *WebUserApi) CreatComment(c *gin.Context) {
	var comment web.Comment
	var extra map[string]interface{}
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加用户相关附加信息
	userid := int(utils.GetUserID(c))
	comment.UserId = &userid
	json.Unmarshal([]byte(comment.Extra), &extra)
	comment.Extra = utils.GetUserExtra(c)

	// 验证
	verify := utils.Rules{
		"Content": {utils.NotEmpty()},
		"ToId":    {utils.NotEmpty()},
		"ToType":  {utils.NotEmpty()},
		"UserId":  {utils.NotEmpty()},
	}
	if err := utils.Verify(comment, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.CreateComment(comment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		var cover string
		cover, _ = extra["cover"].(string)

		// 后面改成消息队列
		momentsService.CreateMoments(web.Moments{
			Content:  comment.Content,
			Cover:    cover,
			FromId:   comment.ToId,
			FromType: comment.ToType,
			IsTop:    utils.BoolToPoint(false),
			Title:    "发布了一条评论",
			Type:     "动态",
			UserId:   &userid,
			Extra:    utils.GetUserExtra(c),
		})
		response.OkWithMessage("创建成功", c)
	}
}

// CreateArticle
// @Tags      用户相关接口
// @Summary   发布文章
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body web.Article true "创建Article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/createArticle [post]
func (userApi *WebUserApi) CreateArticle(c *gin.Context) {

	var article web.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加附加信息
	userid := int(utils.GetUserID(c))
	article.Uid = &userid
	article.Extra = utils.GetUserExtra(c)

	// 验证
	verify := utils.Rules{
		"CategoryId": {utils.NotEmpty()},
		"Content":    {utils.NotEmpty()},
		"Cover":      {utils.NotEmpty()},
		"Status":     {utils.NotEmpty()},
		"Title":      {utils.NotEmpty()},
		"Uid":        {utils.NotEmpty()},
	}
	if err := utils.Verify(article, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.CreateArticle(&article); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		// 后面改成消息队列
		momentsService.CreateMoments(web.Moments{
			Content:  article.Title,
			Cover:    article.Cover,
			FromId:   utils.UintToInt(article.ID),
			FromType: utils.IntToPoint(0),
			IsTop:    utils.BoolToPoint(false),
			Title:    "发布了一条评论",
			Type:     "动态",
			UserId:   &userid,
			Extra:    utils.GetUserExtra(c),
		})
		response.OkWithMessage("创建成功", c)
	}
}

// CreateActivity
// @Tags      用户相关接口
// @Summary   发布活动
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body web.Activity true "创建Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/createActivity [post]
func (userApi *WebUserApi) CreateActivity(c *gin.Context) {

	var activity web.Activity
	err := c.ShouldBindJSON(&activity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 添加附加信息
	userid := int(utils.GetUserID(c))
	activity.UserId = &userid
	activity.Extra = utils.GetUserExtra(c)
	activity.CreatedBy = utils.GetUserID(c)
	if err := activityService.CreateActivity(&activity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		// 后面改成消息队列
		momentsService.CreateMoments(web.Moments{
			Content:  activity.Title,
			Cover:    activity.Cover,
			FromId:   utils.UintToInt(activity.ID),
			FromType: utils.IntToPoint(1),
			IsTop:    utils.BoolToPoint(false),
			Title:    "发布了新的活动",
			Type:     "动态",
			UserId:   &userid,
			Extra:    utils.GetUserExtra(c),
		})
		response.OkWithMessage("创建成功", c)
	}
}

// CreateIntroduce
// @Tags      用户相关接口
// @Summary   发布内推
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body web.Introduce true "创建Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/createIntroduce [post]
func (userApi *WebUserApi) CreateIntroduce(c *gin.Context) {

	var introduce web.Introduce
	var extra map[string]interface{}

	err := c.ShouldBindJSON(&introduce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加附加信息
	userid := int(utils.GetUserID(c))
	introduce.UserId = &userid
	json.Unmarshal([]byte(introduce.Extra), &extra)
	introduce.Extra = utils.GetUserExtra(c)

	verify := utils.Rules{
		"ActivityId": {utils.NotEmpty()},
		"UserId":     {utils.NotEmpty()},
	}
	if err := utils.Verify(introduce, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := introduceService.CreateIntroduce(introduce); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		var (
			cover         string
			activityTitle string
		)
		cover, _ = extra["cover"].(string)
		activityTitle, _ = extra["activityTitle"].(string)

		// 后面改成消息队列
		momentsService.CreateMoments(web.Moments{
			Content:  "快使用我的内推码：【" + introduce.Code + "】来参加活动" + activityTitle + "吧",
			Cover:    cover,
			FromId:   introduce.ActivityId,
			FromType: utils.IntToPoint(1),
			IsTop:    utils.BoolToPoint(false),
			Title:    "发布了活动内推",
			Type:     "动态",
			UserId:   &userid,
			Extra:    utils.GetUserExtra(c),
		})
		response.OkWithMessage("创建成功", c)
	}
}

// AttendActivity
// @Tags      用户相关接口
// @Summary   参与活动
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body web.ActivityRecord true "创建ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/attendActivity [post]
func (userApi *WebUserApi) AttendActivity(c *gin.Context) {

	var activityRecord web.ActivityRecord
	var extra map[string]interface{}

	err := c.ShouldBindJSON(&activityRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加附加信息
	userid := int(utils.GetUserID(c))
	activityRecord.UserId = &userid
	json.Unmarshal([]byte(activityRecord.Extra), &extra)
	activityRecord.Extra = utils.GetUserExtra(c)

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
		var (
			cover         string
			activityTitle string
		)
		cover, _ = extra["cover"].(string)
		activityTitle, _ = extra["activityTitle"].(string)

		// 后面改成消息队列
		momentsService.CreateMoments(web.Moments{
			Content:  "使用内推码：【" + activityRecord.Code + "】参加了活动" + activityTitle,
			Cover:    cover,
			FromId:   activityRecord.ActivityId,
			FromType: utils.IntToPoint(1),
			IsTop:    utils.BoolToPoint(false),
			Title:    "参与了活动",
			Type:     "动态",
			UserId:   &userid,
			Extra:    utils.GetUserExtra(c),
		})
		response.OkWithMessage("创建成功", c)
	}
}

// CreateInformation
// @Tags      用户相关接口
// @Summary   填写简历/添加社团信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body web.OrganizationInformation true "创建OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/createInformation [post]
func (userApi *WebUserApi) CreateInformation(c *gin.Context) {
	var organizationInformation web.OrganizationInformation
	err := c.ShouldBindJSON(&organizationInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加附加信息
	userid := int(utils.GetUserID(c))
	organizationInformation.Oid = &userid

	if err := organizationInformationService.CreateOrganizationInformation(organizationInformation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// VoteComment
// @Tags      用户相关接口
// @Summary   点赞/点踩评论
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body webReq.VoteRequest true "点赞or点踩评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/voteComment [post]
func (userApi *WebUserApi) VoteComment(c *gin.Context) {
	var vote webReq.VoteRequest
	err := c.ShouldBindJSON(&vote)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.ApprovalComment(vote.CommentId, vote.Approval); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// Collection
// @Tags      用户相关接口
// @Summary   收藏社团/文章
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body webReq.CollectionRequest true "收藏社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router    /user/collection [post]
func (userApi *WebUserApi) Collection(c *gin.Context) {

	var req webReq.CollectionRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if DelFlag, err := userService.AddOrRemoveCollection(utils.GetUserID(c), req.Oid); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {

		if !DelFlag {
			// 后面改成消息队列
			momentsService.CreateMoments(web.Moments{
				Content:  req.Describe,
				Cover:    req.Cover,
				FromId:   utils.UintToInt(req.Oid),
				FromType: utils.IntToPoint(2),
				IsTop:    utils.BoolToPoint(false),
				Title:    "收藏了社团:" + req.Name,
				Type:     "动态",
				UserId:   utils.UintToInt(utils.GetUserID(c)),
				Extra:    utils.GetUserExtra(c),
			})
		}
		response.OkWithMessage("操作成功", c)
	}
}
