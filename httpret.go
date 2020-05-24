package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/redpatronus/httpret/utils"
)

func main() {
	var config string
	flag.StringVar(&config, "config", "", "path to config file")
	flag.Parse()

	s := utils.ParseConfig(config)

	log.Println("Starting HTTP listener: " + s.Listen)
	http.HandleFunc("/", s.HttpRet)
	http.ListenAndServe(s.Listen, nil)
}
