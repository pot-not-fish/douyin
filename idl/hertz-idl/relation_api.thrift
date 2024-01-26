namespace go follow_api

include "./user_api.thrift"

struct RelationActionReq {
    1: required string token;
    2: required i64 to_user_id;
    3: required i32 action_type;
}

struct RelationActionResp {
    1: i32 status_code;
    2: string status_msg;
}

struct FollowListReq {
    1: required i64 user_id;
    2: optional string token;
}

struct FollowListResp {
    1: i32 status_code;
    2: string status_msg;
    3: list<user_api.User> user_list;
}

struct FollowerListReq {
    1: required i64 user_id;
    2: optional string token;
}

struct FollowerListResp {
    1: i16 status_code;
    2: string status_msg;
    3: list<user_api.User> user_list;
}

service RelationService {
    RelationActionResp RelationAction(1: RelationActionReq request) (api.post = "/douyin/relation/action/")

    FollowListResp RelationFollow(1: FollowListReq request) (api.get = "/douyin/relation/follow/list/")

    FollowerListResp RelationFollower(1: FollowerListReq request) (api.get = "/douyin/relation/follower/list/")
}
