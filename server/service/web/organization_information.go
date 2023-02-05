package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
)

type OrganizationInformationService struct {
}

// CreateOrganizationInformation 创建OrganizationInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationInformationService *OrganizationInformationService) CreateOrganizationInformation(organizationInformation web.OrganizationInformation) (err error) {
	err = global.GVA_DB.Create(&organizationInformation).Error
	return err
}

// DeleteOrganizationInformation 删除OrganizationInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationInformationService *OrganizationInformationService)DeleteOrganizationInformation(organizationInformation web.OrganizationInformation) (err error) {
	err = global.GVA_DB.Delete(&organizationInformation).Error
	return err
}

// DeleteOrganizationInformationByIds 批量删除OrganizationInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationInformationService *OrganizationInformationService)DeleteOrganizationInformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.OrganizationInformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOrganizationInformation 更新OrganizationInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationInformationService *OrganizationInformationService)UpdateOrganizationInformation(organizationInformation web.OrganizationInformation) (err error) {
	err = global.GVA_DB.Save(&organizationInformation).Error
	return err
}

// GetOrganizationInformation 根据id获取OrganizationInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationInformationService *OrganizationInformationService)GetOrganizationInformation(id uint) (organizationInformation web.OrganizationInformation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&organizationInformation).Error
	return
}

// GetOrganizationInformationInfoList 分页获取OrganizationInformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationInformationService *OrganizationInformationService)GetOrganizationInformationInfoList(info webReq.OrganizationInformationSearch) (list []web.OrganizationInformation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&web.OrganizationInformation{})
    var organizationInformations []web.OrganizationInformation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Oid != nil {
        db = db.Where("oid = ?",info.Oid)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["sort"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&organizationInformations).Error
	return  organizationInformations, total, err
}
