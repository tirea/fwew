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

// Package main contains all the things. affixes_test.go -- you guess it -- tests affixes.go functions.
package main

import (
	"flag"
	"testing"
)

func init() {
	// set relevant option flag(s)
	configuration = ReadConfig()
	debug = flag.Bool("d", configuration.DebugMode, Text("usageD"))
	flag.Parse()
}

func TestPrefix(t *testing.T) {
	w0 := Word{
		Target:       "tsatute",
		Attempt:      "tute",
		PartOfSpeech: "n.",
		Affixes:      map[string][]string{},
	}
	w0 = prefix(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}

func TestSuffix(t *testing.T) {
	w0 := Word{
		Target:       "eltuti",
		Attempt:      "eltu",
		PartOfSpeech: "n.",
		Affixes:      map[string][]string{},
	}
	w0 = suffix(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}

func TestInfix(t *testing.T) {
	w0 := Word{
		Target:         "täpeykiyevarängon",
		Attempt:        "taron",
		InfixLocations: "t<0><1>ar<2>on",
		PartOfSpeech:   "vtr.",
		Affixes:        map[string][]string{},
	}
	w0 = infix(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}

func TestReconstruct(t *testing.T) {
	w0 := Word{
		Target:       "tolaron",
		Attempt:      "taron",
		PartOfSpeech: "vtr.",
		Affixes:      map[string][]string{},
	}
	w0 = Reconstruct(w0)
	if w0.Attempt != w0.Target {
		t.Errorf("Got %s, Want %s", w0.Attempt, w0.Target)
	}
}
