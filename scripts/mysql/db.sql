create database if not exists tiktok;

use tiktok;

create table Users
(
    id int primary key auto_increment not null,
    auth_name varchar(64) unique not null ,
    nick_name varchar(64) not null,
    password varchar(128) not null
);
