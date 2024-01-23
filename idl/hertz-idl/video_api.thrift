namespace go video_api

include "./user_api.thrift"

struct Video {
    1: i64 id;
    2: user_api.User author;
    3: string play_url;
    4: string cover_url;
    5: i64 favorite_count;
    6: i64 comment_count;
    7: bool is_favorite;
    8: string title;
}

struct FeedReq {
    1: i64 latest_time;
    2: optional string token;
}

struct FeedResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 next_time;
    4: list<Video> video_list;
}

struct ListReq {
    1: optional string token;
    2: i64 user_id;
}

struct ListResp {
    1: i32 status_code;
    2: string status_msg;
    3: list<Video> video_list;
}

service VideoService {
    FeedResp Feed(1: FeedReq request) (api.get = "/douyin/feed/")

    ListResp List(1: ListReq request) (api.get = "/douyin/publish/list/")
}
