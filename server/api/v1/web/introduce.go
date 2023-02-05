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

type IntroduceApi struct {
}

var introduceService = service.ServiceGroupApp.WebServiceGroup.IntroduceService


// CreateIntroduce 创建Introduce
// @Tags Introduce
// @Summary 创建Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Introduce true "创建Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /introduce/createIntroduce [post]
func (introduceApi *IntroduceApi) CreateIntroduce(c *gin.Context) {
	var introduce web.Introduce
	err := c.ShouldBindJSON(&introduce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "ActivityId":{utils.NotEmpty()},
        "UserId":{utils.NotEmpty()},
    }
	if err := utils.Verify(introduce, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := introduceService.CreateIntroduce(introduce); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIntroduce 删除Introduce
// @Tags Introduce
// @Summary 删除Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Introduce true "删除Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /introduce/deleteIntroduce [delete]
func (introduceApi *IntroduceApi) DeleteIntroduce(c *gin.Context) {
	var introduce web.Introduce
	err := c.ShouldBindJSON(&introduce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := introduceService.DeleteIntroduce(introduce); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIntroduceByIds 批量删除Introduce
// @Tags Introduce
// @Summary 批量删除Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /introduce/deleteIntroduceByIds [delete]
func (introduceApi *IntroduceApi) DeleteIntroduceByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := introduceService.DeleteIntroduceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIntroduce 更新Introduce
// @Tags Introduce
// @Summary 更新Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body web.Introduce true "更新Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /introduce/updateIntroduce [put]
func (introduceApi *IntroduceApi) UpdateIntroduce(c *gin.Context) {
	var introduce web.Introduce
	err := c.ShouldBindJSON(&introduce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "ActivityId":{utils.NotEmpty()},
          "UserId":{utils.NotEmpty()},
      }
    if err := utils.Verify(introduce, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := introduceService.UpdateIntroduce(introduce); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIntroduce 用id查询Introduce
// @Tags Introduce
// @Summary 用id查询Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query web.Introduce true "用id查询Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /introduce/findIntroduce [get]
func (introduceApi *IntroduceApi) FindIntroduce(c *gin.Context) {
	var introduce web.Introduce
	err := c.ShouldBindQuery(&introduce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reintroduce, err := introduceService.GetIntroduce(introduce.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reintroduce": reintroduce}, c)
	}
}

// GetIntroduceList 分页获取Introduce列表
// @Tags Introduce
// @Summary 分页获取Introduce列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webReq.IntroduceSearch true "分页获取Introduce列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /introduce/getIntroduceList [get]
func (introduceApi *IntroduceApi) GetIntroduceList(c *gin.Context) {
	var pageInfo webReq.IntroduceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := introduceService.GetIntroduceInfoList(pageInfo); err != nil {
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
