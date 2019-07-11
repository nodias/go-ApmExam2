package main

import (
	"context"
	"go-ApmCommon/models"
	"go-ApmCommon/shared/logger"
	"go-ApmExam2/api"
	"go-ApmExam2/router"

	"github.com/urfave/negroni"
)

var config models.TomlConfig

func init() {
	models.Load("config/%s/config.toml")
	config = *models.GetConfig()
	logger.Init()
	api.Init()
}
func main() {
	log := logger.New(context.Background())
	n := negroni.New()
	n.UseHandler(router.NewRouter())
	log.Info("go-ApmExam2 - Server On!")
	n.Run(config.Servers["ApmExam2"].PORT)
}
