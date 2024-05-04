package user

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

	r.Mount("/users").Route(func(userGroup *router.Bundle) {
		userGroup.HandleFunc("GET /me", c.handleGetMyProfile)
		userGroup.HandleFunc("PUT /me", c.handleUpdateMyProfile)
	})
}

func (c *Controller) handleGetMyProfile(w http.ResponseWriter, r *http.Request) {
	var body models.SignUpAccount

	if err := utils.ReadValid(&body, r); err != nil {
		utils.WriteError(w, r, err)
		return
	}

	fmt.Println(body)

	id, err := c.userService.UserRepo.InsertOne(&body)
	if err != nil {
		// Handle validation errors
		utils.WriteError(w, r, err)
		return
	}

	utils.Write(w, r, id)
}

func (c *Controller) handleUpdateMyProfile(w http.ResponseWriter, r *http.Request) {
	var body models.SignUpAccount

	if err := utils.ReadValid(&body, r); err != nil {
		utils.WriteError(w, r, err)
		return
	}

	id, err := c.userService.UserRepo.InsertOne(&body)

	if err != nil {
		// Handle validation errors
		utils.WriteError(w, r, err)
		return
	}

	utils.Write(w, r, id)
}
