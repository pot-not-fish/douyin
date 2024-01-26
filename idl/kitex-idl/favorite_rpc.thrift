namespace go favorite_rpc

// 点赞视频操作请求
struct FavoriteActionReq {
    1: i16 action_type; // 1-点赞，2-取消点赞
    2: string user_id;  // 用户id
    3: string video_id; // 视频id
}

struct FavoriteActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

struct Video {
    1: i64 favorite_count; // 视频点赞总数
    2: bool is_favorite;   // 是否点赞
}

// 视频点赞相关信息请求
struct VideoListReq {
    1: list<i64> video_id; // 视频id
    2: i64 user_id;        // 用户id
}

struct VideoListResp {
    1: i16 code;           // 状态码，0-成功，其他值-失败
    2: string msg;         // 返回状态描述
    3: list<Video> videos; // 视频点赞相关信息
}

struct User {
    1: i64 total_favorited; // 获赞数量
    2: i64 favorite_count;  // 点赞数量
}

// 用户点赞相关信息请求
struct UserListReq {
    1: list<i64> userinfo_id; // 需要查找的用户的id
    2: i64 user_id;           // 用户id 0-未登录
}

struct UserListResp {
    1: i16 code;           // 状态码，0-成功，其他值-失败
    2: string msg;         // 返回状态描述
}

struct VideoActionReq {
    1: i64 video_id; // 视频id
}

struct VideoActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

struct UserActionReq {
    1: i64 user_id; // 用户id
}

struct UserActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

service FavoriteService {
    FavoriteActionResp FavoriteAction(1: FavoriteActionReq request);

    VideoListResp VideoList(1: VideoListReq request);

    UserListResp UserList(1: UserListReq request);

    UserActionResp UserAction(1: UserActionReq request);

    VideoActionResp UserAction(1: VideoActionReq request);
}
