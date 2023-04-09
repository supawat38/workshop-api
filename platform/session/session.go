package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	SessionStore *session.Store
)

func InitSession() {
	var ConfigDefault = session.Config{
		Expiration:   24 * time.Hour,
		CookieName:   "gamemarket",
		KeyGenerator: utils.UUID,
	}
	SessionStore = session.New(ConfigDefault)
}
