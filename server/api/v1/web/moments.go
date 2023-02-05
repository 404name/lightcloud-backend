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

type MomentsApi struct {
}

var momentsService = service.ServiceGroupApp.WebServiceGroup.MomentsService


// CreateMoments 创建Moments
// @Tags Moments
// @Summary 创建Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Moments true "创建Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moments/createMoments [post]
func (momentsApi *MomentsApi) CreateMoments(c *gin.Context) {
	var moments web.Moments
	err := c.ShouldBindJSON(&moments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Content":{utils.NotEmpty()},
        "Cover":{utils.NotEmpty()},
        "FromId":{utils.NotEmpty()},
        "FromType":{utils.NotEmpty()},
        "Title":{utils.NotEmpty()},
        "Type":{utils.NotEmpty()},
        "UserId":{utils.NotEmpty()},
    }
	if err := utils.Verify(moments, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := momentsService.CreateMoments(moments); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMoments 删除Moments
// @Tags Moments
// @Summary 删除Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Moments true "删除Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moments/deleteMoments [delete]
func (momentsApi *MomentsApi) DeleteMoments(c *gin.Context) {
	var moments web.Moments
	err := c.ShouldBindJSON(&moments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := momentsService.DeleteMoments(moments); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMomentsByIds 批量删除Moments
// @Tags Moments
// @Summary 批量删除Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /moments/deleteMomentsByIds [delete]
func (momentsApi *MomentsApi) DeleteMomentsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := momentsService.DeleteMomentsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMoments 更新Moments
// @Tags Moments
// @Summary 更新Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Moments true "更新Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moments/updateMoments [put]
func (momentsApi *MomentsApi) UpdateMoments(c *gin.Context) {
	var moments web.Moments
	err := c.ShouldBindJSON(&moments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Content":{utils.NotEmpty()},
          "Cover":{utils.NotEmpty()},
          "FromId":{utils.NotEmpty()},
          "FromType":{utils.NotEmpty()},
          "Title":{utils.NotEmpty()},
          "Type":{utils.NotEmpty()},
          "UserId":{utils.NotEmpty()},
      }
    if err := utils.Verify(moments, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := momentsService.UpdateMoments(moments); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMoments 用id查询Moments
// @Tags Moments
// @Summary 用id查询Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.Moments true "用id查询Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moments/findMoments [get]
func (momentsApi *MomentsApi) FindMoments(c *gin.Context) {
	var moments web.Moments
	err := c.ShouldBindQuery(&moments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remoments, err := momentsService.GetMoments(moments.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remoments": remoments}, c)
	}
}

// GetMomentsList 分页获取Moments列表
// @Tags Moments
// @Summary 分页获取Moments列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.MomentsSearch true "分页获取Moments列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moments/getMomentsList [get]
func (momentsApi *MomentsApi) GetMomentsList(c *gin.Context) {
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
