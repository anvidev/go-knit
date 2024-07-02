package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"go-starter/internal/config"
	"go-starter/internal/store"
	"go-starter/internal/handler"
	"go-starter/internal/model"
	"go-starter/internal/service"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.MustLoad()

	db := store.Open()
	store := store.New(db)
	services := service.New(store)
	handlers := handler.New(services)

	e := echo.New()
	e.Static("/static", "static")

	if config.IsDevelopment() {
		e.Use(disableCache)
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(config.MustEnv("APP_SECRET")))))
	gob.Register(model.SessionAttributes{})

	e.GET("/", handlers.Landing.ShowLanding)
	e.GET("/sign-in", handlers.Auth.ShowSignIn)
	e.POST("/sign-in", handlers.Auth.PostSignIn)
	e.GET("/sign-up", handlers.Auth.ShowSignUp)
	e.POST("/sign-up", handlers.Auth.PostSignUp)
	e.GET("/sign-out", handlers.Auth.GetSignOut)

	e.GET("/projects", handlers.Project.ShowProjects, withAuth)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.MustEnv("SERVER_ADDR"))))
}

func disableCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-store")
		return next(c)
	}
}

func withAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		user, _ := sess.Values["user"].(model.SessionAttributes)

		if !user.LoggedIn {
			fmt.Println("THIS IS ME")
			return c.Redirect(http.StatusSeeOther, "/sign-in")
		}

		c.Set("user", user)

		ctx := context.WithValue(c.Request().Context(), "user", user)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
