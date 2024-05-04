package controllers

import (
	"ala-coffee-search/utils/router"

	"ala-coffee-search/controllers/product"
)

func InitControllers(r *router.Bundle) {
	product.NewController(r)
}
