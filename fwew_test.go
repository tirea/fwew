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

// Package main contains all the things. affixes_test.go tests fwew.go functions.
package main

import (
	"flag"
	"testing"
)

func TestSimilarity(t *testing.T) {

}

// helper function for TestFwew, basically a means to consider two Word structs equal
func equal(w0, w1 Word) bool {
	if w0.ID == w1.ID && w0.Navi == w1.Navi {
		return true
	} else {
		return false
	}
}

func TestFwew(t *testing.T) {
	// Set relevant option flags
	configuration = ReadConfig()
	reverse = flag.Bool("r", false, Text("usageR"))
	language = flag.String("l", configuration.Language, Text("usageL"))
	posFilter = flag.String("p", configuration.PosFilter, Text("usageP"))
	useAffixes = flag.Bool("a", configuration.UseAffixes, Text("usageA"))
	flag.Parse()

	w0 := fwew("fmetok")[0]
	w1 := Word{ID: "392", Navi: "fmetok"}

	if !equal(w0, w1) {
		t.Errorf("Got %s, Want %s", w0.ID, w1.ID)
	}
}