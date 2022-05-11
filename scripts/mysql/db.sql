create database if not exists tiktok;

use tiktok;

create table Users
(
    id        int primary key auto_increment not null,
    auth_name varchar(64) unique             not null,
    nick_name varchar(64)                    not null,
    password  varchar(128)                   not null,
    ts        timestamp                      not null
);

create table Videos
(
    id        int primary key auto_increment not null,
    author_id int                            not null,
    title     varchar(128)                   not null,
    ts        timestamp                      not null,
    foreign key (author_id) references Users (id)
);

create table Comments
(
    
);