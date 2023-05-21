package env

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Get(envPrefix string, errorOnEmpty bool) (map[string]string, error) {
	vars := make(map[string]string)

	for _, env := range os.Environ() {
		s := strings.SplitN(env, "=", 2)
		name, value := s[0], s[1]

		if strings.HasPrefix(name, envPrefix) {
			name = strings.ToUpper(strings.TrimPrefix(name, envPrefix))
			if value == "" && errorOnEmpty {
				return nil, errors.New(fmt.Sprintf("env [%s] value is empty", name))
			}
			vars[name] = value
		}
	}

	return vars, nil
}
