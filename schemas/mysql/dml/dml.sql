create table users (
    id int primary key auto_increment,
    username varchar(255) not null
);
insert into users (username) values ('user1');

create table todos (
    id int primary key auto_increment,
    user_id int not null,
    title varchar(255) not null
);

insert into todos (user_id, title) values (1, 'todo1');
