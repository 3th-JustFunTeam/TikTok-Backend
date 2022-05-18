SET NAMES utf8mb4;
drop database if exists tiktok;
create database tiktok;

use tiktok;

drop table if exists Users;
create table Users
(
    id         bigint unsigned auto_increment not null,
    auth_name  nvarchar(32) unique            not null,
    nick_name  nvarchar(32)                   not null,
    password   varchar(32)                    not null,
    ts         timestamp                      not null,
    update_ts  timestamp                      not null,
    is_deleted bool                           not null,
    primary key (`id`)
);

drop table if exists Follows;
create table Follows
(
    id           bigint unsigned auto_increment not null,
    user_id      bigint unsigned                not null comment '被关注者ID',
    following_id bigint unsigned                not null comment '关注者ID',
    ts           timestamp                      not null,
    is_deleted   bool                           not null,
    primary key (`id`),
    foreign key (`user_id`) references `Users` (`id`) on delete cascade on update cascade,
    foreign key (`following_id`) references `Users` (`id`) on delete cascade on update cascade
);

drop table if exists Videos;
create table Videos
(
    id         bigint unsigned auto_increment not null,
    author_id  bigint unsigned                not null,
    ts         timestamp                      not null,
    update_ts  timestamp                      not null,
    is_deleted bool                           not null,
    primary key (`id`),
    foreign key (`author_id`) references `Users` (`id`)
);

drop table if exists VideosData;
create table VideosData
(
    id          bigint unsigned auto_increment not null,
    video_id    bigint unsigned                not null,
    title       nvarchar(25)                   not null,
    description nvarchar(50)                   not null,
    video_path  varchar(64)                    not null,
    cover_path  varchar(64)                    not null,
    primary key (`id`)
);

drop table if exists VideoComments;
create table VideoComments
(
    id         bigint unsigned auto_increment not null,
    user_id    bigint unsigned                not null,
    video_id   bigint unsigned                not null,
    ts         timestamp                      not null,
    update_ts  timestamp                      not null,
    is_deleted bool                           not null,
    primary key (`id`),
    foreign key (`user_id`) references `Users` (`id`) on delete cascade on update cascade,
    foreign key (`video_id`) references `Videos` (`id`) on delete cascade on update cascade
);

drop table if exists VideoCommentsData;
create table VideoCommentsData
(
    id               bigint unsigned auto_increment not null,
    video_comment_id bigint unsigned                not null,
    comment_message  nvarchar(100)                  not null,
    primary key (`id`),
    foreign key (`video_comment_id`) references `VideoComments` (`id`) on delete cascade on update cascade
);

drop table if exists Favorites;
create table Favorites
(
    id         bigint unsigned auto_increment not null,
    video_id   bigint unsigned                not null,
    user_id    bigint unsigned                not null,
    ts         timestamp                      not null,
    update_ts  timestamp                      not null,
    is_deleted bool                           not null,
    primary key (`id`),
    foreign key (`user_id`) references `Users` (`id`) on delete cascade on update cascade,
    foreign key (`video_id`) references `Videos` (`id`) on delete cascade on update cascade
);
