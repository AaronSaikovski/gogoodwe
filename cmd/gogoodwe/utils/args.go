package utils

var (
	// Version string
	VersionString string = "v0.0.1"
	infoString    string = "GoGoodwe - A CLI tool to query your SEMS Solar Inverter API."
	reportString  string = "Report Types: (0)-Detail, (1)-Summary, (2)-Point, (3)-Plant, (4)-PlantChart, (5)-PowerFlow"
)

// Args - struct using go-arg- https://github.com/alexflint/go-arg
type Args struct {
	Account        string `arg:"required,-a,--account" help:"SEMS Email Account."`
	Password       string `arg:"required,-p,--password" help:"SEMS Account password."`
	PowerStationID string `arg:"required,-i,--powerstationid" help:"SEMS Powerstation ID."`
	ReportType     int    `arg:",-r,--reporttype" help:"Inverter Report Number"`
}

// Description returns a command line tool to query the GOODWE SEMS Portal APIs and Solar SEMS API.
//
// No parameters.
// Returns a string.
func (Args) Description() string {
	return infoString + "\n" + reportString
}

// Version returns the version string of the Args struct.
//
// No parameters.
// Returns a string.
func (Args) Version() string {
	return "version: " + VersionString
}

// SetVersion sets the version string of the Args struct.
//
// versionString: the new version string to set.
func (Args) SetVersion(versionString string) {
	VersionString = versionString
}
