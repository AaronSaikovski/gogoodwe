package utils

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/valyala/fastjson"
)

var parserPool fastjson.ParserPool

// ParseOutput parses the JSON output from the provided byte slice using a pooled parser.
func ParseOutput(dataOutput []byte) (*fastjson.Value, error) {
	p := parserPool.Get()
	defer parserPool.Put(p)
	return p.ParseBytes(dataOutput)
}

// printOutput prints the provided fastjson.Value in bright yellow color using the aurora package.
//
// Parameters:
// - output: a pointer to a fastjson.Value that represents the output to be printed.
//
// Return type: None.
func PrintOutput(output *fastjson.Value) {
	fmt.Println(aurora.BrightYellow(output))
}
