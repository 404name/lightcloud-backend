package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ActivitySearch struct{
    web.Activity
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
