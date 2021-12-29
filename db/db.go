package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"entgo.io/ent/dialect"
	"github.com/GavukaAlexandr/ent-sandbox/ent"
	"github.com/GavukaAlexandr/ent-sandbox/ent/migrate"
	"go.uber.org/zap"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	// CONFIG     Config
	dbHost       string
	dbPort       string
	dbName       string
	dbUser       string
	dbPassword   string
	Client       *ent.Client
	ClientDriver *entsql.Driver
)

func init() {
	dbHost = url.QueryEscape(os.Getenv("DB_HOST"))
	dbPort = url.QueryEscape(os.Getenv("DB_PORT"))
	dbName = url.QueryEscape(os.Getenv("DB_NAME"))
	dbUser = url.QueryEscape(os.Getenv("DB_USER"))
	dbPassword = url.QueryEscape(os.Getenv("DB_PASSWORD"))

	if dbHost == "" ||
		dbPort == "" ||
		dbName == "" ||
		dbPassword == "" {
		fmt.Println("DB access isn't configured")
		os.Exit(1)
	}
}

// Open new connection
// The returned DB Client is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a DB.
func OpenConnection() error {
	zap.L().Info("open database connection")
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s/%s", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return err
	}

	// Create an ent.Driver from `db`.
	ClientDriver = entsql.OpenDB(dialect.Postgres, db)
	Client = ent.NewClient(ent.Driver(ClientDriver))

	return nil
}

// Run the auto migration tool.
func AutoMigration(ctx *context.Context) {
	zap.L().Info("auto run database migration")
	if err := Client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				// Run custom code here.
				// rows, err := ClientDriver.QueryContext(ctx, "SELECT * FROM snipe_it_cloud_permissions")
				// if err != nil {
				// 	fmt.Println(err)
				// }

				// type row struct {
				// 	age  int
				// 	name string
				// }

				// got := []row{}
				// for rows.Next() {
				// 	var r row
				// 	err = rows.Scan(&r.age, &r.name)
				// 	if err != nil {
				// 		t.Fatalf("Scan: %v", err)
				// 	}
				// 	got = append(got, r)
				// }
				// err = rows.Err()
				// if err != nil {
				// 	t.Fatalf("Err: %v", err)
				// }

				// fmt.Println(got...)
				return next.Create(ctx, tables...)
			})
		})); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
