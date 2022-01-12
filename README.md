create go module

run db with prepared docker-compose: docker compose up -d

https://sqlc.dev
https://pressly.github.io/goose

# sqlc
```goose -dir ./sql/migrations create init sql```
```sqlc generate```
``` goose -dir ./sql/migrations postgres "user=sandbox_user dbname=sandbox_db sslmode=disable password=1" up ```
