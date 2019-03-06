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

// Package main contains all the things. affixes.go handles affix parsing of input.
package main

import (
	"fmt"
	"regexp"
	"strings"
)

var debug = configuration.DebugMode

func prefix(w Word) Word {
	var (
		re            *regexp.Regexp
		reString      string
		attempt       string
		matchPrefixes []string
	)

	// pull this out of the switch because the pos data for verbs is so irregular,
	// the switch condition would be like 25 possibilities long
	if strings.HasPrefix(w.PartOfSpeech, "v") ||
		strings.HasPrefix(w.PartOfSpeech, "svin.") || w.PartOfSpeech == "" {
		inf := w.Affixes[Text("inf")]
		if len(inf) > 0 && (inf[0] == "us" || inf[0] == "awn") {
			reString = "(a|tì)?"
		} else {
			reString = "(a)?(ketsuk|tsuk)?"
		}
	} else {
		switch w.PartOfSpeech {
		case "n.", "pn.":
			reString = "(pe)?(fray)?(tsay)?(fay)?(pay)?(ay)?(fra)?(fì)?(tsa)?(me)?(pxe)?(fne)?(munsna)?"
		case "dem.", "dem., pn.":
			reString = "(pe)?(fray)?(tsay)?(fay)?(pay)?(ay)?(fra)?(me)?(pxe)?(fne)?(munsna)?"
		case "adj.":
			reString = "(nìk|nì|a)?(ke|a)?"
		default:
			return w // Not a type that has a prefix, return word without attempting.
		}
	}

	if strings.HasPrefix(w.Attempt, "e") {
		reString = reString + "(e)?"
		w.Attempt = w.Attempt[1:]
	} else if strings.HasPrefix(w.Attempt, "'e") {
		reString = reString + "('e)?"
		w.Attempt = w.Attempt[2:]
	}
	reString = reString + w.Attempt + ".*"
	if debug {
		fmt.Printf("Prefix reString: %s\n", reString)
	}
	re = regexp.MustCompile(reString)
	tmp := re.FindAllStringSubmatch(w.Target, -1)
	if len(tmp) > 0 && len(tmp[0]) >= 1 {
		matchPrefixes = tmp[0][1:]
	}
	matchPrefixes = DeleteEmpty(matchPrefixes)

	if debug {
		fmt.Printf("matchPrefixes: %s\n", matchPrefixes)
	}

	// no productive prefixes found; why bother to continue?
	if len(matchPrefixes) == 0 {
		return w
	}

	// build what prefixes to put on
	for _, p := range matchPrefixes {
		attempt = attempt + p
	}

	w.Attempt = attempt + w.Attempt

	matchPrefixes = DeleteElement(matchPrefixes, "e")
	if len(matchPrefixes) > 0 {
		w.Affixes[Text("pre")] = matchPrefixes
	}

	return w
}

func suffix(w Word) Word {
	var (
		re            *regexp.Regexp
		reString      string
		attempt       string
		matchSuffixes []string
		adp           = []string{"mungwrr", "kxamlä", "tafkip", "pxisre", "pximaw",
			"ftumfa", "mìkam", "nemfa", "takip", "lisre", "talun", "krrka", "teri",
			"fkip", "pxaw", "pxel", "luke", "rofa", "fpi", "ftu", "kip", "vay", "lok",
			"maw", "sìn", "sre", "few", "kam", "kay", "nuä", "sko", "yoa", "äo", "eo",
			"fa", "hu", "ka", "mì", "na", "ne", "ta", "io", "uo", "ro", "wä", "sì?",
		}
	)
	const (
		adjSufRe string = "(a)?"
		nSufRe   string = "(nga')?(tsyìp)?(o)?(pe)?(ìri)?(ìlä)?(ìl)?(eyä)?(yä)?(ä)?(it)?(ri)?(ru)?(ti)?(tu)?(ur)?(l)?(r)?(t)?(y)?"
	)

	// verbs
	if strings.HasPrefix(w.PartOfSpeech, "v") ||
		strings.HasPrefix(w.PartOfSpeech, "svin.") || w.PartOfSpeech == "" {
		inf := w.Affixes[Text("inf")]
		pre := w.Affixes[Text("pre")]
		// word is verb with <us> or <awn>
		if len(inf) == 1 && (inf[0] == "us" || inf[0] == "awn") {
			// it's a tì-<us> gerund; treat it like a noun
			if len(pre) > 0 && ContainsStr(pre, "tì") && inf[0] == "us" {
				reString = nSufRe
				for _, s := range adp {
					reString += "(" + s + ")?"
				}
				// Just a regular <us> or <awn> verb
			} else {
				reString = adjSufRe
			}
			// It's a tsuk/ketsuk adj from a verb
		} else if len(inf) == 0 && Contains(pre, []string{"tsuk", "ketsuk"}) {
			reString = adjSufRe
		} else {
			reString = "(tswo|yu)?"
			for _, s := range adp {
				reString += "(" + s + ")?"
			}
		}
	} else {
		switch w.PartOfSpeech {
		// nouns and noun-likes
		case "n.", "pn.", "prop.n.", "inter.", "dem.", "dem., pn.":
			// reString = nSufRe
			for _, s := range adp {
				reString += "(" + s + ")?"
			}
			reString += nSufRe // move this down here to stop cases cannibalizing adp
		// adjectives
		case "adj.":
			reString = adjSufRe
		// numbers
		case "num.":
			reString = "(ve)?(a)?"
		default:
			return w // Not a type that has a suffix, return word without attempting.
		}
	}

	// o -> e vowel shift support
	if strings.HasSuffix(w.Attempt, "o") {
		reString = strings.Replace(w.Attempt, "o", "[oe]", -1) + reString
		// a -> e vowel shift support
	} else if strings.HasSuffix(w.Attempt, "a") {
		reString = strings.Replace(w.Attempt, "a", "[ae]", -1) + reString
	} else {
		reString = w.Attempt + reString
	}

	if debug {
		fmt.Printf("Suffix reString: %s\n", reString)
	}
	re = regexp.MustCompile(reString)
	tmp := re.FindAllStringSubmatch(w.Target, -1)
	if len(tmp) > 0 && len(tmp[0]) >= 1 {
		matchSuffixes = tmp[0][1:]
	}
	matchSuffixes = DeleteEmpty(matchSuffixes)
	if debug {
		fmt.Printf("matchSuffixes: %s\n", matchSuffixes)
	}

	// no productive prefixes found; why bother to continue?
	if len(matchSuffixes) == 0 {
		return w
	}

	// build what prefixes to put on
	for _, p := range matchSuffixes {
		attempt = attempt + p
	}

	// o -> e vowel shift support for pronouns with -yä
	if w.PartOfSpeech == "pn." && ContainsStr(matchSuffixes, "yä") {
		if strings.HasSuffix(w.Attempt, "o") {
			w.Attempt = strings.TrimSuffix(w.Attempt, "o") + "e"
			// a -> e vowel shift support
		} else if strings.HasSuffix(w.Attempt, "a") {
			w.Attempt = strings.TrimSuffix(w.Attempt, "a") + "e"
		}
	}
	w.Attempt = w.Attempt + attempt
	w.Affixes[Text("suf")] = matchSuffixes
	return w
}

func infix(w Word) Word {
	// Have we already attempted infixes?
	if _, ok := w.Affixes[Text("inf")]; ok {
		return w
	}
	// Does the word even have infix positions??
	if w.InfixLocations == "\\N" {
		return w
	}

	var (
		re              *regexp.Regexp
		reString        string
		attempt         string
		pos0InfixRe     = "(äp)?(eyk)?"
		pos1InfixRe     = "(ìyev|iyev|ìlm|ìly|ìrm|ìry|ìsy|alm|aly|arm|ary|asy|ìm|imv|ilv|irv|ìy|am|ay|er|iv|ol|us|awn)?"
		pos2InfixRe     = "(eiy|ei|äng|eng|ats|uy)?"
		pos0InfixString string
		pos1InfixString string
		pos2InfixString string
		matchInfixes    []string
	)

	// Hardcode hack for z**enke
	if w.Navi == "zenke" && (strings.Contains(w.Target, "uy") || strings.Contains(w.Target, "ats")) {
		w.InfixLocations = strings.Replace(w.InfixLocations, "ke", "eke", 1)
	}

	reString = strings.Replace(w.InfixLocations, "<0>", pos0InfixRe, 1)
	// handle <ol>ll and <er>rr
	if strings.Contains(reString, "<1>ll") {
		reString = strings.Replace(reString, "<1>ll", pos1InfixRe+"(ll)?", 1)
	} else if strings.Contains(w.InfixLocations, "<1>rr") {
		reString = strings.Replace(reString, "<1>rr", pos1InfixRe+"(rr)?", 1)
	} else {
		reString = strings.Replace(reString, "<1>", pos1InfixRe, 1)
	}
	reString = strings.Replace(reString, "<2>", pos2InfixRe, 1)
	if debug {
		fmt.Printf("Infix reString: %s\n", reString)
	}

	re, _ = regexp.Compile(reString)
	tmp := re.FindAllStringSubmatch(w.Target, -1)
	if len(tmp) > 0 && len(tmp[0]) >= 1 {
		matchInfixes = tmp[0][1:]
	}
	matchInfixes = DeleteEmpty(matchInfixes)
	matchInfixes = DeleteElement(matchInfixes, "ll")
	matchInfixes = DeleteElement(matchInfixes, "rr")

	for _, i := range matchInfixes {
		if i == "äp" || i == "eyk" {
			pos0InfixString = pos0InfixString + i
		} else if ContainsStr([]string{"eiy", "ei", "äng", "eng", "ats", "uy"}, i) {
			pos2InfixString = i
		} else {
			pos1InfixString = i
		}
	}

	attempt = strings.Replace(w.InfixLocations, "<0>", pos0InfixString, 1)
	attempt = strings.Replace(attempt, "<1>", pos1InfixString, 1)
	attempt = strings.Replace(attempt, "<2>", pos2InfixString, 1)

	if ContainsStr(matchInfixes, "eiy") {
		eiy := Index(matchInfixes, "eiy")
		matchInfixes[eiy] = "ei"
	}
	if debug {
		fmt.Printf("matchInfixes: %s\n", matchInfixes)
	}

	// handle <ol>ll and <er>rr
	if strings.Contains(attempt, "olll") {
		attempt = strings.Replace(attempt, "olll", "ol", 1)
	} else if strings.Contains(attempt, "errr") {
		attempt = strings.Replace(attempt, "errr", "er", 1)
	}
	w.Attempt = attempt
	if len(matchInfixes) != 0 {
		w.Affixes[Text("inf")] = matchInfixes
	}

	return w
}

func lenite(w Word) Word {
	// Have we already attempted lenition?
	if _, ok := w.Affixes["lenition"]; ok {
		return w
	}
	lenTable := map[string]string{
		"kx": "k",
		"px": "p",
		"tx": "t",
		"k":  "h",
		"p":  "f",
		"ts": "s",
		"t":  "s",
		"'":  "",
	}
	for k, v := range lenTable {
		if strings.HasPrefix(w.Navi, k) {
			w.Attempt = strings.Replace(w.Attempt, k, v, 1)
			w.Affixes["lenition"] = append(w.Affixes["lenition"], k+"->"+v)
			return w
		}
	}
	return w
}

func matches(w Word) bool {
	return w.Attempt == w.Target
}

// Reconstruct is the main function of affixes.go, responsible for the affixing algorithm
func Reconstruct(w Word) Word {

	w.Attempt = strings.ToLower(w.Navi)
	w.Target = strings.ToLower(w.Target)

	// clone w as wl
	wl := CloneWordStruct(w)
	// wl will be the lenited version of w
	wl = lenite(wl)

	// only try to infix verbs
	if strings.HasPrefix(w.PartOfSpeech, "v") || strings.HasPrefix(w.PartOfSpeech, "svin.") {
		w = infix(w)
		if debug {
			fmt.Println("INFIX")
			fmt.Printf("Navi: %s | Attempt: %s | Target: %s\n", w.Navi, w.Attempt, w.Target)
		}
		if matches(w) {
			return w
		}
	}

	w = prefix(w)
	if debug {
		fmt.Println("PREFIX")
		fmt.Printf("Navi: %s | Attempt: %s | Target: %s\n", w.Navi, w.Attempt, w.Target)
	}
	if matches(w) {
		return w
	}
	wl = prefix(wl)
	if matches(wl) {
		return wl
	}

	if !strings.HasPrefix(w.Attempt, w.Target[0:1]) {
		w = lenite(w)
		if debug {
			fmt.Println("LENITE")
			fmt.Printf("Navi: %s | Attempt: %s | Target: %s\n", w.Navi, w.Attempt, w.Target)
		}
		if matches(w) {
			return w
		}
	}

	w = suffix(w)
	if debug {
		fmt.Println("SUFFIX")
		fmt.Printf("Navi: %s | Attempt: %s | Target: %s\n", w.Navi, w.Attempt, w.Target)
	}
	if matches(w) {
		return w
	}
	wl = suffix(wl)
	if matches(wl) {
		return wl
	}

	w = lenite(w)
	if debug {
		fmt.Println("LENITE")
		fmt.Printf("Navi: %s | Attempt: %s | Target: %s\n", w.Navi, w.Attempt, w.Target)
	}
	if matches(w) {
		return w
	}

	if debug {
		fmt.Println("GIVING UP")
	}
	return Word{ID: "-1"}
}
