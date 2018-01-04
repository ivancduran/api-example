package errs

import (
	"fmt"
	"net/http"

	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var App_Env bool

func init() {
	if os.Getenv("APP_ENV") == "production" {
		App_Env = true
		raven.SetDSN(os.Getenv("SENTRY_DSN"))
	}
}

const (
	ErrUsr = "user"
	ErrExt = "external"
	ErrInt = "internal"
)

// ErrMsg custom message erros
var ErrMsg = map[string]string{
	"SQL":           "SQl syntax problem",
	"Query":         "Query syntax",
	"Bind":          "Can't bind data",
	"Write":         "Can't write file",
	"AssetSelect":   "Can't select video entry",
	"IDne":          "ID don't exist",
	"Upload":        "Can't upload file",
	"BadInfo":       "Some data is not provided",
	"DecodingToken": "Can't decode token",
	"UpdatePlan":    "Can't update account's plan",
	"Deletechannel": "Can't delete instance",
	"SendEmail":     "Can't send email",
	"Mongo":         "Can't connect to mongodb",
	"UpdateAuthKey": "Can't update auth key",
	"Login":         "Email/password incorrect",
	"Cancel":        "Account has been canceled, please send email to contact@easycast.com",
	"Account":       "Account inactive",
	"EmailTaken":    "Email already taken",
}

func Send(status string, msg string, err error) *echo.HTTPError {
	m := msg
	var c int

	switch status {
	case "user":
		c = http.StatusBadRequest
		m += ", verify your data."
		log.Warn("An External error occured." + msg)
		break
	case "external":
		c = http.StatusInternalServerError
		m += " Something went wrong, please contact you're administrator."
		log.Warn("An External error occured." + msg)
		break
	case "internal":
		c = http.StatusNotAcceptable
		log.Warn("An Internal error occured." + msg)
		break
	}

	if App_Env == true {
		raven.CaptureError(err, nil)
	}

	fmt.Println("-------------------------------------")
	fmt.Println("ERROR => ")
	log.Error(err)
	fmt.Println("-------------------------------------")
	fmt.Println("Sending Echo Error: " + m)
	fmt.Println("-------------------------------------")

	return echo.NewHTTPError(c, m)
}
