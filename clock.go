package _65go2

import (
	"time"
)

type Clock struct {
	rate     time.Duration
	divisor  int
	ticks    uint64
	ticker   *time.Ticker
	stopChan chan int
	waiting  map[uint64][]chan int
}

func NewClock(rate time.Duration, divisor int) Clock {
	return Clock{
		rate:     rate,
		divisor:  divisor,
		ticks:    0,
		ticker:   nil,
		stopChan: make(chan int, 1),
		waiting:  make(map[uint64][]chan int),
	}
}

func (clock *Clock) maintainTime() {
	ticks := 0

	for {
		select {
		case <-clock.stopChan:
			clock.ticker = nil
			return
		case _ = <-clock.ticker.C:
			ticks++

			if ticks == clock.divisor {
				clock.ticks++
				ticks = 0

				if Ca, ok := clock.waiting[clock.ticks]; ok {
					for _, C := range Ca {
						C <- 1
					}

					delete(clock.waiting, clock.ticks)
				}
			}
		}
	}
}

func (clock *Clock) start() {
	if clock.ticker == nil {
		clock.ticker = time.NewTicker(clock.rate)
		clock.maintainTime()
	}
}

func (clock *Clock) stop() {
	if clock.ticker != nil {
		clock.stopChan <- 1
	}
}

func (clock *Clock) await(tick uint64) {
	if clock.ticks >= tick {
		return
	}

	C := make(chan int, 1)
	clock.waiting[tick] = append(clock.waiting[tick], C)
	<-C
}
