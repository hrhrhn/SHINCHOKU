```
$ tree backend
.
├── .gitignore
├── Makefile
├── README.md
├── app
│   ├── .realize.yaml
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── db
│   └── mysql
│       └── my.cnf
└── docker-compose.yml


(base) htakagi@Mac dev-gql-practice % make

docker-compose up --build -d
Building app
[+] Building 2.3s (15/15) FINISHED
...
Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
Starting gql-psql-db ... done
Recreating gql-psql-backend ... done


(base) htakagi@Mac dev-gql-practice % make migrate-up

docker-compose exec app migrate --source file://migrations --database postgres://gwp:gwp@db:5432/gwp?sslmode=disable up
20220403160058/u create_tasks (16.2506ms)


(base) htakagi@Mac dev-gql-practice % make start

docker-compose exec app realize start --run
[15:58:05][APP] : Watching 1 file/s 1 folder/s
[15:58:05][APP] : Install started
[15:58:06][APP] : Install completed in 0.940 s
[15:58:06][APP] : Running..
[15:58:06][APP] : ⇨ http server started on [::]:8080





(base) htakagi@Mac dev-gql-practice % cd app
(base) htakagi@Mac app % go get github.com/99designs/gqlgen                             
(base) htakagi@Mac app % docker-compose exec app go run github.com/99designs/gqlgen init

(base) htakagi@Mac dev-gql-practice % make generate
docker-compose exec app go generate ./...
```