package chain

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/executor"
)

type (
	PublishEventStep struct {
		Executor executor.CommandExecutor
		executor.CommandExecutor
	}
	BackupCompensatingStep struct {
		Executor executor.CommandExecutor
		executor.CommandExecutor
	}
	PrepareEventIdStep struct {
		Executor executor.CommandExecutor
		executor.CommandExecutor
	}
)
