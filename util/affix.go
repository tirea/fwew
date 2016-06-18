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

	var result [][]string

	inf_re = strings.Replace(inf_re,"<1>",Text("INFIX_0"),1)
	inf_re = strings.Replace(inf_re,"<2>",Text("INFIX_1"),1)
	inf_re = strings.Replace(inf_re,"<3>",Text("INFIX_2"),1)

	re, err := regexp.Compile(inf_re)
	if err != nil { panic(err) }

	// pull out all infixes used and stash them in the result array
	result = re.FindAllStringSubmatch(w, -1)

	return result
}

// handles other types of words having prefixes
func Prefix(w string, pre_re string) [][]string {

	var result [][]string

	return result

}

// handles other types of words having suffixes
func Suffix(w string, suf_re string) [][]string {

	var result [][]string

	return result
}

func Lenition(w string, w_re string) [][]string {

	var result [][]string

	return result
}