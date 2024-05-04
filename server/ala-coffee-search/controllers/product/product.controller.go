package product

import (
	"net/http"

	"ala-coffee-search/services"
	"ala-coffee-search/utils/router"
)

type Controller struct {
	productService *services.UserService
}

func NewController(r *router.Bundle) {
	c := &Controller{
		// productService: &services.UserService{
		// 	UserRepo: repo.UserRepo,
		// },
	}

	r.Mount("/products").Route(func(productGroup *router.Bundle) {
		productGroup.HandleFunc("GET /{$}", c.GetProducts)
	})
}

func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	// var body models.SignUpAccount

	// if err := utils.ReadValid(&body, r); err != nil {
	// 	utils.WriteError(w, r, err)
	// 	return
	// }

	// fmt.Println(body)

	// id, err := c.userService.UserRepo.InsertOne(&body)

	// if err != nil {
	// 	// Handle validation errors
	// 	utils.WriteError(w, r, err)
	// 	return
	// }

	// utils.Write(w, r, id)
}
