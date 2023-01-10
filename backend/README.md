# backend

1. Setup Postgres
```sh
$ docker run \
-v `pwd`/init.sql:/docker-entrypoint-initdb.d/init.sql:ro \
--name pg-gi \
-e POSTGRES_USER=backend_user \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=gi \
-p 4321:5432 \
-d postgres
```

2. Start Server
```sh
go run main.go
```
