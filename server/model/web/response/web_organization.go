package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
)

type OrganizationHomePageResp struct {
	BaseInfo  *system.SysUser               `json:"baseInfo" form:"baseInfo"`   // 基础信息
	ExtraInfo []web.OrganizationInformation `json:"extraInfo" form:"extraInfo"` // 附加模块
	Comment   []web.Comment                 `json:"comment" form:"comment"`     // 评价信息
	Activity  []web.Activity                `json:"activity" form:"activity"`   // 活动
	Moments   []web.Moments                 `json:"moments" form:"moments"`     // 动态
}

type ActivityResp struct {
	BaseInfo     web.Activity         `json:"baseInfo" form:"baseInfo"`         // 基础信息
	Organization *system.SysUser      `json:"organization" form:"organization"` // 发起社团
	ExtraInfo    []web.ActivityRecord `json:"extraInfo" form:"extraInfo"`       // 参与人员
	Comment      []web.Comment        `json:"comment" form:"comment"`           // 评论
	Moments      []web.Moments        `json:"moments" form:"moments"`           // 动态
}
