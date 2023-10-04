package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nanpipat/go-api-template/core"
)

func APIRun() {
	env := core.NewEnv()

	db, err := core.NewDatabase(env.Config()).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "MySQL: %v", err)
		os.Exit(1)
	}

	e := core.NewHTTPServer(&core.HTTPContextOptions{
		ContextOptions: &core.ContextOptions{
			DB:  db,
			ENV: env,
		},
	})

	e.GET("/healthz", core.WithHTTPContext(func(c core.IHTTPContext) error {
		return c.String(http.StatusOK, "OK!")
	}))

	core.StartHTTPServer(e, env)

}
