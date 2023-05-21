package main

import (
	"fmt"
	"os"
	"simple_conf_manager/internal/args"
	"simple_conf_manager/internal/env"
	"simple_conf_manager/internal/templating"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	arg, err := args.New()
	handleError(err)
	envs, err := env.Get(*arg.EnvPrefix, *arg.ErrorOnEmptyEnv)
	handleError(err)
	handleError(templating.Parse(*arg.Filepath, *arg.OutPutFilepath, envs))
}
