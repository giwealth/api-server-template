package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-service-template/pkg/logger"

	"api-service-template/internal/option"
	"api-service-template/internal/presentation/httpapi"

	"github.com/urfave/cli/v2"
)

var (
	app         = newApp()
	opt         *option.Options
	commit      string
	branch      string
	buildTime   string
	buildAuthor string
	pipeline    string
)

func main() {
	app.Commands = []*cli.Command{
		{
			Name:   "api",
			Usage:  "start api service",
			Action: httpWeb,
		},
		{
			Name:  "version",
			Usage: "show version",
			Action: func(c *cli.Context) error {
				Content := fmt.Sprintf(
					"commit: %s\nbranch: %s\nbuild time: %s\npipeline: %s\nbuild author: %s\n",
					commit, branch, buildTime, pipeline, buildAuthor,
				)
				fmt.Fprintln(os.Stdout, Content)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.DefaultEntry().Fatal(err)
	}
}

func httpWeb(c *cli.Context) error {
	server := httpapi.NewServer()
	srv, err := server.Run(opt)
	if err != nil {
		logger.DefaultEntry().Fatal("start server", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.DefaultEntry().Info("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.DefaultEntry().Fatal("Server forced to shutdown:", err)
	}

	logger.DefaultEntry().Info("Server exiting")

	return nil
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "api-service-template"
	app.Usage = "api-service-template api server"
	app.Version = commit
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "configration file path",
		},
	}
	app.Before = func(c *cli.Context) error {
		cfg := c.String("config")
		opt = option.New(cfg)

		if err := opt.Parse(); err != nil {
			return err
		}

		return opt.Prepare()
	}

	return app
}
