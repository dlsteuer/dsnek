package swu

import (
	"strings"
	"strconv"
	"github.com/bgentry/heroku-go"
	"errors"
	"os"
)

var envVarNotExist = errors.New("environment variable does not exist")

// Get the dyno name
func GetDynoName() string {
	return GetEnvOneOf("DYNO", "PS")
}

// Get the dyno's index
func GetDynoIndex() (int, error) {
	name := GetDynoName()
	if name == "" {
		return 0, envVarNotExist
	}

	spl := strings.Split(GetDynoName(), ".")
	idx, err := strconv.Atoi(spl[len(spl)-1])
	if err != nil {
		return 0, err
	}
	return idx, nil
}

// Get the dyno app name, this is from the DYNO environment variable
func GetDynoAppName() (string, error) {
	name := GetDynoName()
	if name == "" {
		return "", envVarNotExist
	}
	return strings.Split(GetDynoName(), ".")[0], nil
}

// Fetch the heroku dyno count for this application
func GetDynoCount() (int, error) {
	apiKey := os.Getenv("HEROKU_API_KEY")
	name, err := GetDynoAppName()
	if err != nil {
		return 0, err
	}
	client := heroku.Client{
		Password: apiKey,
	}
	dynos, err := client.DynoList(name, nil)
	if err != nil {
		return 0, err
	}
	return len(dynos), nil
}
