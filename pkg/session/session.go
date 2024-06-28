package session

import (
	"go-starter/internal/model"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	SessionCookieKey = "session"
)

func Create(c echo.Context, sessAttr *model.SessionAttributes) error {
	sess, _ := session.Get(SessionCookieKey, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values[model.SessionUserKey] = sessAttr
	if err := sess.Save(c.Request(), c.Response().Writer); err != nil {
		return err
	}
	return nil
}

func Destroy(c echo.Context) error {
	sess, _ := session.Get(SessionCookieKey, c)
	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response().Writer); err != nil {
		return err
	}
	return nil
}
