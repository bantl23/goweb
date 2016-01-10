package main

import (
	"github.com/codegangsta/cli"
	"log"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "goweb"
	app.Usage = "Simple directory web server"
	app.Version = "1.0.0"
	var addr string = ":8080"
	var dir string = "."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "a, addr",
			Value:       addr,
			Usage:       "server address",
			Destination: &addr,
		},
		cli.StringFlag{
			Name:        "d, dir",
			Value:       dir,
			Usage:       "directory to serve",
			Destination: &dir,
		},
	}
	app.Action = func(cmd *cli.Context) {
		log.Println("Serving", dir, "on", addr)
		http.ListenAndServe(addr, http.FileServer(http.Dir(dir)))
	}
	app.Run(os.Args)
}
