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

import "github.com/tirea/fwew/util"

const (
	maxIntDec int = 32767
	maxIntOct int = 77777
)

func valid(input int, reverse bool) bool {
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
func Convert(input int, reverse bool) string {
	if !valid(input, reverse) {
		return util.Text("invalidIntError")
	}
	// convert! :D
	return "" // TODO
}
