package conf

import (
	"io/ioutil"
	"os"
	"strings"
)

/*
 * Dead-simple implementation of the "dotenv" configuration format,
 * which reads config vars from the environment, optionally by sourcing a '.env' file.
 * See https://12factor.net/config for the idea behind this.
 */

func init() {
	env, err := ioutil.ReadFile(".env")
	if err != nil || len(env) == 0 {
		return
	}

	for _, line := range strings.Split(string(env), "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if line[0] == '#' {
			continue
		}

		kv := strings.SplitN(line, "=", 2)

		if kv[0] == "" {
			continue
		}

		if err := os.Setenv(kv[0], kv[1]); err != nil {
			panic(err)
		}
	}
}

func Get(key string) string {
	return os.Getenv(key)
}