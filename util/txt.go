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

var usr, _ = user.Current()
var texts = map[string]string{}

func init() {
	// main program strings
	texts["name"] = "fwew"
	texts["tip"] = "type \"/help\" or \"/commands\" for more info"
	texts["author"] = "Tirea Aean"
	texts["header"] = fmt.Sprintf("%s\n%s\n", Version, texts["tip"])
	texts["languages"] = "de, eng, est, hu, nl, pl, ru, sv"
	texts["prompt"] = "~~> "

	// slash-commands Help
	texts["slashCommandHelp"] = "" +
		"/set       show currently set options, or set given options (separated by space)\n" +
		"/unset     unset given options (separated by space)\n" +
		"/list      list all words that meet given criteria\n" +
		"/random    display given number of random entries\n" +
		"/update    download and update the dictionary file\n" +
		"/commands  show this commands help text\n" +
		"/help      show main help text\n" +
		"/exit      exit/quit the program (aliases /quit /q /wc)\n\n" +
		"examples:\n" +
		"/set\n" +
		"/set i ipa\n" +
		"/unset i\n" +
		"/list pos has svin.\n" +
		"/list pos is vtrm.\n" +
		"/list word starts prr\n" +
		"/list word ends tut\n" +
		"/list word has kang\n" +
		"/list syllables >= 5\n" +
		"/list syllables = 1\n" +
		"/list words first 10\n" +
		"/list words last 20\n" +
		"/random 8\n" +
		"/random random\n" +
		"/random 2 where pos has vtr.\n" +
		"/random random where pos is n.\n" +
		"/update\n" +
		"/commands\n" +
		"/help\n" +
		"/exit\n"

	// flag strings
	texts["usage"] = "usage"
	texts["bin"] = strings.ToLower(texts["name"])
	texts["options"] = "options"
	texts["words"] = "words"
	texts["usageV"] = "show program & dictionary version numbers"
	texts["usageL"] = "use specified language \n\tValid values: " + texts["languages"]
	texts["usageI"] = "display infix location data"
	texts["usageIPA"] = "display IPA data"
	texts["usageS"] = "display source data"
	texts["usageP"] = "search for word(s) with specified part of speech"
	texts["usageR"] = "reverse the lookup direction from Na'vi->local to local->Na'vi"
	texts["usageA"] = "find all matches by using affixes to match the input word"
	texts["usageN"] = "convert numbers octal<->decimal"
	texts["usageM"] = "format output in markdown for bold and italic (mostly useful for fwew-discord bot)"
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
	texts["pre"] = "Prefixes"
	texts["inf"] = "Infixes"
	texts["suf"] = "Suffixes"
	texts["src"] = "source"

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
	texts["dictCloseError"] = "err 9: failed to close dictionary file (" + texts["dictionary"] + ")"
}

// Text function is the accessor for []string texts
func Text(s string) string {
	return texts[s]
}
