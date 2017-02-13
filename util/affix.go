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

type Affixes struct {
	Prefixes  map[string][]string
	Infixes   map[string][]string
	Suffixes  map[string][]string
	IsLenited bool
}

func stripPrefixes(word string) (string, []string) {
	//TODO
	return "", []string{}
}

func stripSuffixes(word string) (string, []string) {
	//TODO
	return "", []string{}
}

func stripInfixes(word string) (string, []string) {
	//TODO
	return "", []string{}
}

func unLenite(word string) (string, bool) {
	//TODO
	return "", false

}

func exists(word string) bool {
	word = strings.ToLower(word)

	affixData, err := os.Open(Text("infixes"))
	defer affixData.Close()
	if err != nil {
		fmt.Println(errors.New(Text("noDataError")))
		os.Exit(1)
	}
	scanner := bufio.NewScanner(affixData)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if word == line {
			affixData.Close()
			return true
		}
	}
	return false
}

func FindAffixes(word string) (string, Affixes) {
	//TODO
	/*
		Affix-stripping algorithm:
			1) Strip all productive prefixes.
			2) If the word exists now, done. Else goto 3.
			3) Strip all productive suffixes.
			4) If the word exists now,  done. Else goto 5.
			5) Undo lenition.
			6) If the word exists now, done. Else goto 7.
			7) Strip all infixes.
			8) If the word exists now, done. Else no results, done.
	*/
	var affixes Affixes

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

	// Strip all productive prefixes

	return "", affixes // Placeholder
}
