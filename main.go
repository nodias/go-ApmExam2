package main

import (
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam2/router"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	config.Load()
}
func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logpaths.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam2"].PORT)
}
