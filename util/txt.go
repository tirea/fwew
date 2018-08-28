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

// Package util handles general program stuff. txt.go handles program strings.
package util

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
)

var usr, err = user.Current()
var texts = map[string]string{}

func init() {
	// main program strings
	texts["name"] = "fwew"
	texts["tagline"] = "The Best Na'vi Dictionary on the Command Line"
	texts["tip"] = "type \"/help\" or \"/commands\" for more info"
	texts["author"] = "Tirea Aean"
	texts["header"] = fmt.Sprintf("%s\n%s\n%s\n", Version, texts["tagline"], texts["tip"])
	texts["languages"] = "de, eng, est, hu, nl, pl, ru, sv"
	texts["prompt"] = "~>> "

	// slash-commands Help
	texts["/set"] = "/set       show currently set options, or set given options (separated by space)\n"
	texts["/unset"] = "/unset     unset given options (separated by space)\n"
	texts["/list"] = "/list      list all words with given part of speech\n"
	texts["/update"] = "/update    download and update the dictionary file\n"
	texts["/commands"] = "/commands  Show this commands help text\n"
	texts["/help"] = "/help      Show main help text\n"
	texts["/exit"] = "/exit      exit/quit the program (aliases /quit /q /wc)\n\n"
	texts["/examples"] = fmt.Sprintf("%s:\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"examples", "/set", "/set i ipa", "/unset i", "/list svin.", "/update", "/commands", "/help", "/exit")
	texts["slashCommandHelp"] = texts["/set"] + texts["/unset"] + texts["/list"] +
		texts["/update"] + texts["/commands"] + texts["/help"] + texts["/exit"] + texts["/examples"]

	// flag strings
	texts["usage"] = "usage"
	texts["bin"] = strings.ToLower(texts["name"])
	texts["options"] = "options"
	texts["words"] = "words"
	texts["usageV"] = "show program & dictionary version numbers"
	texts["usageL"] = "use specified language \n\tValid values: " + texts["languages"]
	texts["usageI"] = "display infix location data"
	texts["usageIPA"] = "display IPA data"
	texts["usageP"] = "search for word(s) with specified part of speech"
	texts["usageR"] = "reverse the lookup direction from Na'vi->local to local->Na'vi"
	texts["usageA"] = "find all matches by using affixes to match the input word"
	texts["usageN"] = "convert numbers octal<->decimal"
	texts["defaultFilter"] = "all"

	// file strings
	texts["homeDir"], _ = filepath.Abs(usr.HomeDir)
	texts["dataDir"] = filepath.Join(texts["homeDir"], ".fwew")
	texts["config"] = filepath.Join(texts["dataDir"], "config.json")
	texts["dictionary"] = filepath.Join(texts["dataDir"], "dictionary.txt")
	texts["dictURL"] = "https://tirea.learnnavi.org/dictionarydata/dictionary.txt"
	texts["dlSuccess"] = texts["dictURL"] + "\nsaved to\n" + texts["dictionary"] + "\n"

	// general message strings
	texts["cset"] = "currently set"
	texts["set"] = "set"
	texts["unset"] = "unset"

	// error message strings
	texts["none"] = "no results\n"
	texts["noDataError"] = "err 1: failed to open and/or read dictionary file (" + texts["dictionary"] + ")"
	texts["fileError"] = "err 2: failed to open and/or read configuration file (" + texts["config"] + ")"
	texts["noOptionError"] = "err 3: invalid option"
	texts["invalidIntError"] = "err 4: input must be a decimal integer in range 0 <= n <= 32767 or octal integer in range 0 <= n <= 77777"
	texts["invalidOctalError"] = "err 5: invalid octal integer"
	texts["invalidDecimalError"] = "err 6: invalid decimal integer"
	texts["invalidLanguageError"] = "err 7: invalid language option"
	texts["invalidPOSFilterError"] = "err 8: invalid part of speech filter"
}

// Text function is the accessor for []string texts
func Text(s string) string {
	return texts[s]
}

// SetText is the setter for []string texts
func SetText(i, s string) {
	texts[i] = s
}
