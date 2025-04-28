/*
MIT License

# Copyright (c) 2024 Aaron Saikovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package powerflow

type Powerflow struct {
	Language string `json:"language"`
	Function any    `json:"function"`
	HasError bool   `json:"hasError"`
	Msg      string `json:"msg"`
	Code     string `json:"code"`
	Data     struct {
		HasGenset                bool `json:"hasGenset"`
		HasMoreInverter          bool `json:"hasMoreInverter"`
		HasPowerflow             bool `json:"hasPowerflow"`
		Powerflow                any  `json:"powerflow"`
		HasGridLoad              bool `json:"hasGridLoad"`
		IsStored                 bool `json:"isStored"`
		IsParallelInventers      bool `json:"isParallelInventers"`
		IsMixedParallelInventers bool `json:"isMixedParallelInventers"`
		IsEvCharge               bool `json:"isEvCharge"`
		EvCharge                 any  `json:"evCharge"`
	} `json:"data"`
	Components struct {
		Para         string `json:"para"`
		LangVer      int    `json:"langVer"`
		TimeSpan     int    `json:"timeSpan"`
		API          string `json:"api"`
		MsgSocketAdr any    `json:"msgSocketAdr"`
	} `json:"components"`
}

func NewPowerflow() *Powerflow {
	return &Powerflow{}
}
