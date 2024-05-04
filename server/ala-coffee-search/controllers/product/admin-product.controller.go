package product

import (
	"fmt"
	"net/http"

	"ala-coffee-search/models"
	"ala-coffee-search/repositories"
	"ala-coffee-search/services"
	"ala-coffee-search/utils"
	"ala-coffee-search/utils/router"
)

type AdminController struct {
	userService *services.UserService
}

func NewAdminController(r *router.Bundle, repo repositories.Repository) {
	c := &AdminController{
		userService: &services.UserService{
			UserRepo: repo.UserRepo,
		},
	}

	r.Mount("/admin/products").Route(func(adminProductGroup *router.Bundle) {
		adminProductGroup.HandleFunc("POST /sign-up", c.HandleSignUp)
	})
}

func (c *AdminController) HandleSignUp(w http.ResponseWriter, r *http.Request) {
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
