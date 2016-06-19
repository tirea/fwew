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

package util

import (
	"strings"
	"regexp"
)

var nPrefixRe string = "(fì|tsa|me|pxe|ay|fay|tsay|fne|sna|munsna|fra|fray|pe|pem|pep|pay)?"
var vPrefixRe string = "(tsuk|ketsuk|tì)?"
var adjPrefixRe string = "(a|nì|ke|kel|kele)?"
var advPrefixRe string = "(nìk)?"
var nSuffixRe string = "(ìl|l|ti|it|t|ru|ur|r|yä|ä|ìri|ri|ya|fkeyk|o|pe|tsyìp|am|ay|y|äo|eo|fa|few|fpi|ftu|ftumfa|hu|io|ìlä|kam|kay|krrka|ka|kxamlä|lisre|lok|luke|mìkam|mì|mungwrr|na|nemfa|ne|nuä|pxaw|pxel|pximaw|maw|pxisre|rofa|ro|sìn|sko|sre|tafkip|takip|fkip|kip|talun|ta|teri|uo|vay|wä|yoa)?"
var numSuffixRe string = "(ve)?"
var adjSuffixRe string = "(a|pin)?"
var vSuffixRe string = "(yu|tswo)?"
var lentable = map[string]string{}

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

// handles verbs with infixes
func Infix(w string, inf_re string) [][]string {

	var orig string = inf_re
	var result [][]string

	inf_re = strings.Replace(inf_re,"<1>",Text("INFIX_0"),1)
	inf_re = strings.Replace(inf_re,"<2>",Text("INFIX_1"),1)
	inf_re = strings.Replace(inf_re,"<3>",Text("INFIX_2"),1)
	orig = strings.Replace(orig,"<1>","",1)
	orig = strings.Replace(orig,"<2>","",1)
	orig = strings.Replace(orig,"<3>","",1)

	re, err := regexp.Compile(inf_re)
	if err != nil { panic(err) }

	// pull out all infixes used and stash them in the result array
	result = re.FindAllStringSubmatch(w, -1)

	//replace the infixed verb in each array 1st position with root word
	if len(result) > 0 {
		result[0][0] = strings.Replace(result[0][0],result[0][0],orig,1)
	}

	return result
}

// handles other types of words having prefixes
func Prefix(w string, pre_re string, pos string) [][]string {

	var orig string = pre_re
	var result [][]string

	switch pos {
		case "n.":
			pre_re = nPrefixRe + pre_re
		case "v.", "vin.", "vim.", "vtr.", "vtrm.":
			pre_re = vPrefixRe + pre_re
		case "adj.":
			pre_re = adjPrefixRe + pre_re
		case "adv.":
			pre_re = advPrefixRe + pre_re
		default:
			break
	}

	re, err := regexp.Compile(pre_re)
	if err != nil { panic(err) }

	result = re.FindAllStringSubmatch(w, -1)

	if len(result) > 0 {
		result [0][0] = strings.Replace(result[0][0],result[0][0],orig,1)
	}

	return result
}

// handles other types of words having suffixes
func Suffix(w string, suf_re string, pos string) [][]string {

	var result [][]string

	return result
}

func Lenition(w string, w_re string) [][]string {

	var result [][]string

	return result
}