package khulnasoft

import (
	"net/url"
	"os"
	"time"
)

// AppMetadata contains metadata about the running KhulnaSoft application.
type AppMetadata struct {
	// The application ID, if the application is not linked to the KhulnaSoft platform this will be an empty string.
	//
	// To link to the KhulnaSoft platform run `khulnasoft app link` from your terminal in the root directory of the KhulnaSoft app.
	AppID string

	// The base URL which can be used to call the API of this running application.
	//
	// For local development it is "http://localhost:<port>", typically "http://localhost:4000".
	//
	// If a custom domain is used for this environment it is returned here, but note that
	// changes only take effect at the time of deployment while custom domains can be updated at any time.
	APIBaseURL url.URL

	// Information about the environment the app is running in.
	Environment EnvironmentMeta

	// Information about the running binary itself.
	Build BuildMeta

	// Information about this deployment of the binary
	Deploy DeployMeta
}

type EnvironmentMeta struct {
	// The name of environment that this application.
	// For local development it is "local".
	Name string

	// The type of environment is this application running in
	// For local development this will be EnvLocal
	Type EnvironmentType

	// The cloud that this environment is running on
	// For local development this is CloudLocal
	Cloud CloudProvider
}

type BuildMeta struct {
	// The git commit that formed the base of this build.
	Revision string

	// true if there were uncommitted changes on top of the Commit.
	UncommittedChanges bool
}

type DeployMeta struct {
	// The deployment ID created by the KhulnaSoft Platform.
	ID string

	// The time the KhulnaSoft Platform deployed this build to the environment.
	Time time.Time
}

// EnvironmentType represents the type of environment.
//
// For more information on environment types see https://khulnasoft.com/docs/deploy/environments#environment-types
//
// Additional environment types may be added in the future.
type EnvironmentType string

const (
	// EnvProduction represents a production environment.
	EnvProduction EnvironmentType = "production"

	// EnvDevelopment represents a long-lived cloud-hosted, non-production environment, such as test environments.
	EnvDevelopment EnvironmentType = "development"

	// EnvEphemeral represents short-lived cloud-hosted, non-production environments, such as preview environments
	// that only exist while a particular pull request is open.
	EnvEphemeral EnvironmentType = "ephemeral"

	// EnvLocal represents the local development environment when using 'khulnasoft run' or `khulnasoft test`.
	//
	// Deprecated: EnvLocal is deprecated and KhulnaSoft will no longer return this value. A locally running environment
	// can be identified by the combination of EnvDevelopment && CloudLocal. This constant will be removed in a future
	// version of KhulnaSoft.
	EnvLocal EnvironmentType = "local"

	// EnvTest represents a running unit test
	EnvTest EnvironmentType = "test"
)

// CloudProvider represents the cloud provider this application is running in.
//
// For more information about how Cloud Providers work with KhulnaSoft, see https://khulnasoft.com/docs/deploy/own-cloud
//
// Additional cloud providers may be added in the future.
type CloudProvider string

const (
	CloudAWS   CloudProvider = "aws"
	CloudGCP   CloudProvider = "gcp"
	CloudAzure CloudProvider = "azure"

	// KhulnaSoftCloud is KhulnaSoft's own cloud offering, and the default provider for new Environments.
	KhulnaSoftCloud CloudProvider = "khulnasoft"

	// CloudLocal is used when an application is running from the KhulnaSoft CLI by using either
	// 'khulnasoft run' or 'khulnasoft test'
	CloudLocal CloudProvider = "local"
)

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking KhulnaSoft APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("KHULNASOFTRUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
