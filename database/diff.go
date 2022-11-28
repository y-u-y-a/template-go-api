package main

import (
	"context"
	"fmt"
	"log"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"

	"y-u-y-a/template-go/config"
	appMigrate "y-u-y-a/template-go/ent/migrate"
)

func main() {
	client, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.AppEnv.DbUser, config.AppEnv.DbPassword,
		config.AppEnv.DbHost, config.AppEnv.DbPort,
		config.AppEnv.DbName,
	))
	if err != nil {
		log.Fatalf("failed connecting to postgreSQL: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	graph, err := entc.LoadGraph("../ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}

	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating local dir: %v", err)
	}
	m, err := schema.NewMigrate(
		client, schema.WithDir(dir),
		appMigrate.WithDropColumn(true),
		appMigrate.WithForeignKeys(true),
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
	)

	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	err = m.Diff(ctx, tbls...)
	// You can use the following method to give the migration files a name.
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
