namespace go comment_api

include "./user_api.thrift"

struct CommentActionReq {
    1: required string token;
    2: required i64 video_id;
    3: required i32 action_type;
    4: optional string comment_text;
    5: optional i64 comment_id;
}

struct CommentActionResp {
    1: i32 status_code;
    2: string status_msg;
    3: Comment comment;
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
