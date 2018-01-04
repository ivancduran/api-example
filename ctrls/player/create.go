package ctrls

import (
	"github.com/ivancduran/api-example/errs"
	"github.com/ivancduran/api-example/generics"
	"github.com/ivancduran/api-example/models"

	"github.com/labstack/echo"
)

// Return the player iframe code as a string
func CreatePlayer(c echo.Context) (interface{}, errs.Custom) {
	// Db := db.New()
	// defer Db.Close()

	p := new(models.Player)
	err := c.Bind(p)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrUsr, errs.ErrMsg["Bind"], err)
	}

	p.Entry, _ = generics.Entry()

	res := map[string]interface{}{
		"message": "player created correctly",
		"player":  p,
		"success": true,
	}

	return res, nil

}
