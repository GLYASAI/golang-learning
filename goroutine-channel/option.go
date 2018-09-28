package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

type apiServer struct {
	LogLevel string
}

func main() {
	a := new(apiServer)
	pflag.CommandLine.StringVar(&a.LogLevel, "log-level", "info", "the entrance log level")
	fmt.Printf(a.LogLevel)
}