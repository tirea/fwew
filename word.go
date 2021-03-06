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

// Package main contains all the things. word.go is home to the Word struct.
package main

import "fmt"

// Word is a struct that contains all the data about a given word
type Word struct {
	ID             string
	LangCode       string
	Navi           string
	Target         string
	Attempt        string
	IPA            string
	InfixLocations string
	PartOfSpeech   string
	Definition     string
	Source         string
	Stressed       string
	Syllables      string
	InfixDots      string
	Affixes        map[string][]string
}

func (w Word) String() string {
	// this string only doesn't get translated or called from Text() because they're var names
	return fmt.Sprintf(""+
		"Id: %s\n"+
		"LangCode: %s\n"+
		"Navi: %s\n"+
		"Target: %s\n"+
		"Attempt: %s\n"+
		"IPA: %s\n"+
		"InfixLocations: %s\n"+
		"PartOfSpeech: %s\n"+
		"Definition: %s\n"+
		"Source: %s\n"+
		"Stressed: %s\n"+
		"Syllables: %s\n"+
		"InfixDots: %s\n"+
		"Affixes: %v\n",
		w.ID,
		w.LangCode,
		w.Navi,
		w.Target,
		w.Attempt,
		w.IPA,
		w.InfixLocations,
		w.PartOfSpeech,
		w.Definition,
		w.Source,
		w.Stressed,
		w.Syllables,
		w.InfixDots,
		w.Affixes)
}

// InitWordStruct is basically a constructor for Word struct
func InitWordStruct(w Word, dataFields []string) Word {
	//const (
	//	idField  int = 0 // dictionary.txt line Field 0 is Database ID
	//	lcField  int = 1 // dictionary.txt line field 1 is Language Code
	//	navField int = 2 // dictionary.txt line field 2 is Na'vi word
	//	ipaField int = 3 // dictionary.txt line field 3 is IPA data
	//	infField int = 4 // dictionary.txt line field 4 is Infix location data
	//	posField int = 5 // dictionary.txt line field 5 is Part of Speech data
	//	defField int = 6 // dictionary.txt line field 6 is Local definition
	//	srcField int = 7 // dictionary.txt line field 7 is Source data
	//  stsField int = 8 // dictionary.txt line field 8 is Stressed syllable #
	//  sylField int = 9 // dictionary.txt line field 9 is syllable breakdown
	//  ifdField int = 10 // dictionary.txt line field 10 is dot-style infix data
	//)
	w.ID = dataFields[idField]
	w.LangCode = dataFields[lcField]
	w.Navi = dataFields[navField]
	w.IPA = dataFields[ipaField]
	w.InfixLocations = dataFields[infField]
	w.PartOfSpeech = dataFields[posField]
	w.Definition = dataFields[defField]
	w.Source = dataFields[srcField]
	w.Stressed = dataFields[stsField]
	w.Syllables = dataFields[sylField]
	w.InfixDots = dataFields[ifdField]
	w.Affixes = map[string][]string{}

	return w
}

// CloneWordStruct is basically a copy constructor for Word struct
func CloneWordStruct(w Word) Word {
	var nw Word
	nw.ID = w.ID
	nw.LangCode = w.LangCode
	nw.Navi = w.Navi
	nw.Target = w.Target
	nw.Attempt = w.Attempt
	nw.IPA = w.IPA
	nw.InfixLocations = w.InfixLocations
	nw.PartOfSpeech = w.PartOfSpeech
	nw.Definition = w.Definition
	nw.Source = w.Source
	nw.Stressed = w.Stressed
	nw.Syllables = w.Syllables
	nw.InfixDots = w.InfixDots
	nw.Affixes = make(map[string][]string)
	for k := range w.Affixes {
		nw.Affixes[k] = make([]string, len(w.Affixes[k]))
		copy(nw.Affixes[k], w.Affixes[k])
	}
	return nw
}
