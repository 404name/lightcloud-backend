package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
)

type HomeInfoResponse struct {
	Version          string         `json:"version"`          //网站版本
	Describe         string         `json:"describe"`         //网站描述
	Href             string         `json:"href"`             //网站跳转
	HrefDesc         string         `json:"hrefDesc"`         //外链描述
	MomentList       []web.Moments  `json:"momentList"`       // 最新动态列表
	HotMomentList    []web.Moments  `json:"hotMomentList"`    // 最热动态列表
	ActivityList     []web.Activity `json:"activityList"`     // 最新活动列表
	OrganizationList interface{}    `json:"organizationList"` // 推荐社团列表
}
