package errs

import "os"

// A Builder allows for gradual construction of an error.
// The zero value is ready for use.
//
// Use Err() to construct the error.
type Builder struct {
	_ int // for godoc to show unexported fields
}

// B is a shorthand for creating a new Builder.
func B() (_ *Builder) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//
	//	https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L27-L27
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Code sets the error code.
func (*Builder) Code(c ErrCode) (_ *Builder) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L30-L34
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Msg sets the error message.
func (*Builder) Msg(msg string) (_ *Builder) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L37-L40
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Msgf is like Msg but uses fmt.Sprintf to construct the message.
func (*Builder) Msgf(format string, args ...interface{}) (_ *Builder) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L43-L46
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Meta appends metadata key-value pairs.
func (*Builder) Meta(metaPairs ...interface{}) (_ *Builder) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L49-L52
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Details sets the details.
func (*Builder) Details(det ErrDetails) (_ *Builder) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L55-L59
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Cause sets the underlying error cause.
func (*Builder) Cause(err error) (_ *Builder) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L62-L77
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Err returns the constructed error.
// It never returns nil.
//
// If Code has not been set or has been set to OK,
// the Code is set to Unknown.
//
// If Msg has not been set and Cause is nil,
// the Msg is set to "unknown error".
func (*Builder) Err() (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/beta/errs/builder.go#L87-L117
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking KhulnaSoft APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("KHULNASOFTRUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
