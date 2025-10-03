create table users (
    id int primary key auto_increment,
    username varchar(255) unique not null,
    password varchar(255) not null,
    created_at datetime default CURRENT_TIMESTAMP
)