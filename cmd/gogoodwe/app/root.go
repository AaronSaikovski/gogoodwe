package app

import (
	"context"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/AaronSaikovski/gogoodwe/pkg/auth"
	"github.com/alexflint/go-arg"
)

// Run is the main program runner.
//
// It takes a version string as a parameter and returns an error.
// The version string is used to set the build information.
// The function parses the command line arguments using the utils.Args struct.
// It checks if the email address and powerstation ID are in the correct format.
// If not, it fails with an error message.
// Finally, it calls the fetchData function to get data from the API and returns any errors.
//
// Parameters:
// - ctx: the context.Context object for cancellation and timeouts.
// - versionString: the version string used to set the build information.
//
// Returns:
// - error: an error if there was a problem with the input or fetching the data from the API.
func Run(ctx context.Context, versionString string) error {

	// Set version build info
	var args utils.Args
	args.SetVersion(versionString)

	// Parse args
	p := arg.MustParse(&args)

	// Check for valid email address input
	if !utils.CheckValidEmail(args.Account) {
		p.Fail("invalid email address format: should be 'user@somedomain.com'")
		return ctx.Err()

	}

	// Check for valid powerstation ID
	if !utils.CheckValidPowerstationID(args.PowerStationID) {
		p.Fail("invalid Powerstation ID format: should be 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'")
		return ctx.Err()
	}

	// User account struct instance
	apiLoginCreds := auth.NewSemsLoginCredentials(args.Account, args.Password, args.PowerStationID)

	// Get the data from the API, return any errors
	if err := loginAndfetchData(ctx, apiLoginCreds, args.ReportType); err != nil {
		return ctx.Err()
	} else {
		ctx.Done()
		return nil
	}

}
