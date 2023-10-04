package main

//generate main.go

import (
	"log"

	"github.com/nanpipat/go-api-template/cmd"
	"github.com/nanpipat/go-api-template/consts"
	"github.com/nanpipat/go-api-template/core"
)

func main() {
	switch core.NewEnv().Config().Service {
	case consts.ServiceAPI:
		cmd.APIRun()
	// case consts.ServiceConsumer:
	// 	cmd.MQRun()
	case consts.ServiceMigration:
		cmd.MigrationRun()
	// case consts.ServiceSeed:
	// 	cmd.SeedRun()
	// case consts.ServiceCronjob:
	// 	cmd.CronjobRun()
	default:
		log.Fatal("Service not support")
	}
}
