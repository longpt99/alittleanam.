package auth

import (
	"fmt"
	"net/http"

	"ala-coffee-search/models"
	"ala-coffee-search/repositories"
	"ala-coffee-search/services"
	"ala-coffee-search/utils"
	"ala-coffee-search/utils/router"
)

type Controller struct {
	userService *services.UserService
}

func NewController(r *router.Bundle, repo repositories.Repository) {
	c := &Controller{
		userService: &services.UserService{
			UserRepo: repo.UserRepo,
		},
	}

	r.Mount("/auth").Route(func(authGroup *router.Bundle) {
		authGroup.HandleFunc("POST /sign-in", c.handleSignIn)
		authGroup.HandleFunc("POST /sign-up", c.handleSignUp)
		authGroup.HandleFunc("POST /forgot-password", c.handleForgotPassword)
		authGroup.HandleFunc("POST /reset-password", c.handleResetPassword)
	})
}

func (c *Controller) handleSignUp(w http.ResponseWriter, r *http.Request) {
	var body models.SignUpAccount

	if err := utils.ReadValid(&body, r); err != nil {
		utils.WriteError(w, r, err)
		return
	}

	id, err := c.userService.CreateAccount(&body)
	if err != nil {
		utils.WriteError(w, r, err)
		return
	}

	utils.Write(w, r, id)
}

func (c *Controller) handleSignIn(w http.ResponseWriter, r *http.Request) {
	var body models.SignInAccount

	if err := utils.ReadValid(&body, r); err != nil {
		utils.WriteError(w, r, err)
		return
	}

	id, err := c.userService.SignInAccount(&body)
	if err != nil {
		// Handle validation errors
		utils.WriteError(w, r, err)
		return
	}

	utils.Write(w, r, id)
}

func (c *Controller) handleForgotPassword(w http.ResponseWriter, r *http.Request) {
	var body models.SignUpAccount

	if err := utils.ReadValid(&body, r); err != nil {
		utils.WriteError(w, r, err)
		return
	}

	fmt.Println(body)

	id, err := c.userService.UserRepo.InsertOne(&body)

	// result, err := user.CreateAccount(&body)
	if err != nil {
		// Handle validation errors
		utils.WriteError(w, r, err)
		return
	}

	utils.Write(w, r, id)
}

func (c *Controller) handleResetPassword(w http.ResponseWriter, r *http.Request) {
	var body models.SignUpAccount

	if err := utils.ReadValid(&body, r); err != nil {
		utils.WriteError(w, r, err)
		return
	}

	fmt.Println(body)

	id, err := c.userService.UserRepo.InsertOne(&body)

	// result, err := user.CreateAccount(&body)
	if err != nil {
		// Handle validation errors
		utils.WriteError(w, r, err)
		return
	}

	utils.Write(w, r, id)
}
