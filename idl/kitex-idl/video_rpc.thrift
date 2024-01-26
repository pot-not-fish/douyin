namespace go video_rpc

// 点赞视频操作请求
struct FavoriteActionReq {
    1: i16 action_type; // 1-点赞，2-取消点赞
    2: i64 user_id;  // 用户id
    3: i64 video_id; // 视频id
}

// 点赞视频操作返回
struct FavoriteActionResp {
    1: i16 code;       // 状态码，0-成功，其他值-失败
    2: string msg;     // 返回状态描述
}

struct Video {
    1: i64 id;             // 视频id
    2: string play_url;    // 播放地址
    3: string cover_url;   // 视频封面
    4: string title;       // 视频标题
    5: i64 favorite_count; // 视频点赞总数
    6: bool is_favorite;   // 是否点赞
}

struct User {
    1: i64 id;              // 用户id
    2: i64 work_count;      // 作品数量
}

// 视频feed流请求
struct FeedReq {
    1: i64 user_id; // 用户id 0-未登录
}

struct FeedResp {
    1: i16 code;           // 状态码，0-成功，其他值-失败
    2: string msg;         // 返回状态描述
    3: list<Video> videos; // 视频列表
}

// 用户点赞视频列表请求
struct FavoriteListReq {
    1: i64 user_id;          // 用户id 0-未登录
    2: i64 favorite_user_id; // 点赞视频列表所属的用户id 
}

struct FavoriteListResp {
    1: i16 code;           // 状态码，0-成功，其他值-失败
    2: string msg;         // 返回状态描述
    3: list<Video> videos; // 视频列表
}

// 用户已发布视频列表请求
struct PublishListReq {
    1: i64 user_id;         // 用户id 0-未登录
    2: i64 publish_user_id; // 发布视频列表所属的用户id 
}

struct PublishListReq {
    1: i16 code;           // 状态码，0-成功，其他值-失败
    2: string msg;         // 返回状态描述
    3: list<Video> videos; // 视频列表
}

// 视频操作请求
// 如果有删除视频的功能，可以将视频标题改为可选，添加操作码 1-发布，2-删除
struct VideoActionReq {
    1: i64 user_id;  // 用户id 0-未登录
    2: string title; // 视频标题
}

struct VideoActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

// 视频列表请求
struct VideoListReq {
    1: list<i64> video_id; // 视频id
}

struct VideoListResp {
    1: i16 code;           // 状态码，0-成功，其他值-失败
    2: string msg;         // 返回状态描述
    3: list<Video> videos; // 视频列表
}

service VideoService {
    FavoriteActionResp FavoriteAction(1: FavoriteActionReq request);

    FeedResp Feed(1: FeedReq request);

    FavoriteListResp FavoriteList(1: FavoriteListReq request);

    PublishListReq PublishList(1: PublishListReq request);

    VideoActionResp Video(1: VideoActionReq request);

    VideoListResp VideoList(1: VideoListReq request);
}