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

// This util library handles affix parsing of input
package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Word struct {
	Id             string
	LangCode       string
	Navi           string
	IPA            string
	InfixLocations string
	PartOfSpeech   string
	Definition     string
	Affixes        map[string]string
}

func InitWordStruct(w Word, dataFields []string) Word {
	const (
		idField  int = 0 // dictionary.tsv line Field 0 is Database ID
		lcField  int = 1 // dictionary.tsv line field 1 is Language Code
		navField int = 2 // dictionary.tsv line field 2 is Na'vi word
		ipaField int = 3 // dictionary.tsv line field 3 is IPA data
		infField int = 4 // dictionary.tsv line field 4 is Infix location data
		posField int = 5 // dictionary.tsv line field 5 is Part of Speech data
		defField int = 6 // dictionary.tsv line field 6 is Local definition
	)
	w.Id = dataFields[idField]
	w.LangCode = dataFields[lcField]
	w.Navi = dataFields[navField]
	w.IPA = dataFields[ipaField]
	w.InfixLocations = dataFields[infField]
	w.PartOfSpeech = dataFields[posField]
	w.Definition = dataFields[defField]

	return w
}

func prefix(prefixed Word) Word {
	//TODO
	return prefixed
}

func suffix(suffixed Word) Word {
	//TODO
	return suffixed
}

func infix(infixed Word) Word {
	//TODO
	return infixed
}

func Lenite(unlenited Word) Word {
	//TODO
	return unlenited

}

func Stem(given Word) Word {
	//TODO
	/*
	 * Stemming algorithm:
	 * 0) Base case: affixed.Navi == given.Navi, return affixed. Else:
	 * 1) Iterate dictionary.tsv
	 * 2) Apply known rules to current line in order to reconstruct given.
	 * 3) profit???
	 */
	var affixed Word
	var fields []string

	// Prepare file for searching
	dictData, err := os.Open(Text("dictionary"))
	defer dictData.Close()
	if err != nil {
		fmt.Println(errors.New(Text("noDataError")))
		os.Exit(1)
	}
	scanner := bufio.NewScanner(dictData)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// Store the fields of the line into fields array in lowercase
		fields = strings.Split(line, "\t")
		// Put the stuff from fields into the Word struct
		affixed = InitWordStruct(affixed, fields)
		// ???
	}

	/*
		hardCodeHax := map[string][]string{}
		hardCodeHax["'awlo"] = []string{"'aw", "lo"}
		hardCodeHax["melo"] = []string{"mune", "lo"}
		hardCodeHax["pxelo"] = []string{"pxe", "lo"}
		hardCodeHax["poltxe"] = []string{"plltxe", "ol"}
		hardCodeHax["molte"] = []string{"mllte", "ol"}
	*/

	/*
		// Comment to avoid 'declared and not used' error
		prodNPrefixes := []string{"munsna", "fì", "fne", "fra", "sna", "tsa", "a"}
		prodVPrefixes := []string{"ketsuk", "tsuk"}
		prodAdjPrefixes := []string{"nì"}
		prodLenPrefixes := []string{"fray", "tsay", "fay", "pay", "pxe", "ay", "me", "pe"}
		prodVInfixes := []string{"ìyev", "iyev", "äng", "eng", "ìlm", "ìly", "ìrm", "ìry", "ìsy", "alm", "aly", "äp", "arm", "ary", "asy", "ats", "eyk", "ìm", "imv", "irv", "ìy", "am", "ay", "ei", "er", "iv", "ol", "uy"}
		prodNSuffixes := []string{"tsyìp", "ìri", "nga'", "ìl", "pe", "yä", "ä", "it", "ri", "ru", "ti", "tu", "ur", "a", "l", "o", "r", "t", "y"}
		prodVSuffixes := []string{"tswo", "yu"}
		prodNumSuffixes := []string{"ve"}
		prodGerundAffix := []string{"tì", "us"}
		prodActPartAffixPre := []string{"us", "a"}
		prodActPartAffixSuf := []string{"a", "us"}
		prodPassPartAffixPre := []string{"awn", "a"}
		prodPassPartAffixSuf := []string{"a", "awn"}
	*/

	return affixed // Placeholder
}
