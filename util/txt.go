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
	"fmt"
	"os/user"
	"path/filepath"
)

var usr, err = user.Current()
var texts = map[string]string{}

func init() {
	// main program strings
	texts["name"] = "Fwew"
	texts["author"] = "Tirea Aean"
	texts["header"] = fmt.Sprintf("%s - Na'vi Dictionary Search - by %s\n`fwew -h` for usage, `fwew -v` for version\n", texts["name"], texts["author"])
	texts["languages"] = "de, eng, est, hu, nl, pl, ru, sv"
	texts["prompt"] = "Fwew:> "

	// flag strings
	texts["usageDebug"] = "Show extremely verbose probing"
	texts["usageV"] = "Show program & dictionary version numbers"
	texts["usageL"] = "Use specified language \n\tValid values: " + texts["languages"]
	texts["usageI"] = "Display infix location data"
	texts["usageIPA"] = "Display IPA data"
	texts["usageP"] = "Search for word(s) with specified part of speech"
	texts["usageR"] = "Reverse the lookup direction from Na'vi->Local to Local->Na'vi"
	texts["usageA"] = "Find all matches by using affixes to match the input word"
	texts["defaultFilter"] = "all"

	// file strings
	texts["homeDir"], _ = filepath.Abs(usr.HomeDir)
	texts["dataDir"] = filepath.Join(texts["homeDir"], ".fwew")
	texts["config"] = filepath.Join(texts["dataDir"], "config.json")
	texts["dictionary"] = filepath.Join(texts["dataDir"], "dictionary.txt")
	texts["prefixes"] = filepath.Join(texts["dataDir"], "prefixes.txt")
	texts["infixes"] = filepath.Join(texts["dataDir"], "infixes.txt")
	texts["suffixes"] = filepath.Join(texts["dataDir"], "suffixes.txt")

	// general message strings
	texts["affixes"] = "Affixes"
	texts["cset"] = "Currently set"
	texts["set"] = "set"
	texts["unset"] = "unset"

	// error message strings
	texts["none"] = "No Results\n"
	texts["noDataError"] = "Data file(s) missing or not installed.\nPlease Install Fwew (run ./install.sh)"
	texts["noOptionError"] = "No such option"
	texts["fileError"] = "File error"

}

func Text(s string) string {
	return texts[s]
}

func SetText(i, s string) {
	texts[i] = s
}
