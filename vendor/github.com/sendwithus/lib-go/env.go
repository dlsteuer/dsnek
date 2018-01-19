package swu

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetEnvVariable(name string, required bool) string {
	variable, exists := os.LookupEnv(name)
	if required && !exists {
		panic(fmt.Sprintf("Unable to find required environment variable %v", name))
	} else if !required && !exists {
		log.Printf("Unable to find environment variable %v\n", name)
	}
	return variable
}

func GetEnvOneOf(args ...string) (envVar string) {
	for _, arg := range args {
		envVar = os.Getenv(arg)
		if envVar != "" {
			return
		}
	}
	return
}

func GetEnvDuration(envVar string, defaults time.Duration) time.Duration {
	value := os.Getenv(envVar)
	if value == "" {
		return defaults
	}
	dur, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}
	return dur
}

func GetEnvInt(envVar string, defaults int) int {
	value := os.Getenv(envVar)
	if value == "" {
		return defaults
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}

func GetEnvString(envVar string, defaults string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaults
	}
	return value
}

func GetEnvBool(envVar string) bool {
	return strings.ToLower(os.Getenv(envVar)) == "true"
}
