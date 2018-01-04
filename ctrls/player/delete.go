package ctrls

import (
	"fmt"

	"github.com/ivancduran/api-example/errs"
	"github.com/ivancduran/api-example/models"

	"github.com/labstack/echo"
)

func DeletePlayer(c echo.Context) (interface{}, errs.Custom) {
	// Db := db.New()
	// defer Db.Close()

	entry := c.Param("entry")

	player := new(models.Player)

	fmt.Println(entry, player)

	// err := player.Delete(Db, entry, accountID)
	// if err != nil {
	// 	return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["SQL"], err)
	// }

	res := map[string]interface{}{
		"message": "player deleted",
		"success": true,
	}

	return res, nil
}
