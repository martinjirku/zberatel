package session

import (
	"time"

	"github.com/alexedwards/scs/v2"
)

func NewSessionManager() *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	return sessionManager
}
