package snapshot_test

import (
	"testing"

	"github.com/aarol/snapshot"
)

func TestSnapshot(t *testing.T) {
	val := 2
	snapshot.MatchesInline(t, val, 2)

	val2 := "123"
	snapshot.MatchesInline(t, val2, "123")
}
