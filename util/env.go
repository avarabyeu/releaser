package util

import (
	"os"
	"strings"
)

//GetEnvVars returns all env vars as a map
func GetEnvVars() map[string]string {
	envs := map[string]string{}
	for _, envVar := range os.Environ() {
		parts := strings.Split(envVar, "=")
		envs[parts[0]] = parts[1]
	}
	return envs
}
