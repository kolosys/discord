package types

import (
	"github.com/kolosys/atomic/snowflake"
)

const (
	DiscordEpoch int64 = 1420070400000 // 2015-01-01T00:00:00.000Z
)

// Snowflake is a Discord snowflake ID.
type Snowflake = snowflake.Snowflake

// NullableSnowflake is a snowflake that can be null in JSON.
type NullableSnowflake = snowflake.Nullable

// SnowflakeSlice is a slice of snowflakes with JSON marshaling.
type SnowflakeSlice = snowflake.Slice

// NewSnowflakeFromString creates a new Discord Snowflake from a string.
func NewSnowflakeFromString(s string) Snowflake {
	return snowflake.NewFromString(s, DiscordEpoch)
}

// NewSnowflakeFromInt64 creates a new Discord Snowflake from an int64.
func NewSnowflakeFromInt64(id int64) Snowflake {
	return snowflake.NewFromInt64(id, DiscordEpoch)
}
