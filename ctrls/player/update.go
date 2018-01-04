package ctrls

import (
	"encoding/json"
	"fmt"

	"github.com/ivancduran/api-example/errs"
	"github.com/ivancduran/api-example/models"

	"github.com/labstack/echo"
)

// UpdatePlayer
func UpdatePlayer(c echo.Context) (interface{}, errs.Custom) {
	// Db := db.New()
	// defer Db.Close()

	entry := c.Param("entry")

	p := new(models.Player)

	fmt.Println(entry, p)

	b := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&b)
	if err != nil {
		return nil, errs.NewCustom(errs.ErrExt, errs.ErrMsg["Bind"], err)
	}

	// err = p.Update(Db, entry, accountID, b)
	// if err != nil {
	// 	return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["SQL"], err)
	// }

	res := map[string]interface{}{
		"msg":     "player updated",
		"success": true,
	}

	return res, nil
}
