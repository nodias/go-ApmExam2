package main

import (
	"github.com/nodias/go-ApmCommon/middleware"
	"github.com/nodias/go-ApmExam2/router"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware("C:/workspace/logs/go-ApmExam2.log")))
	n.UseHandler(router.NewRouter())
	n.Run(":7002")
}
