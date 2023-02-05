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
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CommentApi struct {
}

var commentService = service.ServiceGroupApp.WebServiceGroup.CommentService


// CreateComment 创建Comment
// @Tags Comment
// @Summary 创建Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Comment true "创建Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/createComment [post]
func (commentApi *CommentApi) CreateComment(c *gin.Context) {
	var comment web.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Content":{utils.NotEmpty()},
        "ToId":{utils.NotEmpty()},
        "ToType":{utils.NotEmpty()},
        "UserId":{utils.NotEmpty()},
    }
	if err := utils.Verify(comment, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := commentService.CreateComment(comment); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteComment 删除Comment
// @Tags Comment
// @Summary 删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Comment true "删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comment/deleteComment [delete]
func (commentApi *CommentApi) DeleteComment(c *gin.Context) {
	var comment web.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.DeleteComment(comment); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommentByIds 批量删除Comment
// @Tags Comment
// @Summary 批量删除Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /comment/deleteCommentByIds [delete]
func (commentApi *CommentApi) DeleteCommentByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.DeleteCommentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateComment 更新Comment
// @Tags Comment
// @Summary 更新Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Comment true "更新Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comment/updateComment [put]
func (commentApi *CommentApi) UpdateComment(c *gin.Context) {
	var comment web.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Content":{utils.NotEmpty()},
          "ToId":{utils.NotEmpty()},
          "ToType":{utils.NotEmpty()},
          "UserId":{utils.NotEmpty()},
      }
    if err := utils.Verify(comment, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := commentService.UpdateComment(comment); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindComment 用id查询Comment
// @Tags Comment
// @Summary 用id查询Comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.Comment true "用id查询Comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comment/findComment [get]
func (commentApi *CommentApi) FindComment(c *gin.Context) {
	var comment web.Comment
	err := c.ShouldBindQuery(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recomment, err := commentService.GetComment(comment.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recomment": recomment}, c)
	}
}

// GetCommentList 分页获取Comment列表
// @Tags Comment
// @Summary 分页获取Comment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.CommentSearch true "分页获取Comment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/getCommentList [get]
func (commentApi *CommentApi) GetCommentList(c *gin.Context) {
	var pageInfo webReq.CommentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := commentService.GetCommentInfoList(pageInfo); err != nil {
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
