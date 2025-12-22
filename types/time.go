package types

import (
	"github.com/kolosys/atomic/timestamp"
)

// Timestamp is an ISO8601/RFC3339 timestamp.
type Timestamp = timestamp.Timestamp

// NullableTimestamp is a timestamp that can be null in JSON.
type NullableTimestamp = timestamp.Nullable
