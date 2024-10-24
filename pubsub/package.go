// Package pubsub provides KhulnaSoft applications with the ability
// to create Pub/Sub Topics and multiple Subscriptions on those
// topics in a cloud-agnostic manner.
//
// For more information see https://khulnasoft.com/docs/develop/pubsub
package pubsub

import "os"

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking KhulnaSoft APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("KHULNASOFTRUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
