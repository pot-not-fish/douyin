namespace go video_rpc

struct Video {
    1: i64 id;             // 视频id
    2: string play_url;    // 播放地址
    3: string cover_url;   // 视频封面
    4: string title;       // 视频标题
    5: i64 favorite_count; // 视频点赞总数
    6: i64 comment_count;  // 评论数量
}

// 用户列表请求
struct VideoListReq {
    1: i64 user_id;     // 用户id 0-未登录
    2: i64 owner_id;    // 点赞视频列表所属的用户id 0-feed流请求
    3: i64 action_type; // 操作数 0-feed流请求 1-返回查看用户的点赞列表 2-返回查看用户的发布视频列表
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

// 视频列表信息请求
struct VideoInfoReq {
    1: list<i64> video_id; // 视频id
}

struct VideoListResp {
    1: i16 code;           // 状态码，0-成功，其他值-失败
    2: string msg;         // 返回状态描述
    3: list<Video> videos; // 视频列表
}

service VideoService {
    VideoListResp VideoList(1: VideoListReq request);

    VideoListResp VideoInfo(1: VideoInfoReq request);
    
    VideoActionResp VideoAction(1: VideoActionReq request);
}