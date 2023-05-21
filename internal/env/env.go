package env

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Get(envPrefix string, errorOnEmpty bool) (map[string]string, error) {
	envPrefix = strings.ToUpper(envPrefix)
	vars := make(map[string]string)

	for _, env := range os.Environ() {
		s := strings.SplitN(env, "=", 2)
		name, value := s[0], s[1]
		name = strings.ToUpper(name)
		if strings.HasPrefix(name, envPrefix) {
			if value == "" && errorOnEmpty {
				return nil, errors.New(fmt.Sprintf("env [%s] value is empty", name))
			}
			name = strings.TrimPrefix(name, envPrefix)
			vars[name] = value
		}
	}

	return vars, nil
}
