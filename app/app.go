package app

// Main package - This is the main program entry point
import (
	"github.com/AaronSaikovski/gogoodwe/inverter"
	"github.com/AaronSaikovski/gogoodwe/utils"
	"github.com/alexflint/go-arg"
)

// Run - main program runner
func Run() error {

	//Get the args input data
	var args utils.Args
	p := arg.MustParse(&args)

	//check for valid email address input
	if !utils.CheckValidEmail(args.Account) {
		p.Fail("Invalid Email address format - should be: 'user@somedomain.com'.")
	}

	//check for valid powerstation Id
	if !utils.CheckValidPowerstationID(args.PowerStationID) {
		p.Fail("Invalid Powerstation ID format: - should be: 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'.")
	}

	// Get the data from the API, return any errors. Pass in args as string
	return inverter.FetchData(args.Account, args.Password, args.PowerStationID)
}
