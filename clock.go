package go6502

import (
	"time"
)

type Clock struct {
	rate     time.Duration
	ticks    uint64
	ticker   *time.Ticker
	stopChan chan int
	waiting  map[uint64][]chan int
}

func NewClock(rate time.Duration) *Clock {
	return &Clock{
		rate:     rate,
		ticks:    0,
		ticker:   nil,
		stopChan: make(chan int),
		waiting:  make(map[uint64][]chan int),
	}
}

func (clock *Clock) maintainTime() {
	for {
		select {
		case <-clock.stopChan:
			clock.ticker = nil
			return
		case _ = <-clock.ticker.C:
			clock.ticks++

			if Ca, ok := clock.waiting[clock.ticks]; ok {
				for _, C := range Ca {
					C <- 1
				}

				delete(clock.waiting, clock.ticks)
			}
		}
	}
}

func (clock *Clock) Ticks() uint64 {
	return clock.ticks
}

func (clock *Clock) Start() (ticks uint64) {
	ticks = clock.ticks

	if clock.ticker == nil {
		clock.ticker = time.NewTicker(clock.rate)
		clock.maintainTime()
	}

	return
}

func (clock *Clock) Stop() {
	if clock.ticker != nil {
		clock.stopChan <- 1
	}
}

func (clock *Clock) Await(tick uint64) uint64 {
	if clock.ticks < tick {
		C := make(chan int, 1)
		clock.waiting[tick] = append(clock.waiting[tick], C)
		<-C
	}

	return clock.ticks
}
