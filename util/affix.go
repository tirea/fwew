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
	//TODO
	return false
}

func FindAffixes(word string) (string, Affixes) {
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

	//TODO
	var found Affixes

	return "", found
}
