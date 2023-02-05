import service from '@/utils/request'

// @Tags Activity
// @Summary 创建Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activity true "创建Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activity/createActivity [post]
export const createActivity = (data) => {
  return service({
    url: '/activity/createActivity',
    method: 'post',
    data
  })
}

// @Tags Activity
// @Summary 删除Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activity true "删除Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /activity/deleteActivity [delete]
export const deleteActivity = (data) => {
  return service({
    url: '/activity/deleteActivity',
    method: 'delete',
    data
  })
}

// @Tags Activity
// @Summary 删除Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /activity/deleteActivity [delete]
export const deleteActivityByIds = (data) => {
  return service({
    url: '/activity/deleteActivityByIds',
    method: 'delete',
    data
  })
}

// @Tags Activity
// @Summary 更新Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Activity true "更新Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /activity/updateActivity [put]
export const updateActivity = (data) => {
  return service({
    url: '/activity/updateActivity',
    method: 'put',
    data
  })
}

// @Tags Activity
// @Summary 用id查询Activity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Activity true "用id查询Activity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /activity/findActivity [get]
export const findActivity = (params) => {
  return service({
    url: '/activity/findActivity',
    method: 'get',
    params
  })
}

// @Tags Activity
// @Summary 分页获取Activity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Activity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activity/getActivityList [get]
export const getActivityList = (params) => {
  return service({
    url: '/activity/getActivityList',
    method: 'get',
    params
  })
}
