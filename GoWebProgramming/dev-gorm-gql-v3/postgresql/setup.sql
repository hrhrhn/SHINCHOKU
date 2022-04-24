drop table if exists todos;
drop table if exists users;

create table users (
    id serial primary key,
    user_name varchar(255)
);


create table todos (
    id serial primary key,
    content text, 
    done boolean,
    user_id integer references users(id)
);