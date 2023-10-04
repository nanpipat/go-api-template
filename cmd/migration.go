package cmd

import (
	"fmt"
	"os"

	"github.com/nanpipat/go-api-template/core"
	"github.com/nanpipat/go-api-template/migrations"
	"gorm.io/gorm"
)

func MigrationRun() {
	env := core.NewEnv()

	fmt.Println(env.Config(), "env.Config()")

	mysql, err := core.NewDatabaseWithConfig(env.Config(), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "MSSQL: %v", err)
		os.Exit(1)
	}

	ctx := core.NewContext(&core.ContextOptions{
		DB:  mysql,
		ENV: env,
	})

	err = migrations.RunMigration(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Migrate: %v", err)
		os.Exit(1)
	}
}
