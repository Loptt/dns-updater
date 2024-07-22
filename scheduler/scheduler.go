package scheduler

import (
	"log"
	"time"

	"github.com/Loptt/dns-updater/updater"
)

type Scheduler struct {
	interval time.Duration
	ui       updater.UpdaterInterface
}

func (s *Scheduler) Run() {
	runTicker := time.NewTicker(s.interval)
    defer runTicker.Stop()

	for range runTicker.C {
		if err := s.ui.Update(); err != nil {
			log.Printf("Error: failed to update record: %v\n", err)
		}
	}
}

func NewScheduler(interval time.Duration, ui updater.UpdaterInterface) *Scheduler {
	return &Scheduler{interval: interval, ui: ui}
}

