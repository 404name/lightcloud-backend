// 自动生成模板OrganizationInformation
package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// OrganizationInformation 结构体
type OrganizationInformation struct {
      global.GVA_MODEL
      Content  string `json:"content" form:"content" gorm:"column:content;comment:信息json内容;"`
      Oid  *int `json:"oid" form:"oid" gorm:"column:oid;comment:社团id(用户id);size:10;"`
      Need  *bool `json:"need" form:"need" gorm:"column:need;comment:评分;"`
      Rendering  string `json:"rendering" form:"rendering" gorm:"column:rendering;comment:渲染组件(默认);size:255;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:板块排序;size:10;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:信息类别;size:10;"`
}


// TableName OrganizationInformation 表名
func (OrganizationInformation) TableName() string {
  return "web_organization_information"
}

