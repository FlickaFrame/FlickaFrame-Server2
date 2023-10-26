// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	ID       int64  `json:"id"`
	Phone    string `json:"phone"`
	NickName string `json:"nick_name"`
	Sex      int64  `json:"sex"`
	AvtarUrl string `json:"avatar"`
	Info     string `json:"info"`
}

type RegisterReq struct {
	Phone    string `json:"phone" validate:"e164,required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nick_name,option"`
}

type RegisterResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"user_info"`
}

type VideoUserInfo struct {
	NickName  string `json:"nickName"`  // 昵称
	AvatarUrl string `json:"avatarUrl"` // 头像地址
}

type Video struct {
	ID            int64         `json:"id"`         // 视频ID
	Title         string        `json:"title"`      // 视频标题
	PlayUrl       string        `json:"playUrl"`    // 视频播放地址
	ThumbUrl      string        `json:"thumbUrl"`   // 视频封面地址
	FavoriteCount int64         `json:"favNum"`     // 点赞数
	CommentCount  int64         `json:"commentNum"` // 评论数
	ShareNum      int64         `json:"shareNum"`   // 分享数
	CreatedAt     string        `json:"createdAt"`  // 视频创建时间(毫秒时间戳)
	IsFav         bool          `json:"isFav"`      // 当前用户是否已点赞
	IsFollow      bool          `json:"isFollow"`   // 当前用户是否已关注该用户
	Tags          []string      `json:"tags"`       // 视频标签
	Author        VideoUserInfo `json:"author"`     // 作者信息
}

type FeedReq struct {
	LatestTime int64  `json:"latest_time,optional" form:"latestTime,optional"` // 最新视频时间(毫秒时间戳)
	Limit      int    `json:"limit,optional" form:"limit,optional"`            // 请求数量
	Token      string `json:"token,optional" form:"token,optional"`            // 登录token
	AuthorID   uint   `json:"author_id,optional" form:"authorID,optional"`     // 作者ID
}

type FeedResp struct {
	VideoList []*Video `json:"video_list"`
	NextTime  int64    `json:"next_time"` // 下次请求时间(毫秒时间戳)
	Length    int      `json:"length"`    // 视频列表长度
}
