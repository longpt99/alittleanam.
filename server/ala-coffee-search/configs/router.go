package configs

import (
	"net/http"

	"ala-coffee-search/controllers"
	"ala-coffee-search/utils"
	"ala-coffee-search/utils/errs"
	"ala-coffee-search/utils/router"
)

func RouterConfig(r *router.Bundle) {
	controllers.InitControllers(r)

	// r.Use(middleware.Logger)
	// r.Get("/documentation/*", httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost:8080/documentation/doc.json"), // The Swagger JSON endpoint URL
	// ))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteError(w, r, errs.E(http.StatusNotFound))
	})
}
