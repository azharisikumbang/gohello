drop table if exists user_groups;
create table user_groups (
    id int primary key auto_increment,
    name varchar(255) unique not null,
    description text default null,
    number int default null,
    created_at datetime default CURRENT_TIMESTAMP
);

drop table if exists users;
create table users (
    id int primary key auto_increment,
    username varchar(255) unique not null,
    password varchar(255) not null,
    user_group_id int default 2, -- GUEST
    created_at datetime default CURRENT_TIMESTAMP,
    foreign key(user_group_id) 
        references user_groups(id) 
        on delete set null 
);

create index uk_users_username on users(username);
create index uk_users_user_group_id on users(user_group_id);
create index uk_users_groups_name on user_groups(name);

insert into user_groups (name, number) values 
    ('ROOT', 1),
    ('GUEST', 99999);

-- root root
insert into users (username, password, user_group_id) values 
    ('root', '$2a$12$wGDDmxjGYgKN92QE4Ii6ZOwxSslsV15gYYfAMmlgXUvqq.X98Zr8S', 1);
