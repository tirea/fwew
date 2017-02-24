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
	"regexp"
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
	// this string only doesn't get translated or called from util.Text() because they're var names
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

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func contains(s []string, q []string) bool {
	if len(q) == 0 || len(s) == 0 {
		return false
	}
	// search for any instance of a thing in q...
	for _, x := range q {
		// ... that exists in s.
		for _, y := range s {
			if y == x {
				return true
			}
		}
	}

	return false
}

func prefix(w Word) Word {
	//TODO
	var re *regexp.Regexp
	var reString string
	var attempt string = ""
	var lenPre []string = []string{"pe", "fray", "tsay", "fay", "pay", "ay", "me", "pxe"}

	switch w.PartOfSpeech {
	case "n.":
		reString = "(pe)?(fray)?(tsay)?(fay)?(pay)?(ay)?(fra)?(fì)?(tsa)?(me)?(pxe)?(fne)?(munsna)?"
	case "adj.":
		reString = "(nì|a)?"
	case "vin.", "vtr.", "vim.", "vtrm.", "v.":
		reString = "(ketsuk|tsuk)?"
	default:
		return w // Not a type that has a prefix, return word without attempting.
	}

	re = regexp.MustCompile(reString)
	matchPrefixes := re.FindAllStringSubmatch(w.Target, -1)[0][1:]
	matchPrefixes = delete_empty(matchPrefixes)

	// no productive prefixes found; why bother to continue?
	if len(matchPrefixes) == 0 {
		return w
	}

	// build what prefixes to put on
	for _, p := range matchPrefixes {
		attempt = attempt + p
	}

	//TODO: fix this:
	// current fail: everything works except non-leniting prefixes

	// check for leniting prefix
	if contains(matchPrefixes, lenPre) {
		// lenite first
		w = lenite(w)
		// then add prefixes
		w.Attempt = attempt + w.Attempt
	} else {
		// otherwise just add the prefixes to create the attempt
		w.Attempt = attempt + w.Navi
	}

	w.Affixes["Prefixes"] = matchPrefixes

	//prodGerundAffix := []string{"tì", "us"}
	//prodActPartAffixPre := []string{"a", "us"}
	//prodPassPartAffixPre := []string{"a", "awn"}

	return w // placeholder
}

func suffix(w Word) Word {
	//TODO
	/*
		prodNSuffixes := []string{"tsyìp", "ìri", "nga'", "ìl", "pe", "yä", "ä", "it", "ri", "ru", "ti", "tu", "ur", "l", "o", "r", "t", "y"}
		prodVSuffixes := []string{"tswo", "yu"}
		prodNumSuffixes := []string{"ve"}
		prodActPartAffixSuf := []string{"us", "a"}
		prodPassPartAffixSuf := []string{"awn", "a"}
	*/
	return w
}

func infix(w Word) Word {
	//TODO
	/*
		hardCodeHax := map[string][]string{}
		hardCodeHax["poltxe"] = []string{"plltxe", "ol"}
		hardCodeHax["molte"] = []string{"mllte", "ol"}
		prodVInfixes := []string{"ìyev", "iyev", "äng", "eng", "ìlm", "ìly", "ìrm", "ìry", "ìsy", "alm", "aly", "äp", "arm", "ary", "asy", "ats", "eyk", "ìm", "imv", "irv", "ìy", "am", "ay", "ei", "er", "iv", "ol", "uy"}
	*/
	return w
}

func lenite(w Word) Word {
	// Have we already attempted lenition?
	if _, ok := w.Affixes["lenition"]; ok {
		return w
	}
	switch {
	case strings.HasPrefix(w.Navi, "kx"):
		w.Attempt = strings.Replace(w.Navi, "kx", "k", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "kx->k")
		return w
	case strings.HasPrefix(w.Navi, "px"):
		w.Attempt = strings.Replace(w.Navi, "px", "p", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "px->p")
		return w
	case strings.HasPrefix(w.Navi, "tx"):
		w.Attempt = strings.Replace(w.Navi, "tx", "t", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "tx->t")
		return w
	case strings.HasPrefix(w.Navi, "k"):
		w.Attempt = strings.Replace(w.Navi, "k", "h", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "k->h")
		return w
	case strings.HasPrefix(w.Navi, "p"):
		w.Attempt = strings.Replace(w.Navi, "p", "f", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "p->f")
		return w
	case strings.HasPrefix(w.Navi, "ts"):
		w.Attempt = strings.Replace(w.Navi, "ts", "s", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "ts->s")
		return w
	case strings.HasPrefix(w.Navi, "t"):
		w.Attempt = strings.Replace(w.Navi, "t", "s", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "t->s")
		return w
	case strings.HasPrefix(w.Navi, "'"):
		w.Attempt = strings.Replace(w.Navi, "'", "", 1)
		w.Affixes["lenition"] = append(w.Affixes["lenition"], "'->")
		return w
	default:
		return w
	}
}

func Reconstruct(w Word) Word {
	//TODO
	w = prefix(w)

	if w.Attempt == w.Target {
		return w
	}

	w = lenite(w)

	if w.Attempt == w.Target {
		return w
	}

	return Word{Id: "-1"}
}
