package vo

type UserInfo struct {
	// UserId 用户id
	UserId uint `json:"id"`
	// NickName 用户昵称
	NickName string `json:"name"`
	// FollowCount 关注数
	FollowCount uint64 `json:"follow_count"`
	// FollowCount 粉丝数
	FollowerCount uint64 `json:"follower_count"`
	//IsFollow  true-已关注，false-未关注
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
}
