package job

import (
	"fmt"
	"github.com/oklog/run"
	"github.com/shoppigram-com/marketplace-api/packages/logger"
)

type Runner interface {
	Run() error
	Shutdown(error)
}

var (
	l = logger.New("INFO")
)

// Run runs the given job in a goroutine and adds it to the run.Group
func Run(r Runner, g *run.Group, enabled bool, name string) {
	if enabled {
		// Wrap the Run method with panic recovery
		runWithRecovery := func() (err error) {
			defer func() {
				if rec := recover(); rec != nil {
					l.Error(fmt.Sprintf("panic in job %s: %v", name, rec))
					err = fmt.Errorf("panic occurred: %v", rec)
				}
			}()
			err = r.Run()
			return
		}

		// Wrap the Shutdown method with panic recovery
		shutdownWithRecovery := func(err error) {
			defer func() {
				if rec := recover(); rec != nil {
					l.Error(fmt.Sprintf("panic during shutdown of job %s: %v", name, rec))
				}
			}()
			r.Shutdown(err)
		}

		g.Add(runWithRecovery, shutdownWithRecovery)
	} else {
		l.Info(fmt.Sprintf("job %s is disabled", name))
	}
}
