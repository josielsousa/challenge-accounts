package jwt

type Jwt struct {
	appKey []byte
}

func New(appKey []byte) *Jwt {
	return &Jwt{
		appKey: appKey,
	}
}
