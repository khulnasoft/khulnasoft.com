package et

import (
	"khulnasoft.com/pubsub"
)

// Topic returns a TopicHelper for the given topic.
func Topic[T any](topic *pubsub.Topic[T]) (_ TopicHelpers[T]) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/et/pubsub.go#L8-L10
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// TopicHelpers provides functions for interacting with the backing topic implementation
// during unit tests. It is designed to help test code that uses the pubsub.Topic
//
// Note all functions on this TopicHelpers are scoped to the current test
// and will only impact and observe state from the current test
type TopicHelpers[T any] interface {
	// PublishedMessages returns a slice of all messages published during this test on this topic.
	PublishedMessages() []T
}