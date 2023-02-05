import service from '@/utils/request'

// @Tags Moments
// @Summary 创建Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Moments true "创建Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moments/createMoments [post]
export const createMoments = (data) => {
  return service({
    url: '/moments/createMoments',
    method: 'post',
    data
  })
}

// @Tags Moments
// @Summary 删除Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Moments true "删除Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moments/deleteMoments [delete]
export const deleteMoments = (data) => {
  return service({
    url: '/moments/deleteMoments',
    method: 'delete',
    data
  })
}

// @Tags Moments
// @Summary 删除Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /moments/deleteMoments [delete]
export const deleteMomentsByIds = (data) => {
  return service({
    url: '/moments/deleteMomentsByIds',
    method: 'delete',
    data
  })
}

// @Tags Moments
// @Summary 更新Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Moments true "更新Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /moments/updateMoments [put]
export const updateMoments = (data) => {
  return service({
    url: '/moments/updateMoments',
    method: 'put',
    data
  })
}

// @Tags Moments
// @Summary 用id查询Moments
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Moments true "用id查询Moments"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /moments/findMoments [get]
export const findMoments = (params) => {
  return service({
    url: '/moments/findMoments',
    method: 'get',
    params
  })
}

// @Tags Moments
// @Summary 分页获取Moments列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Moments列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /moments/getMomentsList [get]
export const getMomentsList = (params) => {
  return service({
    url: '/moments/getMomentsList',
    method: 'get',
    params
  })
}
