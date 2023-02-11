namespace go relation

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

struct RelationActionRequest {
    1: string token //用户鉴权token
    2: i64 to_user_id //对方用户id
    3: i32 action_type //1-关注，2-取消关注
}

struct RelationActionResponse {
    1: Code status_code // 状态码
    2: string status_msg // 状态描述
}

struct FollowListRequest {
    1: i64 user_id //用户id
    2: string token //用户鉴权token
}

struct FollowListResponse {
    1: Code status_code // 状态码
    2: string status_msg // 状态描述
    3: list<User> user_list //用户信息列表
}

struct FollowerListRequest {
    1: i64 user_id //用户id
    2: string token //用户鉴权token
}

struct FollowerListResponse {
    1: Code status_code // 状态码
    2: string status_msg // 状态描述
    3: list<User> user_list //用户信息列表
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest req)(api.post="/douyin/relation/action/")
    FollowListResponse FollowList(1: FollowListRequest req)(api.get="/douyin/relation/follow/list/")
    FollowerListResponse FollowerList(1: FollowerListRequest req)(api.get="/douyin/relation/follower/list/")
}