package cache

// NewCluster declares a new cache cluster.
//
// See https://khulnasoft.com/docs/develop/caching for more information.
func NewCluster(name string, cfg ClusterConfig) (_ *Cluster) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/pkgfn.go#L8-L14
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}
