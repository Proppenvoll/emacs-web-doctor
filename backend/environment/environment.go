package environment

import (
	"errors"
	"log"
	"os"
)

func convertYesOrNoToBool(yesOrNo string) (bool, error) {
	switch yesOrNo {
	case "yes":
		return true, nil
	case "no":
		return false, nil
	}

	return false, errors.New("Provided string needs to be 'yes' or 'no'")
}

var CorsActive bool
var Port string

func init() {
	var corsActiveError error
	CorsActive, corsActiveError = convertYesOrNoToBool(os.Getenv("CORS_ACTIVE"))

	if corsActiveError != nil {
		log.Fatalln("Environment variable CORS_ACTIVE needs a value of 'yes' or 'no'")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		log.Fatalln("Missing environment variable PORT")
	}
}
