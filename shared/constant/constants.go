// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/ramensoft/handinger-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Assistant string  // Always "assistant"
type Completed string  // Always "completed"
type Cron string       // Always "cron"
type Delayed string    // Always "delayed"
type Interval string   // Always "interval"
type Message string    // Always "message"
type OutputText string // Always "output_text"
type Scheduled string  // Always "scheduled"
type URL string        // Always "url"
type Worker string     // Always "worker"

func (c Assistant) Default() Assistant   { return "assistant" }
func (c Completed) Default() Completed   { return "completed" }
func (c Cron) Default() Cron             { return "cron" }
func (c Delayed) Default() Delayed       { return "delayed" }
func (c Interval) Default() Interval     { return "interval" }
func (c Message) Default() Message       { return "message" }
func (c OutputText) Default() OutputText { return "output_text" }
func (c Scheduled) Default() Scheduled   { return "scheduled" }
func (c URL) Default() URL               { return "url" }
func (c Worker) Default() Worker         { return "worker" }

func (c Assistant) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Completed) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Cron) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Delayed) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Interval) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Message) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c OutputText) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Scheduled) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c URL) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Worker) MarshalJSON() ([]byte, error)     { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
