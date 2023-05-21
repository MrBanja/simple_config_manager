package templating

import (
	"errors"
	"os"
	"path"
	"text/template"
)

func Parse(filepath, outFilepath string, envs map[string]string) error {
	base := path.Base(filepath)
	t, err := template.New(base).ParseFiles(filepath)
	if err != nil {
		return errors.New("filepath: " + err.Error())
	}

	f, err := os.Create(outFilepath)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return errors.New("out_filepath: " + err.Error())
	}

	err = t.Execute(f, envs)
	if err != nil {
		return err
	}
	return nil
}
