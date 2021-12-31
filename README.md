create go module

run db with prepared docker-compose: docker compose up -d

add packages:
  go get entgo.io/ent/cmd/ent
  go get github.com/jackc/pgx/v4
  go get -u go.uber.org/zap

Create Your First Schema​ go run ```entgo.io/ent/cmd/ent init User```

add fields for schema/entity ```go generate ./ent```

for describe schema use ```go run entgo.io/ent/cmd/ent describe ./ent/schema```

if err missing go.sum entry for module providing package golang.org/x/sys/execabs, run in cli: ```go mod tidy``` https://stackoverflow.com/a/67203642

create db package with client connection and driver

if You need Upsert features: .OnConflictColumns() || .OnConflict().UpdateNewValues() etc...
  1) add it feature to ./ent/generator.go
  add it ```//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --idtype string --feature sql/upsert ./schema``` with --feature flag
  instead ```//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema```
  1) get latest ent ```go get entgo.io/ent@master```
  2) go generate ./ent



why pgx https://github.com/jackc/pgx
https://entgo.io/blog
https://medium.com/a-journey-with-go/go-ent-graph-based-orm-by-facebook-d9ba6d2290c6
https://betterprogramming.pub/implement-a-graphql-server-with-ent-and-gqlgen-in-go-8840f086b8a8

https://sqlc.dev
