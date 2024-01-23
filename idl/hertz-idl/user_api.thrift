namespace go user_api

// id改string
// total_favorited改i64

struct User {
    1: i64 id;
    2: string name;
    3: i64 follow_count;
    4: i64 follower_count;
    5: bool is_follow;
    6: string avatar;
    7: string background_image;
    8: string signature;
    9: i64 total_favorited;
    10: i64 work_count;
    11: i64 favorite_count;
}

// 用户注册接口
struct RegisterReq {
    1: required string username;
    2: required string password;
}

struct RegisterResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
    4: string token;
}

// 用户信息接口
struct UserinfoReq {
    1: required i64 user_id;
    2: optional string token;
}

struct UserinfoResp {
    1: i32 status_code;
    2: string status_msg;
    3: User user;
}

service UserService {
    RegisterResp Register(1: RegisterReq request) (api.post = "/douyin/user/register/");

    UserinfoResp Userinfo(1: UserinfoReq request) (api.get = "/douyin/user/");
}