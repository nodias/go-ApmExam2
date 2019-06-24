package main

import (
	"github.com/nodias/go-ApmCommon/middleware"
	"github.com/nodias/go-ApmCommon/model"
	"github.com/nodias/go-ApmExam2/router"
	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	config.New("config.toml")
}
func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logpaths["local"].Path)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["local2"].PORT)
}
