create database tyres;

use tyres;

create table users {
    id int primary key auto_increment,
    username varchar(255) not null,
    hashed_password varchar(255) not null,
    email varchar(255) not null,
    phone varchar(255) not null,
    is_admin boolean not null,
    created_at varchar(255) not null,
    updated_at varchar(255) not null
};

create table posts {
    id int primary key auto_increment,
    title varchar(255) not null,
    content text not null,
    img varchar(255) not null,
    price int not null,
    created_at varchar(255) not null,
    updated_at varchar(255) not null
};
