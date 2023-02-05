import service from '@/utils/request'

// @Tags OrganizationInformation
// @Summary 创建OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrganizationInformation true "创建OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organizationInformation/createOrganizationInformation [post]
export const createOrganizationInformation = (data) => {
  return service({
    url: '/organizationInformation/createOrganizationInformation',
    method: 'post',
    data
  })
}

// @Tags OrganizationInformation
// @Summary 删除OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrganizationInformation true "删除OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organizationInformation/deleteOrganizationInformation [delete]
export const deleteOrganizationInformation = (data) => {
  return service({
    url: '/organizationInformation/deleteOrganizationInformation',
    method: 'delete',
    data
  })
}

// @Tags OrganizationInformation
// @Summary 删除OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organizationInformation/deleteOrganizationInformation [delete]
export const deleteOrganizationInformationByIds = (data) => {
  return service({
    url: '/organizationInformation/deleteOrganizationInformationByIds',
    method: 'delete',
    data
  })
}

// @Tags OrganizationInformation
// @Summary 更新OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrganizationInformation true "更新OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /organizationInformation/updateOrganizationInformation [put]
export const updateOrganizationInformation = (data) => {
  return service({
    url: '/organizationInformation/updateOrganizationInformation',
    method: 'put',
    data
  })
}

// @Tags OrganizationInformation
// @Summary 用id查询OrganizationInformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrganizationInformation true "用id查询OrganizationInformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /organizationInformation/findOrganizationInformation [get]
export const findOrganizationInformation = (params) => {
  return service({
    url: '/organizationInformation/findOrganizationInformation',
    method: 'get',
    params
  })
}

// @Tags OrganizationInformation
// @Summary 分页获取OrganizationInformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrganizationInformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organizationInformation/getOrganizationInformationList [get]
export const getOrganizationInformationList = (params) => {
  return service({
    url: '/organizationInformation/getOrganizationInformationList',
    method: 'get',
    params
  })
}
