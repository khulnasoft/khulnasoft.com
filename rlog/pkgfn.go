package rlog

import "os"

// Debug logs a debug-level message.
// The variadic key-value pairs are treated as they are in With.
func Debug(msg string, keysAndValues ...any) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/rlog/pkgfn.go#L12-L14
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Info logs an info-level message.
// The variadic key-value pairs are treated as they are in With.
func Info(msg string, keysAndValues ...any) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/rlog/pkgfn.go#L18-L20
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Warn logs a warn-level message.
// The variadic key-value pairs are treated as they are in With.
func Warn(msg string, keysAndValues ...any) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/rlog/pkgfn.go#L24-L26
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Error logs an error-level message.
// The variadic key-value pairs are treated as they are in With.
func Error(msg string, keysAndValues ...any) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/rlog/pkgfn.go#L30-L32
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// With adds a variadic number of fields to the logging context.
// The keysAndValues must be pairs of string keys and arbitrary data.
func With(keysAndValues ...any) (_ Ctx) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/rlog/pkgfn.go#L36-L38
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
