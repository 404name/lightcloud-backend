package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ActivityRecordSearch struct{
    web.ActivityRecord
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartScore  *int  `json:"startScore" form:"startScore"`
    EndScore  *int  `json:"endScore" form:"endScore"`
    request.PageInfo
}
