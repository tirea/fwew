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

// Package main contains all the things. numbers.go contains all the stuff for the number parsing
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

// "word number portion": octal value
var numTable = map[string]int{
	"kizazam":  070000,
	"kizaza":   070000,
	"puzazam":  060000,
	"puzaza":   060000,
	"mrrzazam": 050000,
	"mrrzaza":  050000,
	"rrzazam":  050000,
	"rrzaza":   050000,
	"tsìzazam": 040000,
	"tsìzaza":  040000,
	"pxezazam": 030000,
	"pxezaza":  030000,
	"mezazam":  020000,
	"mezaza":   020000,
	"ezazam":   020000,
	"ezaza":    020000,
	"zazam":    010000,
	"zaza":     010000,
	"kivozam":  07000,
	"kivoza":   07000,
	"puvozam":  06000,
	"puvoza":   06000,
	"mrrvozam": 05000,
	"mrrvoza":  05000,
	"rrvozam":  05000,
	"rrvoza":   05000,
	"tsìvozam": 04000,
	"tsìvoza":  04000,
	"pxevozam": 03000,
	"pxevoza":  03000,
	"mevozam":  02000,
	"mevoza":   02000,
	"evozam":   02000,
	"evoza":    02000,
	"vozam":    01000,
	"voza":     01000,
	"kizam":    0700,
	"kiza":     0700,
	"puzam":    0600,
	"puza":     0600,
	"mrrzam":   0500,
	"mrrza":    0500,
	"rrzam":    0500,
	"rrza":     0500,
	"tsìzam":   0400,
	"tsìza":    0400,
	"pxezam":   0300,
	"pxeza":    0300,
	"mezam":    0200,
	"meza":     0200,
	"ezam":     0200,
	"eza":      0200,
	"zam":      0100,
	"za":       0100,
	"kivol":    070,
	"kivo":     070,
	"puvol":    060,
	"puvo":     060,
	"mrrvol":   050,
	"mrrvo":    050,
	"rrvol":    050,
	"rrvo":     050,
	"tsìvol":   040,
	"tsìvo":    040,
	"pxevol":   030,
	"pxevo":    030,
	"mevol":    020,
	"mevo":     020,
	"evol":     020,
	"evo":      020,
	"vol":      010,
	"vo":       010,
	"hin":      07,
	"fu":       06,
	"mrr":      05,
	"rr":       05,
	"sìng":     04,
	"pey":      03,
	"mun":      02,
	"un":       02,
	"aw":       01,
}

func unwordify(input string) string {
	var (
		matchNumbers []string
		re           *regexp.Regexp
		reString     string
		s            string
		n            int
	)
	s = strings.ToLower(input)
	// kew
	if s == "kew" {
		return "0"
	}
	// 'aw mune pxey tsìng mrr pukap kinä
	for i, w := range naviVocab[0] {
		if s == w && w != "" {
			return strconv.FormatInt(int64(i), 8)
		}
	}
	// build regexp for all other numbers
	reString = "a?(mezazam|mezaza|pxezazam|pxezaza|tsìzazam|tsìzaza|mrrzazam|mrrzaza|puzazam|puzaza|kizazam|kizaza|zazam|zaza)?a?"
	reString += "a?(mevozam|mevoza|evozam|evoza|pxevozam|pxevoza|tsìvozam|tsìvoza|mrrvozam|mrrvoza|rrvozam|rrvoza|puvozam|puvoza|kivozam|kivoza|vozam|voza)?a?"
	reString += "a?(mezam|meza|ezam|eza|pxezam|pxeza|tsìzam|tsìza|mrrzam|mrrza|rrzam|rrza|puzam|puza|kizam|kiza|zam|za)?a?"
	reString += "a?(mevol|mevo|evol|evo|pxevol|pxevo|tsìvol|tsìvo|mrrvol|mrrvo|rrvol|rrvo|puvol|puvo|kivol|kivo|vol|vo)?a?"
	reString += "a?(aw|mun|un|pey|sìng|mrr|rr|fu|hin)?a?"
	re = regexp.MustCompile(reString)
	tmp := re.FindAllStringSubmatch(s, -1)
	if len(tmp) > 0 && len(tmp[0]) >= 1 {
		matchNumbers = tmp[0][1:]
	}

	for _, w := range matchNumbers {
		n += numTable[w]
	}

	return strconv.FormatInt(int64(n), 8)
}

func wordify(input string) string {
	rev := Reverse(input)
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
		if err != nil {
			return fmt.Sprintf("%s: %s\n", Text("invalidDecimalError"), input)
		}
		if !Valid(i, reverse) {
			return fmt.Sprintf("%s\n", Text("invalidIntError"))
		}
		o := strconv.FormatInt(int64(i), 8)
		output += fmt.Sprintf("Octal: %s\n", o)
		output += fmt.Sprintf("Na'vi: %s\n", wordify(o))
	} else {
		var io int64
		var err error
		if IsLetter(input) {
			io, err = strconv.ParseInt(unwordify(input), 8, 64)
		} else {
			io, err = strconv.ParseInt(input, 8, 64)
		}
		if err != nil {
			return fmt.Sprintf("%s: %s\n", Text("invalidOctalError"), input)
		}
		if !Valid(io, reverse) {
			return fmt.Sprintf("%s\n", Text("invalidIntError"))
		}
		d := strconv.FormatInt(int64(io), 10)
		o := strconv.FormatInt(int64(io), 8)
		output += fmt.Sprintf("Decimal: %s\n", d)
		if IsLetter(input) {
			output += fmt.Sprintf("Octal: %s\n", o)
		} else {
			output += fmt.Sprintf("Na'vi: %s\n", wordify(input))
		}
	}
	return output
}
