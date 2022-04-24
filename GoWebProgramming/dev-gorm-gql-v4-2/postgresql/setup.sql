drop table if exists details_v4;
drop table if exists todos_v4;
drop table if exists users_v4;
drop table if exists parents_v4;

create table parents_v4 (
    id serial primary key,
    parent_name varchar(255)
);

create table users_v4 (
    id serial primary key,
    user_name varchar(255),
    parent_id integer references parents_v4(id)
        on delete set null
        on update cascade
        null
);

create table todos_v4 (
    id serial primary key,
    content text, 
    done boolean,
    user_id integer references users_v4(id) on delete cascade on update cascade
);

create table details_v4 (
    id serial primary key,
    content text, 
    done boolean,
    todo_id integer references todos_v4(id),
    user_id integer references users_v4(id) on delete cascade on update cascade
);
