namespace go follow_rpc

struct User {
    1: i64 follow_count;   // 关注总数
    2: i64 follower_count; // 粉丝总数
    3: bool is_follow;     // 是否关注 不能关注自己
}

struct RelationActionReq {
    1: i16 action_type; // 1-关注，2-取消关注
    2: i64 user_id;     // 用户id
    3: i64 follow_id;   // 关注者的id
}

struct RelationActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

struct UserListReq {
    1: list<i64> userinfo_id; // 需要查找的用户的id
    2: i64 user_id;           // 用户id 0-未登录
}

struct UserListResp {
    1: i16 code;        // 状态码，0-成功，其他值-失败
    2: string msg;      // 返回状态描述
    3: list<User> user; // 用户关注信息列表
}

struct UserActionReq {
    1: i64 user_id; // 用户id
}

struct UserActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
}

service FollowService {
    RelationActionResp RelationAction(1: RelationActionReq request);

    UserListResp UserList(1: UserListReq request);

    UserActionResp UserAction(1: UserActionReq request);
}
