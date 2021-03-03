package main

import (
	"flag"

	"github.com/redpatronus/httpret/run"
	"github.com/redpatronus/httpret/utils"
)

func main() {
	var config string
	flag.StringVar(&config, "config", "", "path to config file")
	flag.Parse()

	s := utils.ParseConfig(config)
	run.HttpServiceStart(s)
}
