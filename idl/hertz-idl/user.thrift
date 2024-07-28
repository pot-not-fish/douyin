namespace go user_api

struct User {
    1: i64 id;
    2: i64 follow_count;
    3: i64 fans_count;
    4: bool is_follow;
    5: i64 work_count;
    6: i64 favorite_count;
    7: i64 total_favorited;
}

struct RegisterReq {
}

struct RegisterResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
}

struct UserinfoReq {
    1: required i64 user_id;
}

struct UserinfoResp {
    1: i32 status_code;
    2: string status_msg;
    3: User user;
}

service UserService {
    RegisterResp Register(1: RegisterReq request) (api.post = "/douyin/user/register");

    UserinfoResp Userinfo(1: UserinfoReq request) (api.get = "/douyin/user");
}
