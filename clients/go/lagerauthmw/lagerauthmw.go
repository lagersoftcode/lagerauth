package lagerauthmw

// LagerAuthMiddleware is the main strcut on to which methods are attached
type LagerAuthMiddleware struct {
	config //embed config struct
}

func New(clientID, secretKey string, opts ...Opt) LagerAuthMiddleware {
	lmw := LagerAuthMiddleware{
		config{
			ClientID:   clientID,
			SecretKey:  secretKey,
			OAuthURL:   "https://oauth.lagersoft.com",
			CookieName: "lmwauth",
			MountPoint: "/oauth",
		},
	}

	// set options
	for _, f := range opts {
		f(&lmw)
	}

	return lmw
}

// Opt(ional) argument following the functional Opts pattern:
type Opt func(*LagerAuthMiddleware)

func OAuthURLOpt(url string) Opt {
	return func(lmw *LagerAuthMiddleware) {
		lmw.OAuthURL = url
	}
}

func CookieNameOpt(cookie string) Opt {
	return func(lmw *LagerAuthMiddleware) {
		lmw.CookieName = cookie
	}
}

func MountPointOpt(mountPoint string) Opt {
	return func(lmw *LagerAuthMiddleware) {
		lmw.MountPoint = mountPoint
	}
}

type config struct {
	ClientID   string
	SecretKey  string
	OAuthURL   string
	CookieName string
	MountPoint string
}
