package m65go2

import (
	"time"
)

// Represents a clock signal for an IC.  Once a Clock is started, it
// maintains a 'ticks' counters which is incremented at a specific
// interval.
type Clock struct {
	rate     time.Duration
	ticks    uint64
	ticker   *time.Ticker
	stopChan chan int
	waiting  map[uint64][]chan int
}

// Returns a pointer to a new Clock which increments its ticker at
// intervals of 'rate'.  The returned Clock has not been started and
// its ticks counter is zero.
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

// Returns the current value of the Clock's ticks counter.
func (clock *Clock) Ticks() uint64 {
	return clock.ticks
}

// Starts the Clock
func (clock *Clock) Start() (ticks uint64) {
	ticks = clock.ticks

	if clock.ticker == nil {
		clock.ticker = time.NewTicker(clock.rate)
		clock.maintainTime()
	}

	return
}

// Stops the clock
func (clock *Clock) Stop() {
	if clock.ticker != nil {
		clock.stopChan <- 1
	}
}

// Blocks the calling thread until the given tick has arrived.
// Returns immediately if the clock has already passed the given tick.
func (clock *Clock) Await(tick uint64) uint64 {
	if clock.ticks < tick {
		C := make(chan int, 1)
		clock.waiting[tick] = append(clock.waiting[tick], C)
		<-C
	}

	return clock.ticks
}
