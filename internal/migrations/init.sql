create database if not exists tyres;

use tyres;

create table if not exists users (
    id int primary key auto_increment,
    username varchar(255) not null,
    hashed_password varchar(255) not null,
    email varchar(255) not null,
    phone varchar(255) not null
);

create table if not exists posts (
    id int primary key auto_increment,
    title varchar(255) not null,
    content text not null,
    img varchar(255) not null,
    price int not null
);
