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

//MapValues returns array of values of the map
func MapValues(m map[string]string) []string {
	vals := make([]string, len(m))

	ind := 0
	for _, v := range m {
		vals[ind] = v
		ind++
	}
	return vals
}
