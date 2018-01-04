package api

import (
	"net/http"

	ctrls "github.com/ivancduran/api-example/ctrls/player"

	"github.com/ivancduran/api-example/errs"

	"github.com/labstack/echo"
)

// PlayerGetIframe Return the player
func PlayerGetIframe(c echo.Context) error {
	res, err := ctrls.PlayerGetIframe(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

// CreatePlayer Create a new player
func CreatePlayer(c echo.Context) error {
	res, err := ctrls.CreatePlayer(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

// GetAllPlayers Get all players by accountid
func GetAllPlayers(c echo.Context) error {
	res, err := ctrls.GetAllPlayers(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

// GetPlayer Get player information based on the id
func GetPlayer(c echo.Context) error {
	res, err := ctrls.GetPlayer(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

// UpdatePlayer desc
func UpdatePlayer(c echo.Context) error {
	res, err := ctrls.UpdatePlayer(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}

// DeletePlayer delete file
func DeletePlayer(c echo.Context) error {
	res, err := ctrls.DeletePlayer(c)
	if err != nil {
		return errs.Send(err.Result())
	}

	return c.JSON(http.StatusOK, res)
}
