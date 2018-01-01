//	This file is part of Fwew.
//	Fwew is free software: you can redistribute it and/or modify
// 	it under the terms of the GNU General Public License as published by
// 	the Free Software Foundation, either version 3 of the License, or
// 	(at your option) any later version.
//
//	Fwew is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Fwew.  If not, see http://gnu.org/licenses/

// Package numbers contains all the stuff for the number parsing
package numbers

import (
	"fmt"
	"strconv"

	"github.com/tirea/fwew/util"
)

const (
	maxIntDec int64 = 32767
	maxIntOct int64 = 77777
)

// Validate range of integers for input
func valid(input int64, reverse bool) bool {
	if reverse {
		if 0 <= input && input <= maxIntDec {
			return true
		}
		return false
	}
	if 0 <= input && input <= maxIntOct {
		return true
	}
	return false
}

// Convert is the main number conversion function
func Convert(input string, reverse bool) string {
	output := ""
	if reverse {
		i, err := strconv.ParseInt(input, 10, 64)
		if !valid(i, reverse) {
			return util.Text("invalidIntError")
		}
		o := strconv.FormatInt(int64(i), 8)
		if err != nil {
			return err.Error()
		}
		output += fmt.Sprintf("Octal: %s\n", o)
		// output += fmt.Sprintf("Na'vi: %s\n", "na'vi number here")
	} else {
		io, err := strconv.ParseInt(input, 8, 64)
		if !valid(io, reverse) {
			return util.Text("invalidIntError")
		}
		d := strconv.FormatInt(int64(io), 10)
		if err != nil {
			return err.Error()
		}
		output += fmt.Sprintf("Decimal: %s\n", d)
		// output += fmt.Sprintf("Na'vi: %s\n", "na'vi number here")
	}
	return output
}
