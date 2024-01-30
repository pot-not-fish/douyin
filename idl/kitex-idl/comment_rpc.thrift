namespace go comment_rpc

struct Comment {
    1: i64 id;             // 发布id
    2: string content;     // 发布内容
    3: string create_date; // 发布时间
}

// 评论操作请求
struct CommentActionReq {
    1: i16 action_type;              // 1-评论，2-删除评论
    2: i64 user_id;                  // 用户id
    3: i64 video_id;                 // 视频id
    4: optional string comment_text; // 评论内容
}

struct CommentActionResp {
    1: i16 code;        // 状态码，0-成功，其他值-失败
    2: string msg;      // 返回状态描述
    3: Comment comment; // 返回发布的评论内容
}

// 视频评论列表请求
struct CommentListReq {
    1: i64 video_id; // 视频id
}

struct CommentListResp {
    1: i16 code;              // 状态码，0-成功，其他值-失败
    2: string msg;            // 返回状态描述
    3: list<Comment> comment; // 返回发布的评论内容
}

service CommentService {
    CommentActionResp CommentAction(1: CommentActionReq request);

    CommentListResp CommentList(1: CommentListReq request);
}
