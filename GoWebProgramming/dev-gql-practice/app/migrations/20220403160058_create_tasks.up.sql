DROP TRIGGER update_tri ON todos;
DROP FUNCTION set_update_time();

DROP TABLE IF EXISTS todos;

create table todos (
    id serial primary key,
    title varchar(255) not null,
    notes text, 
    completed boolean not null default false,
    due timestamp null default null,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp null default null
);

create function set_update_time() returns opaque as '
  begin
    new.update_at := ''now'';
    return new;
  end;
' language 'plpgsql';

create trigger update_tri before update on todos for each row
  execute procedure set_update_time();