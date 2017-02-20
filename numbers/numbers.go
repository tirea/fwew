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

// This util library contains all the stuff for the number parsing
package numbers

const (
	maxIntDec int = 32767
	maxIntOct int = 77777
)

func valid(input int, reverse bool) bool {
	if reverse {
		if 0 <= input && input <= maxIntDec {
			return true
		} else {
			return false
		}
	} else {
		if 0 <= input && input <= maxIntOct {
			return true
		} else {
			return false
		}
	}
}

func Convert(input int, reverse bool) {
	if !valid(input, reverse) {

	}
}
