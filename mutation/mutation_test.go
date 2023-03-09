//go:build mutation

package tonality_test

import (
	"testing"

	"github.com/gtramontina/ooze"
)

func TestMutation(t *testing.T) {
	ooze.Release(
		t,
		ooze.WithRepositoryRoot("../"),
		ooze.WithTestCommand("make test"),
		ooze.IgnoreSourceFiles("^vendor/"),
	)
}
