namespace go publish

enum Code {
     Success = 0
     Error = 1
}

struct User {
    1: i64 id // 用户id
    2: string name // 用户名称
    3: i64 follow_count // 关注总数
    4: i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注,false-未关注
}

struct Video {
    1: i64 id //视频唯一标识
	2: User author //视频作者信息
	3: string play_url //视频播放地址
	4: string cover_url //视频封面地址
	5: i64 favorite_count //视频的点赞总数
	6: i64 comment_count //视频的评论总数
	7: bool is_favorite //true-已点赞，false-未点赞
	8: string title //视频标题
}

struct DouyinPublishActionRequest {
    1: string token // 用户鉴权token
    2: byte data // 视频数据
    3: string title // 视频标题
}

struct DouyinPublishActionResponse {
    1: i64 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
}

struct DouyinPublishListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}

struct DouyinPublishListResponse {
    1: i64 status_code // 状态码，0-成功，其他值-失败
    2: string status_msg // 返回状态描述
    3: list<Video> video_list // 用户发布的视频列表
}

service PublishService {
    DouyinPublishActionResponse PublishAction(1: DouyinPublishActionRequest req)(api.post="/douyin/publish/action/")
    DouyinPublishListResponse PublishList(1: DouyinPublishListRequest req)(api.get="/douyin/publish/list/")
}