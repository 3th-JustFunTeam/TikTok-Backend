# API

## Basic API

### /douyin/feed/ - 视频流接口

Method: GET

Definition:

```protobuf
syntax = "proto2";
package douyin.core;

message douyin_feed_request {
  optional int64 latest_time = 1; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

message douyin_feed_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated Video video_list = 3; // 视频列表
  optional int64 next_time = 4; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

message Video {
  required int64 id = 1; // 视频唯一标识
  required User author = 2; // 视频作者信息
  required string play_url = 3; // 视频播放地址
  required string cover_url = 4; // 视频封面地址
  required int64 favorite_count = 5; // 视频的点赞总数
  required int64 comment_count = 6; // 视频的评论总数
  required bool is_favorite = 7; // true-已点赞，false-未点赞
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
}
```

## /douyin/user/register/ - 用户注册接口

Method: POST


Definition:

```protobuf
syntax = "proto2";
package douyin.core;

message douyin_user_register_request {
  required string username = 1; // 注册用户名，最长32个字符
  required string password = 2; // 密码，最长32个字符
}

message douyin_user_register_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required int64 user_id = 3; // 用户id
  required string token = 4; // 用户鉴权token
}
```

## /douyin/user/login/ - 用户登录接口

Method: POST

Definition:

```protobuf
syntax = "proto2";
package douyin.core;

message douyin_user_login_request {
  required string username = 1; // 登录用户名
  required string password = 2; // 登录密码
}

message douyin_user_login_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required int64 user_id = 3; // 用户id
  required string token = 4; // 用户鉴权token
}
```

## /douyin/user/ - 用户信息

Method: POST

Definition:

```protobuf
syntax = "proto2";
package douyin.core;

message douyin_user_request {
  required int64 user_id = 1; // 用户id
  required string token = 2; // 用户鉴权token
}

message douyin_user_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required User user = 3; // 用户信息
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
}
```

## /douyin/publish/action/ - 视频投稿

Method: POST

Definition:

```protobuf
syntax = "proto2";
package douyin.core;

message douyin_publish_action_request {
  required int64 user_id = 1; // 用户id
  required string token = 2; // 用户鉴权token
  required bytes data = 3; // 视频数据
}

message douyin_publish_action_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
}

```

## /douyin/publish/list/ - 发布列表

Method: GET

Definition:

```protobuf
syntax = "proto2";
package douyin.core;

message douyin_publish_list_request {
  required int64 user_id = 1; // 用户id
  required string token = 2; // 用户鉴权token
}

message douyin_publish_list_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated Video video_list = 3; // 用户发布的视频列表
}

message Video {
  required int64 id = 1; // 视频唯一标识
  required User author = 2; // 视频作者信息
  required string play_url = 3; // 视频播放地址
  required string cover_url = 4; // 视频封面地址
  required int64 favorite_count = 5; // 视频的点赞总数
  required int64 comment_count = 6; // 视频的评论总数
  required bool is_favorite = 7; // true-已点赞，false-未点赞
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
}
```