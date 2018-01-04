package ctrls

import (
	"github.com/ivancduran/api-example/errs"

	"github.com/labstack/echo"
)

func PlayerGetIframe(c echo.Context) (interface{}, errs.Custom) {

	player := c.Param("player")
	params := map[string]interface{}{
		"player": player,
	}

	return params, nil

}
