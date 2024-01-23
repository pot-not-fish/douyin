namespace go favorite_api

include "./video_api.thrift"

struct FavoriteActionReq {
    1: required string token;
    2: required i64 video_id;
    3: required i32 action_type;
}

struct FavoriteActionResp {
    1: i32 status_code;
    2: string status_msg;
}

struct FavoriteListReq {
    1: optional string token;
    2: required i64 user_id;
}

struct FavoriteListResp {
    1: i32 status_code;
    2: string status_msg;
    3: list<video_api.Video> video_list;
}

service FavoriteService {
    FavoriteActionResp FavoriteAction(1: FavoriteActionReq request) (api.post = "/douyin/favorite/action/")

    FavoriteListResp FavoriteList(1: FavoriteListReq request) (api.get = "/douyin/favorite/list/")
}
