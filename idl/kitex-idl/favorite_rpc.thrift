namespace go favorite_rpc

struct IsFavoriteReq {
    1: list<i64> user_id;
    2: list<i64> video_id; 
}

struct IsFavoriteResp {
    1: i16 status_code;
    2: string status_msg;
    3: list<bool> is_favorite;
}

service FavoriteService {
    IsFavoriteResp IsFavorite(1: IsFavoriteReq request);
}
