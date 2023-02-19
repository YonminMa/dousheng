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
    6: string avatar //用户头像
    7: string background_image //用户个人页顶部大图
    8: string signature //个人简介
    9: i64 total_favorited //获赞数量
    10: i64 work_count //作品数量
    11: i64 favorite_count //点赞数量
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

struct PublishActionRequest {
    1: string token // 用户鉴权token
    2: byte data // 视频数据
    3: string title // 视频标题
}

struct PublishActionResponse {
    1: Code status_code // 状态码
    2: string status_msg // 返回状态描述
}

struct PublishListRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}

struct PublishListResponse {
    1: Code status_code // 状态码
    2: string status_msg // 返回状态描述
    3: list<Video> video_list // 用户发布的视频列表
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req)(api.post="/douyin/publish/action/")
    PublishListResponse PublishList(1: PublishListRequest req)(api.get="/douyin/publish/list/")
}