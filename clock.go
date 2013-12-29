package _65go2

import (
	"time"
)

type Clock struct {
	rate    time.Duration
	divisor int
	ticker  *time.Ticker
}

func NewClock(rate time.Duration, divisor int) Clock {
	return Clock{rate: rate, divisor: divisor, ticker: nil}
}

func (clock *Clock) start() <-chan time.Time {
	if clock.ticker == nil {
		clock.ticker = time.NewTicker(clock.rate)
		return clock.ticker.C
	}

	return nil
}

func (clock *Clock) stop() {
	if clock.ticker != nil {
		clock.ticker.Stop()
	}

	clock.ticker = nil
}
