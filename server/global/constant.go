package global

// 文章类别
const (
	CommentTypeArticle  = iota // 文章
	CommentTypeActivity        // 活动
	CommentTypeUser            // 用户/社团
	CommentTypeMessage         // 留言
)

// 动态跳转类别
const (
	MomentsTypeArticle      = iota // 文章
	MomentsTypeActivity            // 活动
	MomentsTypeOrganization        // 用户/社团
)
