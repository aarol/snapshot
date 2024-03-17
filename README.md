# Snapshot

Package snapshot implements [snapshot testing](https://vitest.dev/guide/snapshot.html#inline-snapshots) for Go.

It works by modifying the *test file itself* during execution.

## Installation

```sh
go get github.com/aarol/snapshot
```

## Example

* Step 1: Call the `MatchesInline` function inside your tests:

```go
package snapshot_test

import (
	"testing"

	"github.com/aarol/snapshot"
)

func TestSnapshot(t *testing.T) {
	val := 1
	snapshot.MatchesInline(t, val)

	val2 := "123"
	snapshot.MatchesInline(t, val2)
}
```

* Step 2: Run the tests with `go test` and observe the added parameters:

```go

func TestSnapshot(t *testing.T) {
	val := 1
	snapshot.MatchesInline(t, val, 1)

	val2 := "123"
	snapshot.MatchesInline(t, val2, "123")
}
```

Now, any future values of `val` or `val2` will be compared against the hard-coded value. If it were to change, the test would error:

```sh
> go test .

--- FAIL: TestSnapshot (0.00s)
    snapshot_test.go:11: snapshot doesn't match: want 1, got 2
```
