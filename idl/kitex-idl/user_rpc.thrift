namespace go user_rpc

struct User {
    1: i64 id;              // 用户id
    2: string name;         // 用户名
    3: i64 follow_count;    // 关注总数
    4: i64 follower_count;  // 粉丝总数
    6: string avatar;       // 用户头像
    7: string background;   // 用户背景
    8: string signature;    // 用户签名
    9: i64 total_favorited; // 获赞数量
    10: i64 work_count;     // 作品数量
    11: i64 favorite_count; // 点赞数量
}

// 用户信息请求
struct UserListReq {
    1: list<i64> userinfo_id; // 需要查找的用户的id
    2: i64 user_id;           // 用户id 0-未登录
}

struct UserListResp {
    1: i16 code;         // 状态码，0-成功，其他值-失败
    2: string msg;       // 返回状态描述
    3: list<User> users; // 用户信息列表
}

// 账户操作请求
struct UserActionReq {
    1: i16 action_type; // 操作码 1-注册用户 2-登录用户
    2: string username; // 用户名
    3: string password; // 密码
}

struct UserActionResp {
    1: i16 code;   // 状态码，0-成功，其他值-失败
    2: string msg; // 返回状态描述
    3: i64 id;     // 用户id
}

service UserService {
    UserListResp UserList(1: UserListReq request);

    UserActionResp UserAction(1: UserActionReq request);
}
