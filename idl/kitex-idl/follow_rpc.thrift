namespace go follow_rpc

struct IsFollowReq {
    1: list<i64> to_user_id;
    2: list<i64> from_user_id;
}

struct IsFollowResp {
    1: i16 status_code;
    2: string status_msg;
    3: list<bool> is_follow;
}

service FollowService {
    IsFollowResp IsFollow(1: IsFollowReq request);
}
 