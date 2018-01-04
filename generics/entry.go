package generics

import (
	randmin "math/rand"
	"time"

	hashids "github.com/speps/go-hashids"
)

// Generate uniq random string
func Entry() (string, error) {

	hd := hashids.NewData()
	hd.Salt = "defrgtcdvfbgswdexsfrdesdws12swko"
	hd.Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hd.MinLength = 15

	h, err := hashids.NewWithData(hd)
	timeint := int(time.Now().UnixNano())
	randint := randmin.Intn(9999)

	e, err := h.Encode([]int{timeint, randint})

	if err != nil {
		return "", err
	}

	return e, nil
}
