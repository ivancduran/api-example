package ctrls

import (
	"fmt"

	"github.com/ivancduran/api-example/errs"
	"github.com/ivancduran/api-example/models"

	"github.com/labstack/echo"
)

// Return the player iframe code as a string
func GetAllPlayers(c echo.Context) (interface{}, errs.Custom) {
	// Db := db.New()
	// defer Db.Close()

	p := []models.Player{
		models.Player{
			Name: "player 1",
		},
		models.Player{
			Name: "player 2",
		},
	}

	ps := new(models.Player)

	fmt.Println(p, ps)

	// err := ps.FindAllById(Db, int(accountID), &p)
	// if err != nil {
	// 	return nil, errs.NewCustom(errs.ErrInt, errs.ErrMsg["SQL"], err)
	// }

	return p, nil

}
