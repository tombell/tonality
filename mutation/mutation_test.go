//go:build mutation

package mutation

import (
	"testing"

	"github.com/gtramontina/ooze"
)

func TestMutation(t *testing.T) {
	ooze.Release(
		t,
		ooze.ForceColors(),
		ooze.WithRepositoryRoot("../"),
		ooze.WithTestCommand("make test"),
		ooze.WithMinimumThreshold(0.5),
		ooze.Parallel(),
		ooze.IgnoreSourceFiles("^vendor/"),
	)
}
