namespace go user

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
    6: optional string avatar //用户头像
    7: optional string background_image //用户个人页顶部大图
    8: optional string signature //个人简介
    9: optional i64 total_favorited //获赞数量
    10: optional i64 work_count //作品数量
    11: optional i64 favorite_count //点赞数量
}

struct UserLoginRequest {
    1: string username // 登录用户名
    2: string password // 登录密码
}

struct UserLoginResponse {
    1: Code status_code // 状态码
    2: string status_msg // 状态描述
    3: i64 user_id // 用户id
    4: string token // 用户鉴权token
}

struct UserRegisterRequest {
    1: string username // 注册用户名，最长32个字符
    2: string password // 密码，最长32个字符
}

struct UserRegisterResponse {
    1: Code status_code // 状态码
    2: string status_msg // 状态描述
    3: i64 user_id // 用户id
    4: string token // 用户鉴权token
}


struct UserInfoRequest {
    1: i64 user_id // 用户id
    2: string token // 用户鉴权token
}

struct UserInfoResponse {
    1: Code status_code // 状态码
    2: string status_msg // 状态描述
    3: User user // 用户信息
}

service UserService {
    UserRegisterResponse RegisterUser(1: UserRegisterRequest req)(api.post="/douyin/user/register/")
    UserLoginResponse LoginUser(1: UserLoginRequest req)(api.post="/douyin/user/login/")
    UserInfoResponse UserInfo(1: UserInfoRequest req)(api.get="/douyin/user/")
}
