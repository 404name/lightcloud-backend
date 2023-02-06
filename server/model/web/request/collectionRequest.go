package request

type CollectionRequest struct {
	Oid      uint   `json:"oid" form:"oid"`           // 社团id
	Cover    string `json:"cover" form:"cover"`       // 社团封面
	Name     string `json:"name" form:"name"`         // 社团名字
	Describe string `json:"describe" form:"describe"` // 社团描述
}
