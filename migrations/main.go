package migrations

import (
	"github.com/nanpipat/go-api-template/consts"
	"github.com/nanpipat/go-api-template/core"
)

func RunMigration(ctx core.IContext) error {
	tables := []interface{}{
		// &User{},
	}

	if ctx.ENV().String(consts.EnvIsThanos) == "true" || ctx.ENV().String(consts.EnvIsThanos) == "TRUE" {
		err := ctx.DB().Migrator().DropTable(tables...)
		if err != nil {
			return err
		}
	}

	err := ctx.DB().AutoMigrate(tables...)
	if err != nil {
		return err
	}

	return nil
}
