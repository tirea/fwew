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

// This util library handles all program strings
package util

import (
	"os/user"
	"path/filepath"
)

var usr, err = user.Current()
var texts = map[string]string{}

func init() {
	texts["name"] = "Fwew"
	texts["version"] = "1.3.1-BETA (02 FEB 2017)"
	texts["dictVersion"] = "Na'vi Dictionary 13.31 (07 JAN 2017)"
	texts["author"] = "Tirea Aean"
	texts["baseLang"] = "Na'vi"
	texts["header"] = texts["name"] + " " + texts["version"] + " by " + texts["author"] + "\n" +
		"Crossplatform " + texts["baseLang"] + " Dictionary Search" + "\n" +
		"fwew -h for usage, see README\n"
	texts["languages"] = "de, eng, est, hu, nl, pl, ru, sv"
	texts["language"] = "eng"
	texts["none"] = "No Results\n"
	texts["usageDebug"] = "Show extremely verbose probing"
	texts["usageV"] = "Show program & dictionary version numbers"
	texts["usageL"] = "Use specified language \n\tValid values: " + texts["languages"]
	texts["usageI"] = "Display infix location data"
	texts["usageIPA"] = "Display IPA data"
	texts["usageP"] = "Search for word(s) with specified part of speech"
	texts["usageR"] = "Reverse the lookup direction from Na'vi->Local to Local->Na'vi"
	texts["defaultFilter"] = "all"
	texts["homeDir"], _ = filepath.Abs(usr.HomeDir)
	texts["dataDir"] = filepath.Join(texts["homeDir"], ".fwew")
	texts["config"] = filepath.Join(texts["dataDir"], "config.json")
	texts["dictionary"] = filepath.Join(texts["dataDir"], "dictionary.tsv")
	texts["noDataError"] = "Dictionary data file missing or not installed.\nPlease Install Fwew (run ./install.sh)"
}

func Text(s string) string {
	return texts[s]
}

func SetText(i, s string) {
	texts[i] = s
}
