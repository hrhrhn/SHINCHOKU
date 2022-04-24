view SQL
```
(base) htakagi@Mac ~ % psql -U gwp -d gwp                         

psql (14.2)
Type "help" for help.

gwp=> select * from posts;
 id |        content         |         author          
----+------------------------+-------------------------
  2 | Hello World in gwp DB! | Sau Sheong コメント待ち
  3 | Updated post           | Sau Sheong
(2 rows)

```

POST
```
curl -i -X POST -H "Content-Type: application/json"  -d '{"content":"My first post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/
```


GET (id=1)
```
curl -i -X GET http://127.0.0.1:8080/post/1
```

PUT (id=1)
```
curl -i -X PUT -H "Content-Type: application/json"  -d '{"content":"Updated post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/1
```

DELETE
```
curl -i -X DELETE http://127.0.0.1:8080/post/1
```