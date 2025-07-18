## Reel Views Backend
Апи развернуто на хостинге. 

Base URL: https://reel-views.onrender.com

Документация к ендпоинтам находится в ./docs/swagger.yaml

Либо https://reel-views.onrender.com/swagger/index.html

Для **локального** развертывания:
```cmd
docker run --env=POSTGRES_USER=postgres --env=POSTGRES_PASSWORD=postgres --env=POSTGRES_DB=postgres --network=bridge -p 5432:5432 -d postgres
docker run --env=REDIS_VERSION=7.4.3 --volume=/data --network=bridge --workdir=/data -p 6379:6379 -d redis:7
go build -tags netgo -ldflags '-s -w' -o  main ./cmd/app/main.go
```
