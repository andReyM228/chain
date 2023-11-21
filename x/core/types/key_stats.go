package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StatsKeyPrefix is the prefix to retrieve all Stats
	StatsKeyPrefix = "Stats/value/"
)

// StatsKey returns the store key to retrieve a Stats from the index fields
func StatsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
