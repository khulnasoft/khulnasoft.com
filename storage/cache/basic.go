package cache

import (
	"context"
	"os"
)

// NewStringKeyspace creates a keyspace that stores string values in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
func NewStringKeyspace[K any](cluster *Cluster, cfg KeyspaceConfig) (_ *StringKeyspace[K]) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L16-L25
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// StringKeyspace represents a set of cache keys that hold string values.
type StringKeyspace[K any] struct {
	_ int // for godoc to show unexported fields
}

// Get gets the value stored at key.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/get/ for more information.
func (*StringKeyspace[K]) Get(ctx context.Context, key K) (_ string, _ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L36-L38
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Set updates the value stored at key to val.
//
// See https://redis.io/commands/set/ for more information.
func (*StringKeyspace[K]) Set(ctx context.Context, key K, val string) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L43-L45
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// SetIfNotExists sets the value stored at key to val, but only if the key does not exist beforehand.
// If the key already exists, it reports an error matching KeyExists.
//
// See https://redis.io/commands/setnx/ for more information.
func (*StringKeyspace[K]) SetIfNotExists(ctx context.Context, key K, val string) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L51-L53
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Replace replaces the existing value stored at key to val.
// If the key does not already exist, it reports an error matching Miss.
//
// See https://redis.io/commands/set/ for more information.
func (*StringKeyspace[K]) Replace(ctx context.Context, key K, val string) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L59-L61
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// GetAndSet updates the value of key to val and returns the previously stored value.
// If the key does not already exist, it sets it and returns "", nil.
//
// See https://redis.io/commands/getset/ for more information.
func (*StringKeyspace[K]) GetAndSet(ctx context.Context, key K, val string) (oldVal string, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L67-L69
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// GetAndDelete deletes the key and returns the previously stored value.
// If the key does not already exist, it does nothing and returns "", nil.
//
// See https://redis.io/commands/getdel/ for more information.
func (*StringKeyspace[K]) GetAndDelete(ctx context.Context, key K) (oldVal string, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L75-L77
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Delete deletes the specified keys.
//
// If a key does not exist it is ignored.
//
// It reports the number of keys that were deleted.
//
// See https://redis.io/commands/del/ for more information.
func (*StringKeyspace[K]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L86-L88
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*StringKeyspace[K]) With(opts ...WriteOption) (_ *StringKeyspace[K]) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L96-L98
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Append appends to the string with the given key.
//
// If the key does not exist it is first created and set as the empty string,
// causing Append to behave like Set.
//
// It returns the new string length.
//
// See https://redis.io/commands/append/ for more information.
func (*StringKeyspace[K]) Append(ctx context.Context, key K, val string) (newLen int64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L108-L121
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// GetRange returns a substring of the string value stored in key.
//
// The from and to values are zero-based indices, but unlike Go slicing
// the 'to' value is inclusive.
//
// Negative values can be used in order to provide an offset starting
// from the end of the string, so -1 means the last character
// and -len(str) the first character, and so forth.
//
// If the string does not exist it returns the empty string.
//
// See https://redis.io/commands/setrange/ for more information.
func (*StringKeyspace[K]) GetRange(ctx context.Context, key K, from, to int64) (val string, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L135-L146
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// SetRange overwrites part of the string stored at key, starting at
// the zero-based offset and for the entire length of val, extending
// the string if necessary to make room for val.
//
// If the offset is larger than the current string length stored at key,
// the string is first padded with zero-bytes to make offset fit.
//
// Non-existing keys are considered as empty strings.
//
// See https://redis.io/commands/setrange/ for more information.
func (*StringKeyspace[K]) SetRange(ctx context.Context, key K, offset int64, val string) (newLen int64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L158-L171
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Len reports the length of the string value stored at key.
//
// Non-existing keys are considered as empty strings.
//
// See https://redis.io/commands/strlen/ for more information.
func (*StringKeyspace[K]) Len(ctx context.Context, key K) (length int64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L178-L189
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// NewIntKeyspace creates a keyspace that stores int64 values in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
func NewIntKeyspace[K any](cluster *Cluster, cfg KeyspaceConfig) (_ *IntKeyspace[K]) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L195-L204
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// IntKeyspace is a cache keyspace that stores int64 values.
type IntKeyspace[K any] struct {
	_ int // for godoc to show unexported fields
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*IntKeyspace[K]) With(opts ...WriteOption) (_ *IntKeyspace[K]) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L217-L219
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Get gets the value stored at key.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/get/ for more information.
func (*IntKeyspace[K]) Get(ctx context.Context, key K) (_ int64, _ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L225-L227
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Set updates the value stored at key to val.
//
// See https://redis.io/commands/set/ for more information.
func (*IntKeyspace[K]) Set(ctx context.Context, key K, val int64) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L232-L234
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// SetIfNotExists sets the value stored at key to val, but only if the key does not exist beforehand.
// If the key already exists, it reports an error matching KeyExists.
//
// See https://redis.io/commands/setnx/ for more information.
func (*IntKeyspace[K]) SetIfNotExists(ctx context.Context, key K, val int64) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L240-L242
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Replace replaces the existing value stored at key to val.
// If the key does not already exist, it reports an error matching Miss.
//
// See https://redis.io/commands/set/ for more information.
func (*IntKeyspace[K]) Replace(ctx context.Context, key K, val int64) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L248-L250
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// GetAndSet updates the value of key to val and returns the previously stored value.
// If the key does not already exist, it sets it and returns 0, nil.
//
// See https://redis.io/commands/getset/ for more information.
func (*IntKeyspace[K]) GetAndSet(ctx context.Context, key K, val int64) (oldVal int64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L256-L258
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// GetAndDelete deletes the key and returns the previously stored value.
// If the key does not already exist, it does nothing and returns 0, nil.
//
// See https://redis.io/commands/getdel/ for more information.
func (*IntKeyspace[K]) GetAndDelete(ctx context.Context, key K) (oldVal int64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L264-L266
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Delete deletes the specified keys.
//
// If a key does not exist it is ignored.
//
// It reports the number of keys that were deleted.
//
// See https://redis.io/commands/del/ for more information.
func (*IntKeyspace[K]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L275-L277
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Increment increments the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before incrementing.
//
// Negative values can be used to decrease the value,
// but typically you want to use the Decrement method for that.
//
// See https://redis.io/commands/incrby/ for more information.
func (*IntKeyspace[K]) Increment(ctx context.Context, key K, delta int64) (newVal int64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L289-L302
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Decrement decrements the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before decrementing.
//
// Negative values can be used to increase the value,
// but typically you want to use the Increment method for that.
//
// See https://redis.io/commands/decrby/ for more information.
func (*IntKeyspace[K]) Decrement(ctx context.Context, key K, delta int64) (newVal int64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L314-L328
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// NewFloatKeyspace creates a keyspace that stores float64 values in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
func NewFloatKeyspace[K any](cluster *Cluster, cfg KeyspaceConfig) (_ *FloatKeyspace[K]) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L334-L343
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// FloatKeyspace is a cache keyspace that stores float64 values.
type FloatKeyspace[K any] struct {
	_ int // for godoc to show unexported fields
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*FloatKeyspace[K]) With(opts ...WriteOption) (_ *FloatKeyspace[K]) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L356-L358
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Get gets the value stored at key.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/get/ for more information.
func (*FloatKeyspace[K]) Get(ctx context.Context, key K) (_ float64, _ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L364-L366
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Set updates the value stored at key to val.
//
// See https://redis.io/commands/set/ for more information.
func (*FloatKeyspace[K]) Set(ctx context.Context, key K, val float64) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L371-L373
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// SetIfNotExists sets the value stored at key to val, but only if the key does not exist beforehand.
// If the key already exists, it reports an error matching KeyExists.
//
// See https://redis.io/commands/setnx/ for more information.
func (*FloatKeyspace[K]) SetIfNotExists(ctx context.Context, key K, val float64) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L379-L381
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Replace replaces the existing value stored at key to val.
// If the key does not already exist, it reports an error matching Miss.
//
// See https://redis.io/commands/set/ for more information.
func (*FloatKeyspace[K]) Replace(ctx context.Context, key K, val float64) (_ error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L387-L389
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// GetAndSet updates the value of key to val and returns the previously stored value.
// If the key does not already exist, it sets it and returns 0, nil.
//
// See https://redis.io/commands/getset/ for more information.
func (*FloatKeyspace[K]) GetAndSet(ctx context.Context, key K, val float64) (oldVal float64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L395-L397
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// GetAndDelete deletes the key and returns the previously stored value.
// If the key does not already exist, it does nothing and returns 0, nil.
//
// See https://redis.io/commands/getdel/ for more information.
func (*FloatKeyspace[K]) GetAndDelete(ctx context.Context, key K) (oldVal float64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L403-L405
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Delete deletes the specified keys.
//
// If a key does not exist it is ignored.
//
// It reports the number of keys that were deleted.
//
// See https://redis.io/commands/del/ for more information.
func (*FloatKeyspace[K]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L414-L416
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Increment increments the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before incrementing.
//
// Negative values can be used to decrease the value,
// but typically you want to use the Decrement method for that.
//
// See https://redis.io/commands/incrbyfloat/ for more information.
func (*FloatKeyspace[K]) Increment(ctx context.Context, key K, delta float64) (newVal float64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L428-L441
	doPanic("khulnasoft apps must be run using the khulnasoft command")
	return
}

// Decrement decrements the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before decrementing.
//
// Negative values can be used to increase the value,
// but typically you want to use the Increment method for that.
//
// See https://redis.io/commands/incrbyfloat/ for more information.
func (*FloatKeyspace[K]) Decrement(ctx context.Context, key K, delta float64) (newVal float64, err error) {
	// KhulnaSoft will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/khulnasoft/khulnasoft/blob/v1.41.9/runtimes/go/storage/cache/basic.go#L453-L466
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
