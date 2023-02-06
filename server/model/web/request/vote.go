package request

type VoteRequest struct {
	CommentId int  `json:"commentId" form:"commentId"`
	Approval  bool `json:"approval" form:"approval"`
}
