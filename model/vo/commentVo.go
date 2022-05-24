package vo

type CommentVo struct {
	Id         uint     `json:"id"`
	User       UserInfo `json:"user"`
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}
