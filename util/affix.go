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
	"fmt"
	"os"
	"strings"
)

type Word struct {
	Id string 
    LangCode string
    Navi string
    IPA string
    InfixLocations string
    PartOfSpeech string
    Definition string
    Affixes map[string]string
}

func prefix(prefixed Word) Word{
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

	hardCodeHax := map[string][]string{}
	hardCodeHax["'awlo"] = []string{"'aw", "lo"}
	hardCodeHax["melo"] = []string{"mune", "lo"}
	hardCodeHax["pxelo"] = []string{"pxe", "lo"}
	hardCodeHax["poltxe"] = []string{"plltxe", "ol"}
	hardCodeHax["molte"] = []string{"mllte", "ol"}

	/* comment to avoid 'declared and not used' error
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
