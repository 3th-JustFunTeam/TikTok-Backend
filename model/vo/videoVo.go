package vo

// VideoVo 视频接口需要的数据
type VideoVo struct {
	// Id 视频id
	Id uint `json:"id"`
	// Author 视频作者
	Author UserInfo `json:"author"`
	// PlayUrl 视频播放地址
	PlayUrl string `json:"play_url"`
	// CoverUrl 视频封面地址
	CoverUrl string `json:"cover_url"`
	// favoriteCount 视频的点赞总数
	FavoriteCount uint64 `json:"favorite_count"`
	// CommentCount 视频的评论总数
	CommentCount uint64 `json:"comment_count"`
	// IsFavorite 是否对视频点赞
	IsFavorite bool `json:"is_favorite"`
	// Title 视频标题
	Title string `json:"title"`
}
