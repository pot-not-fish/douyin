namespace go favorite_rpc

// 点赞视频操作请求
struct FavoriteActionReq {
    1: i16 action_type; // 1-点赞，2-取消点赞
    2: i64 user_id;  // 用户id
    3: i64 video_id; // 视频id
}

struct FavoriteActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

// 查看视频是否点赞
struct IsFavoriteReq {
    1: list<i64> user_id;  // 用户id列表
    2: list<i64> video_id; // 视频id列表
}

struct IsFavoriteResp {
    1: i16 code;               // 状态码，0-成功，其他值-失败
    2: string msg;             // 返回状态描述
    3: list<bool> is_favorite; // 是否点赞
}

// 查看用户点赞视频
struct FavoriteVideoReq {
    1: i64 user_id;  // 用户id
    2: i64 owner_id; // 所访问的用户id
}

struct FavoriteVideoResp {
    1: i16 code;               // 状态码，0-成功，其他值-失败
    2: string msg;             // 返回状态描述
    3: list<i64> video_id;     // 点赞的视频id列表
    4: list<bool> is_favorite; // 用户是否点赞
}

service FavoriteService {
    FavoriteActionResp FavoriteAction(1: FavoriteActionReq request);

    IsFavoriteResp IsFavorite(1: IsFavoriteReq request);
}
