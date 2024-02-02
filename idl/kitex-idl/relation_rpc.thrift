namespace go follow_rpc

// 关注操作
struct RelationActionReq {
    1: i16 action_type; // 1-关注，2-取消关注
    2: i64 user_id;     // 用户id
    3: i64 follow_id;   // 关注者的id
}

struct RelationActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

// 是否关注
struct IsFollowReq {
    1: i64 user_id; // 用户id
    2: list<i64> follow_id; // 关注者id
}

struct IsFollowResp {
    1: i16 code;             // 状态码，0-成功，其他值-失败
    2: string msg;           // 返回状态描述
    3: list<bool> is_follow; // 是否关注
}

struct RelationListReq {
    1: i64 user_id;     // 用户id
    2: i64 owner_id;    // 所访问的用户id
    3: i16 action_type; // 操作码 1-关注列表 2-粉丝列表
}

struct RelationListResp {
    1: i16 code;             // 状态码，0-成功，其他值-失败
    2: string msg;           // 返回状态描述
    3: list<i64> user_id;    // 访问的用户所关注的用户的id
    4: list<bool> is_follow; // 是否关注
}

service FollowService {
    RelationActionResp RelationAction(1: RelationActionReq request);

    IsFollowResp IsFollow(1: IsFollowReq request);

    RelationListResp RelationList(1: RelationListReq request);
}
