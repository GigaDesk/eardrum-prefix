package parse

import (
	"fmt"

	"github.com/google/uuid"
)

// ParseUUIDSlice converts a slice of string UUIDs to a slice of uuid.UUID structs.
func ParseUUIDSlice(strs []string) ([]uuid.UUID, error) {
	uuids := make([]uuid.UUID, len(strs))
	for i, str := range strs {
		parsed, err := uuid.Parse(str)
		if err != nil {
			return nil, fmt.Errorf("failed to parse string at index %d (%q): %w", i, str, err)
		}
		uuids[i] = parsed
	}
	return uuids, nil
}