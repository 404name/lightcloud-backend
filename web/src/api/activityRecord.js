import service from '@/utils/request'

// @Tags ActivityRecord
// @Summary 创建ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivityRecord true "创建ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activityRecord/createActivityRecord [post]
export const createActivityRecord = (data) => {
  return service({
    url: '/activityRecord/createActivityRecord',
    method: 'post',
    data
  })
}

// @Tags ActivityRecord
// @Summary 删除ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivityRecord true "删除ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /activityRecord/deleteActivityRecord [delete]
export const deleteActivityRecord = (data) => {
  return service({
    url: '/activityRecord/deleteActivityRecord',
    method: 'delete',
    data
  })
}

// @Tags ActivityRecord
// @Summary 删除ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /activityRecord/deleteActivityRecord [delete]
export const deleteActivityRecordByIds = (data) => {
  return service({
    url: '/activityRecord/deleteActivityRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags ActivityRecord
// @Summary 更新ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ActivityRecord true "更新ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /activityRecord/updateActivityRecord [put]
export const updateActivityRecord = (data) => {
  return service({
    url: '/activityRecord/updateActivityRecord',
    method: 'put',
    data
  })
}

// @Tags ActivityRecord
// @Summary 用id查询ActivityRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ActivityRecord true "用id查询ActivityRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /activityRecord/findActivityRecord [get]
export const findActivityRecord = (params) => {
  return service({
    url: '/activityRecord/findActivityRecord',
    method: 'get',
    params
  })
}

// @Tags ActivityRecord
// @Summary 分页获取ActivityRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ActivityRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activityRecord/getActivityRecordList [get]
export const getActivityRecordList = (params) => {
  return service({
    url: '/activityRecord/getActivityRecordList',
    method: 'get',
    params
  })
}
