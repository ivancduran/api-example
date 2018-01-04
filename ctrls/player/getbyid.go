package ctrls

import (
	"fmt"

	"github.com/ivancduran/api-example/errs"
	"github.com/ivancduran/api-example/models"

	"github.com/labstack/echo"
)

func GetPlayer(c echo.Context) (interface{}, errs.Custom) {
	// Db := db.New()
	// defer Db.Close()

	entry := c.Param("entry")

	p := new(models.Player)
	fmt.Println(entry, p)

	// err := p.FindById(Db, entry, int(accountID))
	// if err != nil {
	// 	return nil, errs.NewCustom(errs.ErrExt, errs.ErrMsg["Bind"], err)
	// }

	return p, nil
}
