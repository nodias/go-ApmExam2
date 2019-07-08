package main

import (
	"context"
	"go-ApmCommon/logger"
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam2/api"
	"go-ApmExam2/router"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	model.Load("config/%s/config.toml")
	config = *model.GetConfig()
	logger.Init()
	api.Init()
	log := logger.New(context.Background())
	log.Info(config)
}
func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.Logging(config.Logconfig.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam2"].PORT)
}
