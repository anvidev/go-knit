package handler

import (
	"errors"
	"go-starter/internal/model"
	"go-starter/internal/service"
	"go-starter/internal/view/auth"
	"go-starter/pkg/argon2"
	"go-starter/pkg/session"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	InternalServerError     = errors.New("Internal server error")
	InvalidCredentialsError = errors.New("Invalid credentials")
	UserAlreadyExistError   = errors.New("A user with that email already exist")
	IncompleteFormError     = errors.New("Please fill out the form")
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(s service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: s,
	}
}

func (h AuthHandler) ShowSignIn(c echo.Context) error {
	return render(c, auth.ShowSignIn("Good to see you again"))
}

func (h AuthHandler) PostSignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if len(email) == 0 || len(password) == 0 {
		return render(c, auth.FormSignIn(auth.FormSignInData{
			Message: IncompleteFormError.Error(),
		}))
	}

	user, err := h.userService.GetByEmail(email)
	if err != nil {
		return render(c, auth.FormSignIn(auth.FormSignInData{
			Values: map[string]string{
				"email":    email,
				"password": password,
			},
			Message: InvalidCredentialsError.Error(),
		}))
	}

	valid, err := argon2.Compare(password, user.HashedPassword)
	if err != nil {
		return render(c, auth.FormSignIn(auth.FormSignInData{
			Values: map[string]string{
				"email":    email,
				"password": password,
			},
			Message: InternalServerError.Error(),
		}))
	}
	if !valid {
		return render(c, auth.FormSignIn(auth.FormSignInData{
			Values: map[string]string{
				"email":    email,
				"password": password,
			},
			Message: InvalidCredentialsError.Error(),
		}))
	}

	sessAtrr := model.SessionAttributes{
		Name:     user.Name,
		Email:    user.Email,
		LoggedIn: true,
	}
	if err := session.Create(c, &sessAtrr); err != nil {
		return err
	}

	return hxRedirect(c, "/projects")
}

func (h AuthHandler) ShowSignUp(c echo.Context) error {
	return render(c, auth.ShowSignUp("Welcome!"))
}

func (h AuthHandler) PostSignUp(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if len(name) == 0 || len(email) == 0 || len(password) == 0 {
		return render(c, auth.FormSignUp(auth.FormSignUpData{
			Message: IncompleteFormError.Error(),
		}))
	}

	hashedPassword, err := argon2.Hash(password)
	if err != nil {
		slog.Error("PostSignUp, ", "Password hashing error:", err)
		return render(c, auth.FormSignUp(auth.FormSignUpData{
			Values: map[string]string{
				"name":     name,
				"email":    email,
				"password": password,
			},
			Message: InternalServerError.Error(),
		}))
	}

	newUser, err := h.userService.Create(name, email, hashedPassword)
	if err != nil {
		slog.Error("PostSignUp, ", "Creating user error: ", err)
		return render(c, auth.FormSignUp(auth.FormSignUpData{
			Values: map[string]string{
				"name":     name,
				"email":    email,
				"password": password,
			},
			Message: UserAlreadyExistError.Error(),
		}))
	}

	sessAttr := model.SessionAttributes{
		Name:     newUser.Name,
		Email:    newUser.Email,
		LoggedIn: true,
	}
	if err := session.Create(c, &sessAttr); err != nil {
		slog.Error("PostSignUp, ", "Saving session error: ", err)
		return render(c, auth.FormSignUp(auth.FormSignUpData{
			Values: map[string]string{
				"name":     name,
				"email":    email,
				"password": password,
			},
			Message: InternalServerError.Error(),
		}))
	}

	return hxRedirect(c, "/projects")
}

func (h AuthHandler) GetSignOut(c echo.Context) error {
	if err := session.Destroy(c); err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/sign-in")
}
