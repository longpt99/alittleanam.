package configs

import (
	"net/http"

	"ala-coffee-notification/controllers"
	"ala-coffee-notification/utils"
	"ala-coffee-notification/utils/errs"
	"ala-coffee-notification/utils/router"
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
