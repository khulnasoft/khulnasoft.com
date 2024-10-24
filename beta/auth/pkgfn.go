package auth

// UserID reports the uid of the user making the request.
// The second result is true if there is a user and false
// if the request was made without authentication details.
func UserID() (_ UID, _ bool) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/auth/pkgfn.go#L13-L15
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Data returns the structured auth data for the request.
// It returns nil if the request was made without authentication details,
// and the API endpoint does not require them.
//
// Expected usage is to immediately cast it to the registered auth data type:
//
//	usr, ok := auth.Data().(*user.Data)
//	if !ok { /* ... */ }
func Data() (_ any) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/auth/pkgfn.go#L25-L27
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}
