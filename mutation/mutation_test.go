//go:build mutation

package mutation

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
		ooze.ForceColors(),
	)
}
