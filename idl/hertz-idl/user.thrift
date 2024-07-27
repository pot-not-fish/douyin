namespace go user_api

struct User {
    1: string id;
    2: string name;
    3: i64 follow_count;
    4: i64 follower_count;
    5: bool is_follow;
    7: string signature;
    8: i64 total_favorited;
    9: i64 work_count;
    10: i64 favorite_count;
}

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

struct UserinfoReq {
    1: required i64 user_id;
    2: optional string token;
}

struct UserinfoResp {
    1: i32 status_code;
    2: string status_msg;
    3: User user;
}

struct LoginReq {
    1: required string username;
    2: required string password;
}

struct LoginResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
    4: string token;
}

service UserService {
    RegisterResp Register(1: RegisterReq request) (api.post = "/douyin/user/register");

    UserinfoResp Userinfo(1: UserinfoReq request) (api.get = "/douyin/user");

    LoginResp Login(1: LoginReq request) (api.post = "/douyin/user/login");
}
