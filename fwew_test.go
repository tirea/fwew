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

const (
	numWords   = 2501
	numMemes   = 7
	totalWords = numWords + numMemes
)

func TestSimilarity(t *testing.T) {
	f0 := similarity("fmetok", "fmetok")
	if f0 != 1.0 {
		t.Errorf("Wanted %f, Got %f", 1.0, f0)
	}

	f1 := similarity("meoauniaea", "eltu")
	if f1 != 0.0 {
		t.Errorf("Wanted %f, Got %f", 0.0, f1)
	}
}

// helper function for TestFwew, basically a means to consider two Word structs equal
func testEqualWord(w0, w1 Word) bool {
	if w0.ID == w1.ID && w0.Navi == w1.Navi {
		return true
	}
	return false
}

func TestFwew(t *testing.T) {
	// Set relevant option flags
	configuration = ReadConfig()
	reverse = flag.Bool("r", false, Text("usageR"))
	language = flag.String("l", configuration.Language, Text("usageL"))
	posFilter = flag.String("p", configuration.PosFilter, Text("usageP"))
	useAffixes = flag.Bool("a", configuration.UseAffixes, Text("usageA"))
	flag.Parse()

	var w Word

	w0 := fwew("fmetok")[0]
	w = Word{ID: "392", Navi: "fmetok"}
	if !testEqualWord(w, w0) {
		t.Errorf("Wanted %s, Got %s\n", w, w0)
	}

	w1 := fwew("")
	if w1 != nil {
		t.Errorf("empty string did not yield empty Word slice\n")
	}

	w2 := fwew("tseyä")[0]
	w = Word{ID: "5268", Navi: "tsaw"}
	// if w3.ID != "5268" && w3.Navi != "tsaw" {
	if !testEqualWord(w, w2) {
		t.Errorf("Wanted %s, Got %s\n", w, w2)
	}

	w5 := fwew("oey")[0]
	w = Word{ID: "1380", Navi: "oe"}
	if !testEqualWord(w, w5) {
		t.Errorf("Wanted %s, Got %s\n", w, w5)
	}

	w6 := fwew("ngey")[0]
	w = Word{ID: "1348", Navi: "nga"}
	if !testEqualWord(w, w6) {
		t.Errorf("Wanted %s, Got %s\n", w, w6)
	}

	*reverse = true
	w7 := fwew("test")[0]
	w = Word{ID: "392", Navi: "fmetok"}
	if !testEqualWord(w, w7) {
		t.Errorf("Wanted %s, Got %s\n", w, w7)
	}

	*useAffixes = false
	*reverse = false
	w8 := fwew("fmetok")
	if len(w8) != 1 {
		t.Errorf("Wanted 1 word, Got %d\n", len(w8))
	}
}

// helper function for TestSyllableCount, basically cuts down on repetitive code
func testEqualInt(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Wanted %d, Got %d\n", expected, actual)
	}
}

func TestSyllableCount(t *testing.T) {
	var w Word

	w = Word{Navi: "nari si"}
	testEqualInt(t, 3, syllableCount(w))

	w = Word{Navi: "lu"}
	testEqualInt(t, 1, syllableCount(w))

	w = Word{Navi: "ätxäle si"}
	testEqualInt(t, 4, syllableCount(w))

	w = Word{Navi: "tireapängkxo"}
	testEqualInt(t, 5, syllableCount(w))

	w = Word{Navi: "tìng tseng"}
	testEqualInt(t, 2, syllableCount(w))
}

func TestCountLines(t *testing.T) {
	testEqualInt(t, totalWords, countLines())
}
