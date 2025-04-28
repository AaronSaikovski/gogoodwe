package utils

import (
	"log"

	"github.com/logrusorgru/aurora"
)

// HandleError - Generic error handler.
//
// Parameters:
// - err: the error to handle.
func HandleError(err error) {
	log.Fatal(aurora.BrightRed(err.Error()))
}
