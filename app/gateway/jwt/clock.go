package jwt

import "time"

type Clock struct{}

func NewClock() Clock {
	return Clock{}
}

func (s Clock) Now() time.Time {
	return time.Now()
}

func (s Clock) Until(t time.Time) time.Duration {
	return time.Until(t)
}
