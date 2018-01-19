package swu

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	_routeLogging bool
	awsSession    *session.Session
)

func getEnv(name string, required bool) string {
	v := os.Getenv(name)
	if required && v == "" {
		panic(fmt.Sprintf("$%s not set", name))
	}

	return v
}

func initConfig() {
	_routeLogging = getEnv("ROUTE_LOGGING", false) != ""
}

func GetAWSSession() *session.Session {
	if awsSession == nil {
		region := os.Getenv("SWU_AWS_REGION")
		if region == "" {
			region = "us-east-1"
		}
		var err error
		awsSession, err = session.NewSession(&aws.Config{
			Region: &region,
		})
		if err != nil {
			panic(err)
		}
	}
	return awsSession
}
