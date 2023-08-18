package utils

import (
	"log"

	"github.com/logrusorgru/aurora"
)

// HandleError - Generic error handler
func HandleError(err error) {
	log.Fatal(aurora.BrightRed(err.Error()))
}