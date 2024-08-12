package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)
	mux.Get("/", app.HomePage)
	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Post("/activate-account", app.ActivateAccount)

	mux.Get("/test-email", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HERE................")

		errchan := make(chan error)
		m := Mail{
			Domain:      "localhost",
			Host:        "localhost",
			Port:        1025,
			Encryption:  "none",
			FromAddress: "info@mycompany.com",
			FromName:    "info",
			ErrorChan:   errchan,
		}

		msg := Message{
			To:      "me@here.com",
			Subject: "Test mail",
			Data:    "Hello, world",
		}

		m.sendMail(msg, make(chan error))
	})

	return mux
}
