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

//	This util library contains functions for handling affixed words

// COMPLETE LOGIC OVERHAUL REQUIRED

package util

import (
	"regexp"
	"strings"
	//"fmt"
)

var nPrefixes = []string{"fì", "tsa", "me", "pxe", "ay", "fay", "tsay", "fne", "sna", "munsna", "fra", "fray", "pe", "pem", "pep", "pay"}
var vPrefixes = []string{"tsuk", "ketsuk", "tì"}
var adjPrefixes = []string{"a", "nì", "ke", "kel", "kele"}
var advPrefixRe string = "(nìk)?"
var nSuffixes = []string{"ìl", "l", "ti", "it", "t", "ru", "ur", "r", "yä", "ä", "ìri", "ri", "ya", "fkeyk", "o", "pe", "tsyìp", "am", "ay", "y", "äo", "eo", "fa", "few", "fpi", "ftu", "ftumfa", "hu", "io", "ìlä", "kam", "kay", "krrka", "ka", "kxamlä", "lisre", "lok", "luke", "mìkam", "mì", "mungwrr", "na", "nemfa", "ne", "nuä", "pxaw", "pxel", "pximaw", "maw", "pxisre", "rofa", "ro", "sìn", "sko", "sre", "tafkip", "takip", "fkip", "kip", "talun", "ta", "teri", "uo", "vay", "wä", "yoa"}
var numSuffix string = "ve"
var adjSuffixes = []string{"a", "pin"}
var vSuffixes = []string{"yu", "tswo"}
var lentable = map[string]string{}
var result [][]string

func init() {
	lentable["kx"] = "k"
	lentable["px"] = "p"
	lentable["tx"] = "t"
	lentable["k"] = "h"
	lentable["p"] = "f"
	lentable["ts"] = "s"
	lentable["t"] = "s"
	lentable["'"] = ""
}

// simple containment check function
func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// handles verbs with infixes pretty well and fast
// doesn't work for poltxe and the like
func Infix(w string, inf_re string) [][]string {

	inf_re = strings.Replace(inf_re, "<1>", Text("INFIX_0"), 1)
	inf_re = strings.Replace(inf_re, "<2>", Text("INFIX_1"), 1)
	inf_re = strings.Replace(inf_re, "<3>", Text("INFIX_2"), 1)

	re, err := regexp.Compile(inf_re)
	if err != nil { panic(err) }

	// pull out all infixes used and stash them in the result array
	result = re.FindAllStringSubmatch(w, -1)

	return result
}

// handles other types of words having prefixes
// Doesn't work
func Prefix(w string, pre_re string, pos string) [][]string {

	var prefixRe string

	switch pos {
		case "n.":
			for _, p := range nPrefixes {
				prefixRe = prefixRe + "("+p+")?"
			}
			pre_re = prefixRe + pre_re
		case "v.", "vin.", "vim.", "vtr.", "vtrm.":
			for _, p := range vPrefixes {
				prefixRe = prefixRe + "("+p+")?"
			}
			pre_re = prefixRe + pre_re
		case "adj.":
			for _, p := range adjPrefixes {
				prefixRe = prefixRe + "("+p+")?"
			}
			pre_re = prefixRe + pre_re
		case "adv.":
			pre_re = advPrefixRe + pre_re
		default:
			prefixRe = ""
			//pre_re stays as-is
	}

	// DEBUG
	//fmt.Println("<DEBUG:util.Prefix() pre_re>", pre_re, "</DEBUG>")

	re, err := regexp.Compile(pre_re)
	if err != nil { panic(err) }

	result = re.FindAllStringSubmatch(w, -1)

	return result

}

// handles other types of words having suffixes
// stub
func Suffix(w string, suf_re string, pos string) [][]string {


	return result
}

// handles of course reversing lenition
// stub
func Lenition(w string, w_re string) [][]string {

	
	return result
}
