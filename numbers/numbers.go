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
	"strings"

	"github.com/tirea/fwew/util"
)

const (
	maxIntDec int64 = 32767
	maxIntOct int64 = 77777
)

var naviVocab = [][]string{
	// 0 1 2 3 4 5 6 7 actual
	{"", "'aw", "mune", "pxey", "tsìng", "mrr", "pukap", "kinä"},
	// 0 1 2 3 4 5 6 7 last digit
	{"", "aw", "mun", "pey", "sìng", "mrr", "fu", "hin"},
	// 0 1 2 3 4 5 6 7 first or middle digit
	{"", "", "me", "pxe", "tsì", "mrr", "pu", "ki"},
	// 0 1 2 3 4 powers of 8
	{"", "vo", "zam", "vozam", "zazam"},
	// 0 1 2 3 4 powers of 8 last digit
	{"", "l", "", "", ""},
}

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

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func wordify(input string) string {
	rev := reverse(input)
	output := ""
	if len(input) == 1 {
		if input == "0" {
			return "kew"
		}
		inty, _ := strconv.Atoi(input)
		return naviVocab[0][inty]
	}
	for i, d := range rev {
		switch i {
		case 0: // 7777[7]
			output = naviVocab[1][int(d-'0')] + output
			if int(d-'0') == 1 && rev[1] != '0' {
				output = naviVocab[4][1] + output
			}
		case 1: // 777[7]7
			if int(d-'0') > 0 {
				output = naviVocab[2][int(d-'0')] + naviVocab[3][1] + output
			}
		case 2: // 77[7]77
			if int(d-'0') > 0 {
				output = naviVocab[2][int(d-'0')] + naviVocab[3][2] + output
			}
		case 3: // 7[7]777
			if int(d-'0') > 0 {
				output = naviVocab[2][int(d-'0')] + naviVocab[3][3] + output
			}
		case 4: // [7]7777
			if int(d-'0') > 0 {
				output = naviVocab[2][int(d-'0')] + naviVocab[3][4] + output
			}
		}
	}
	for _, d := range []string{"01", "02", "03", "04", "05", "06", "07"} {
		if rev[0:2] == d {
			output = output + naviVocab[4][1]
		}
	}
	output = strings.Replace(output, "mm", "m", -1)
	return output
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
		output += fmt.Sprintf("Na'vi: %s\n", wordify(o))
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
		output += fmt.Sprintf("Na'vi: %s\n", wordify(input))
	}
	return output
}
