package args

import (
	"errors"
	"flag"
	"os"
	"path"
)

type Args struct {
	Filepath        *string
	OutPutFilepath  *string
	EnvPrefix       *string
	ErrorOnEmptyEnv *bool

	err error
}

func New() (*Args, error) {
	a := &Args{
		Filepath:        flag.String("filepath", "", "Path to the configuration template file"),
		OutPutFilepath:  flag.String("output_filepath", "", "Path to the output configuration file"),
		EnvPrefix:       flag.String("env_prefix", "", "Prefix for the environment variable"),
		ErrorOnEmptyEnv: flag.Bool("error_on_empty", false, "Returns an error if the environment variable value is empty"),
	}
	a.parse()
	return a, a.err
}

func (a *Args) filepathValidate() {
	if a.err != nil {
		return
	}
	if a.Filepath == nil || *a.Filepath == "" {
		a.err = errors.New("filepath is required")
		return
	}

	f, err := os.Stat(*a.Filepath)
	if err != nil {
		a.err = errors.New("wrong filepath: " + err.Error())
		return
	}
	if f.IsDir() {
		a.err = errors.New("filepath is a directory")
		return
	}
}

func (a *Args) outputFilepathValidate() {
	if a.err != nil {
		return
	}
	if a.Filepath == nil || *a.Filepath == "" {
		a.err = errors.New("filepath is required")
		return
	}
	dir := path.Dir(*a.Filepath)
	f, err := os.Stat(dir)
	if err != nil {
		a.err = errors.New("wrong filepath: " + err.Error())
		return
	}
	if !f.IsDir() {
		a.err = errors.New("parent filepath is not directory")
		return
	}
}

func (a *Args) envPrefixValidate() {
	if a.err != nil {
		return
	}
	if a.EnvPrefix == nil || *a.EnvPrefix == "" {
		a.err = errors.New("env prefix is required")
		return
	}
}

func (a *Args) parse() {
	flag.Parse()
	a.filepathValidate()
	a.outputFilepathValidate()
	a.envPrefixValidate()
}
