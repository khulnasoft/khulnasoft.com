package khulnasoft

// Meta returns metadata about the running application.
//
// Meta will never return nil.
func Meta() (_ *AppMetadata) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/pkgfn.go#L16-L18
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// CurrentRequest returns the Request that is currently being handled by the calling goroutine
//
// It is safe for concurrent use and will return a new Request on each evocation, so can be mutated by the
// calling code without impacting future calls.
//
// CurrentRequest never returns nil.
func CurrentRequest() (_ *Request) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/pkgfn.go#L26-L28
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}
