package routes

import (
	"github.com/ivancduran/api-example/generics"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ApiRoutes(app *echo.Group) {
	// api := app.Group("/private")
	PlayerPublicAPI(app)

	// private
	private := makeRestrictedRoutes(app)
	PlayerPrivateAPI(private)
}

func WebhookRoutes() {

}

func makeRestrictedRoutes(app *echo.Group) *echo.Group {
	// restricted group
	j := new(generics.JWT)
	restricted := app.Group("/private")
	restricted.Use(middleware.JWT(j.Secret()))

	return restricted
}
