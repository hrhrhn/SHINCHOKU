PostgresQL

user=gwp dbname=gwp password=gwp

```
psql -U gwp -f ./postgresql/setup.sql -d gwp
```


Run Server
```
cd src/
go run server.go
```