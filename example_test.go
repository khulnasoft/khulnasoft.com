package khulnasoft_test

import (
	"fmt"
	"time"

	"khulnasoft.com"
)

type Client interface{}

var client Client

func NewRedshiftClient() Client         { return nil }
func NewBigQueryClient() Client         { return nil }
func LocalFileWriter(dir string) Client { return nil }

// Change the implementation of some code based on which cloud provider is being used.
func ExampleMeta() {
	switch khulnasoft.Meta().Environment.Cloud {
	case khulnasoft.CloudAWS:
		client = NewRedshiftClient()
	case khulnasoft.CloudGCP:
		client = NewBigQueryClient()
	case khulnasoft.CloudLocal:
		client = LocalFileWriter("/tmp/app-writes.txt")
	default:
		panic("unsupported cloud provider")
	}
}

// Check if the application is running in production
func ExampleMeta_2() {
	if khulnasoft.Meta().Environment.Type != khulnasoft.EnvProduction {
		fmt.Println("running in development")
	} else {
		fmt.Println("running in production")
	}
	// Output: running in development
}

func ExampleCurrentRequest() {
	req := khulnasoft.CurrentRequest()
	elapsed := time.Since(req.Started)

	if req.Type == khulnasoft.APICall {
		fmt.Printf("%s.%s has been running for %.3f seconds", req.Service, req.Endpoint, elapsed.Seconds())
	}
	// Output: myservice.api has been running for 0.543 seconds
}
