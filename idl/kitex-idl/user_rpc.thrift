namespace go user_rpc

struct User {
    1: i64 user_id;
    2: string name;
    3: i64 follow_count;
    4: i64 follower_count;
    6: string avatar;
    7: string background;
    8: string signature;
    9: i64 total_favorited;
    10: i64 work_count;
    11: i64 favorite_count;
}

struct RetriveUserReq {
    1: list<i64> user_id;
}

struct RetriveUserResp {
    1: i16 status_code;
    2: string status_msg;
    3: list<User> user;
}

service UserService {
    RetriveUserResp Userinfo(1: RetriveUserReq request);
}
