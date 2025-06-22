package container

import (
	"rv/pkg/cron"
)

type workers struct {
	c  *Container
	cr *cron.Cron
}

func (c *Container) getWorkers() *workers {
	if c.workers == nil {
		c.workers = &workers{
			cr: cron.NewCron(),
			c:  c,
		}
	}
	return c.workers
}

/*
func (w *workers) getMatchJob() *match.Cron {
	if w.match == nil {
		w.match = match.NewCron(w.c.getLogger(), w.c.getApplication().getMatchApplicationService())
	}
	return w.match
}*/


func (w *workers) start() error {
	/*if err := w.cr.AddFunc(w.c.getConfig().Cron.DeclineStuckMatches, w.getMatchJob().DeclineStuckMatches); err != nil {
		return fmt.Errorf("DeclineStuckMatches: %v", err)
	}
*/
	w.cr.Start()
	return nil
}

func (w *workers) stop() {
	w.cr.Stop()
}

