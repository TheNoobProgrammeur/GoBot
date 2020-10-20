package commands

import (
	"errors"
	"myTest/src/bot/request"
	"regexp"
)

func Jocke(route string) (string, error) {

	regRoute := regexp.MustCompile("^Any|Programming|Miscellaneous|Darck|Pun$")

	if regRoute.MatchString(route) {
		rep, _ := request.GetJock(route)
		return rep, nil
	} else {
		return "error", errors.New("Parametre non valide")
	}
}
