import service from '@/utils/request'

// @Tags Introduce
// @Summary 创建Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Introduce true "创建Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /introduce/createIntroduce [post]
export const createIntroduce = (data) => {
  return service({
    url: '/introduce/createIntroduce',
    method: 'post',
    data
  })
}

// @Tags Introduce
// @Summary 删除Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Introduce true "删除Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /introduce/deleteIntroduce [delete]
export const deleteIntroduce = (data) => {
  return service({
    url: '/introduce/deleteIntroduce',
    method: 'delete',
    data
  })
}

// @Tags Introduce
// @Summary 删除Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /introduce/deleteIntroduce [delete]
export const deleteIntroduceByIds = (data) => {
  return service({
    url: '/introduce/deleteIntroduceByIds',
    method: 'delete',
    data
  })
}

// @Tags Introduce
// @Summary 更新Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Introduce true "更新Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /introduce/updateIntroduce [put]
export const updateIntroduce = (data) => {
  return service({
    url: '/introduce/updateIntroduce',
    method: 'put',
    data
  })
}

// @Tags Introduce
// @Summary 用id查询Introduce
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Introduce true "用id查询Introduce"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /introduce/findIntroduce [get]
export const findIntroduce = (params) => {
  return service({
    url: '/introduce/findIntroduce',
    method: 'get',
    params
  })
}

// @Tags Introduce
// @Summary 分页获取Introduce列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Introduce列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /introduce/getIntroduceList [get]
export const getIntroduceList = (params) => {
  return service({
    url: '/introduce/getIntroduceList',
    method: 'get',
    params
  })
}
