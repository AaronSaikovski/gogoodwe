package powerstation

import (
	"errors"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/logrusorgru/aurora"
	"github.com/valyala/fastjson"
)

// parseOutput parses the JSON output from the provided byte slice.
//
// It takes a byte slice as input and returns a *fastjson.Value and an error.
func parseOutput(dataOutput []byte) (*fastjson.Value, error) {
	//parse JSON output
	var parser fastjson.Parser
	output, err := parser.Parse(string(dataOutput))
	if err != nil {
		utils.HandleError(errors.New("error: parsing powerstation data"))
		return nil, err
	}

	return output, nil

}

func printOutput(output *fastjson.Value) {
	fmt.Println(aurora.BrightYellow(output))
}
