package routes

import (
	"github.com/ivancduran/api-example/api"

	"github.com/labstack/echo"
)

func PlayerPublicAPI(app *echo.Group) {
	app.GET("/player/:player", api.PlayerGetIframe)
}

func PlayerPrivateAPI(app *echo.Group) {
	// all the id are entryid
	app.POST("/player", api.CreatePlayer)
	app.GET("/player", api.GetAllPlayers)
	app.GET("/player/:entry", api.GetPlayer)
	app.PUT("/player/:entry", api.UpdatePlayer)
	app.DELETE("/player/:entry", api.DeletePlayer)
}
