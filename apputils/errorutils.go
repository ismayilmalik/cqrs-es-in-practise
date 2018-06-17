package apputils

import (
	"log"
)

func HandleErrorWithFatal(err error) {
	if err != nil {
		log.Fatalf("Applicaiton will be shut down.Error occured: %v", err)
	}
}