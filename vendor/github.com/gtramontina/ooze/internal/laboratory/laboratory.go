package laboratory

import (
	"github.com/gtramontina/ooze/internal/future"
	"github.com/gtramontina/ooze/internal/gomutatedfile"
	"github.com/gtramontina/ooze/internal/ooze"
	"github.com/gtramontina/ooze/internal/result"
)

type TestRunner interface {
	Test(repository ooze.TemporaryRepository) result.Result[string]
}

type TemporaryDirectory interface {
	New() string
}

type Laboratory struct {
	testRunner         TestRunner
	temporaryDirectory TemporaryDirectory
}

func New(testRunner TestRunner, temporaryDirectory TemporaryDirectory) *Laboratory {
	return &Laboratory{
		testRunner:         testRunner,
		temporaryDirectory: temporaryDirectory,
	}
}

func (l *Laboratory) Test(
	repository ooze.Repository,
	file *gomutatedfile.GoMutatedFile,
) future.Future[result.Result[string]] {
	tempRepository := repository.LinkAllToTemporaryRepository(l.temporaryDirectory.New())
	defer tempRepository.Remove()

	file.WriteTo(tempRepository)

	return future.Resolved(l.testRunner.Test(tempRepository))
}
