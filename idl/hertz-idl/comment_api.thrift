namespace go comment_api

include "./user_api.thrift"

// 评论操作接口
struct CommentActionReq {
    1: required string token;        // 用户token
    2: required i64 video_id;        // 视频id
    3: required i32 action_type;     // 操作码 1-发布评论，2-删除评论
    4: optional string comment_text; // 评论内容
    5: optional i64 comment_id;      // 要删除的评论的id
}

struct CommentActionResp {
    1: i32 status_code;   // 状态码，0-成功，其他值-失败
    2: string status_msg; // 返回状态描述
    3: Comment comment;   // 返回评论的内容
}

struct Comment {
    1: i64 id;
    2: user_api.User user;
    3: string content;
    4: string create_date;
}

struct CommentListReq {
    1: optional string token;
    2: required i64 video_id;
}

struct CommentListResp {
    1: i32 status_code;
    2: string status_msg;
    3: list<Comment> comment_list; 
}

service CommentService {
    CommentActionResp CommentAction(1: CommentActionReq request) (api.post = "/douyin/comment/action/")

    CommentListResp CommentList(1: CommentListReq request) (api.get = "/douyin/comment/list/")
}
