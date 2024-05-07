package configs

import (
	"os"

	"github.com/gorilla/sessions"
)

type Session struct {
	Store *sessions.CookieStore
}

var UserSession Session

func (s Session) Init() {
	UserSession = Session{
		Store: sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY"))),
	}
}