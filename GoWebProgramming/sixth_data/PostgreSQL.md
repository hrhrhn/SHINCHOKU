前提  
`.zshrc`でPostgreSQLのパスを通す
```
export PATH=$PATH:/Applications/Postgres.app/Contents/Versions/14/bin
```

<br>

<!-- PostgreSQLに入る
```
$ psql
``` -->

<br>

ユーザーの作成
```
$ createuser -P -d gwp

Enter password for new role: 
Enter it again: 
```
パスワードも同名とする

<br>

同名のデータベースを作成
```
$ createdb gwp
```

`setup.sql`を実行
```
$ psql -U gwp -f setup.sql -d gwp

CREATE TABLE
```
---
コマンド
```
gwp=# \l
                                  List of databases
   Name    |  Owner   | Encoding |   Collate   |    Ctype    |   Access privileges   
-----------+----------+----------+-------------+-------------+-----------------------
 chitchat  | htakagi  | UTF8     | en_US.UTF-8 | en_US.UTF-8 | 
 gwp       | htakagi  | UTF8     | en_US.UTF-8 | en_US.UTF-8 | 
 htakagi   | htakagi  | UTF8     | en_US.UTF-8 | en_US.UTF-8 | 
 postgres  | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | 
 template0 | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/postgres          +
           |          |          |             |             | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/postgres          +
           |          |          |             |             | postgres=CTc/postgres
(6 rows)

gwp=# \c gwp
You are now connected to database "gwp" as user "htakagi".


gwp-# \d
            List of relations
 Schema |     Name     |   Type   | Owner 
--------+--------------+----------+-------
 public | posts        | table    | gwp
 public | posts_id_seq | sequence | gwp
(2 rows)


gwp-# \d posts
                                    Table "public.posts"
 Column  |          Type          | Collation | Nullable |              Default              
---------+------------------------+-----------+----------+-----------------------------------
 id      | integer                |           | not null | nextval('posts_id_seq'::regclass)
 content | text                   |           |          | 
 author  | character varying(255) |           |          | 
Indexes:
    "posts_pkey" PRIMARY KEY, btree (id)


gwp=# select * from posts;
 id | content | author 
----+---------+--------
(0 rows)
```