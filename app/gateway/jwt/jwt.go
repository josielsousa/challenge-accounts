package jwt

import "time"

type Jwt struct {
	appKey []byte
	C      Clocker
}

//go:generate moq -rm -out jwt_mock.go . Clocker
type Clocker interface {
	Now() time.Time
	Until(t time.Time) time.Duration
}

func New(appKey []byte) *Jwt {
	return &Jwt{
		appKey: appKey,
		C:      NewClock(),
	}
}
