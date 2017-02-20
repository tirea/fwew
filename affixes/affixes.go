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
package affixes

import (
	"fmt"
	"strings"
)

type Word struct {
	Id             string
	LangCode       string
	Navi           string
	Target         string
	Attempt        string
	IPA            string
	InfixLocations string
	PartOfSpeech   string
	Definition     string
	Affixes        map[string][]string
}

func (w Word) String() string {
	return fmt.Sprintf("Id: %s\nLangCode: %s\nNavi: %s\nTarget: %s\nAttempt: %s\nIPA: %s\nInfixLocations: %s\nPartOfSpeech: %s\nDefinition: %s\nAffixes: %v\n",
		w.Id, w.LangCode, w.Navi, w.Target, w.Attempt, w.IPA, w.InfixLocations, w.PartOfSpeech, w.Definition, w.Affixes)
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
	w.Affixes = map[string][]string{}

	return w
}

func prefix(prefixed Word) Word {
	//TODO
	//var re *regexp.Regexp
	//var nounRe string = "(pe|fray|tsay|fay|pay|ay|fra)?(fì|tsa)?(me|pxe)?(fne)?(munsna)?"
	//var adjRe string = "(nì|a)?"
	//var vRe string = "(ketsuk|tsuk)?"

	//prodGerundAffix := []string{"tì", "us"}
	//prodActPartAffixPre := []string{"a", "us"}
	//prodPassPartAffixPre := []string{"a", "awn"}

	return prefixed // placeholder
}

func suffix(suffixed Word) Word {
	//TODO
	/*
		prodNSuffixes := []string{"tsyìp", "ìri", "nga'", "ìl", "pe", "yä", "ä", "it", "ri", "ru", "ti", "tu", "ur", "l", "o", "r", "t", "y"}
		prodVSuffixes := []string{"tswo", "yu"}
		prodNumSuffixes := []string{"ve"}
		prodActPartAffixSuf := []string{"us", "a"}
		prodPassPartAffixSuf := []string{"awn", "a"}
	*/
	return suffixed
}

func infix(infixed Word) Word {
	//TODO
	/*
		hardCodeHax := map[string][]string{}
		hardCodeHax["poltxe"] = []string{"plltxe", "ol"}
		hardCodeHax["molte"] = []string{"mllte", "ol"}
		prodVInfixes := []string{"ìyev", "iyev", "äng", "eng", "ìlm", "ìly", "ìrm", "ìry", "ìsy", "alm", "aly", "äp", "arm", "ary", "asy", "ats", "eyk", "ìm", "imv", "irv", "ìy", "am", "ay", "ei", "er", "iv", "ol", "uy"}
	*/
	return infixed
}

func lenite(unlenited Word) Word {
	var lenTable map[string]string = map[string]string{
		"kx": "k",
		"px": "p",
		"tx": "t",
		"k":  "h",
		"p":  "f",
		"t":  "s",
		"ts": "s",
		"'":  "",
	}

	for key, value := range lenTable {
		if strings.HasPrefix(unlenited.Navi, key) {
			unlenited.Attempt = strings.Replace(unlenited.Navi, key, value, 1)
			unlenited.Affixes["lenition"] = append(unlenited.Affixes["lenition"], key+"->"+value)
		}
	}

	return unlenited
}

func Reconstruct(w Word) Word {
	//TODO
	w = prefix(w)
	w = lenite(w)

	if w.Attempt == w.Target {
		return w
	} else {
		return Word{Id: "-1"}
	}
}
